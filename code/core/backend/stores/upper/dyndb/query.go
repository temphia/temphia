package dyndb

import (
	"fmt"

	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/stores/upper/dyndb/processer"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/upper/db/v4"
)

func (d *DynDB) simpleQuery(txid uint32, req store.SimpleQueryReq) (*store.QueryResult, error) {
	records := make([]map[string]interface{}, 0)
	err := d.txOr(txid, func(sess db.Session) error {

		pp.Println(req)

		selects := make([]interface{}, 0, len(req.Selects))

		for _, s := range req.Selects {
			selects = append(selects, s)
		}

		conds, err := transformFilters(req.FilterConds)
		if err != nil {
			return err
		}

		// tbl := sess.Collection(d.tns.Table(req.TenantId, req.Group, req.Table))

		// fixme => search join
		// https://tour.upper.io/sql-builder/02

		err = sess.SQL().
			Select(selects...).
			From(d.tns.Table(req.TenantId, req.Group, req.Table)).
			Where(conds).
			OrderBy(store.KeyPrimary).
			Paginate(uint(req.Count)).
			Page(uint(req.Page + 1)). // cz page starts from 1
			All(&records)

		return err
	})

	if err != nil {
		return nil, err
	}

	cols, err := d.cache.CachedColumns(req.TenantId, req.Group, req.Table)
	if err != nil {
		return nil, err
	}

	pcr := processer.New(d.vendor, cols)
	err = pcr.FromRowsDBType(records)
	if err != nil {
		return nil, err
	}

	resp := &store.QueryResult{
		Columns: cols,
		Count:   req.Count,
		Page:    req.Page,
		Rows:    records,
	}

	if req.LoadExtraMeta {

		revRefs, err := d.ListReverseColumnRef(req.TenantId, req.Group, req.Table)
		if err != nil {
			return nil, err
		}

		// fixme => remove server_side hooks ?

		// hooks, err := d.ListHook(req.TenantId, req.Group, req.Table)
		// if err != nil {
		// 	return nil, err
		// }

		views, err := d.ListView(req.TenantId, req.Group, req.Table)
		if err != nil {
			return nil, err
		}

		resp.ExtraMeta = &store.QueryMeta{
			ReverseRefs: revRefs,
			Hooks:       nil,
			Views:       views,
		}
	}

	return resp, err
}

func (d *DynDB) processer(tenantId, group, table string) processer.Processer {
	cols, err := d.cache.CachedColumns(tenantId, group, table)
	if err != nil {
		return nil
	}
	return processer.New(d.vendor, cols)
}

func (d *DynDB) getRow(txid uint32, req store.GetRowReq) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := d.txOr(txid, func(sess db.Session) error {
		tbl := sess.Collection(d.tns.Table(req.TenantId, req.Group, req.Table))
		return tbl.Find(store.KeyPrimary, req.Id).One(&data)
	})
	if err != nil {
		return nil, err
	}

	err = d.processer(req.TenantId, req.Group, req.Table).FromRowDBType(data)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (d *DynDB) _FTSQuery(txid uint32, req store.FTSQueryReq) (*store.QueryResult, error) {
	return nil, nil
}

func (d *DynDB) rawQuery(txid uint32, req store.RawQueryReq) (interface{}, error) {

	return nil, nil
}

func (d *DynDB) templateQuery(txid uint32, req store.TemplateQueryReq) (*store.QueryResult, error) {
	return nil, nil
}

func (d *DynDB) RefResolve(txid uint32, tenantId, gslug string, req *store.RefResolveReq) (*store.QueryResult, error) {

	rows := make([]map[string]interface{}, 0)

	err := d.txOr(txid, func(sess db.Session) error {
		key := ""

		switch req.Type {
		case store.RefHardPriId, store.RefSoftPriId:
			key = fmt.Sprintf("%s IN", store.KeyPrimary)
		case store.RefSoftText:
			key = fmt.Sprintf("%s IN", req.Object)
		default:
			return easyerr.NotImpl()
		}

		return sess.Collection(d.tns.Table(tenantId, gslug, req.Target)).Find(db.Cond{
			key: req.RowIds,
		}).All(&rows)
	})

	if err != nil {
		return nil, err
	}

	cols, err := d.cache.CachedColumns(tenantId, gslug, req.Target)
	if err != nil {
		return nil, err
	}

	pcr := processer.New(d.vendor, cols)
	err = pcr.FromRowsDBType(rows)
	if err != nil {
		return nil, err
	}

	return &store.QueryResult{
		Columns: cols,
		Rows:    rows,
	}, nil
}

func (d *DynDB) refLoad(txid uint32, tenantId, gslug string, req *store.RefLoadReq) (*store.QueryResult, error) {
	rows := make([]map[string]interface{}, 0)

	err := d.txOr(txid, func(sess db.Session) error {
		sess.Collection(d.tns.Table(tenantId, gslug, req.Target)).Find(
			db.Cond{
				fmt.Sprintf("%s >", store.KeyPrimary): req.CursorRowId,
			},
		).All(&rows)
		return nil
	})
	if err != nil {
		return nil, err
	}

	cols, err := d.cache.CachedColumns(tenantId, gslug, req.Target)
	if err != nil {
		return nil, err
	}

	pcr := processer.New(d.vendor, cols)
	err = pcr.FromRowsDBType(rows)
	if err != nil {
		return nil, err
	}

	return &store.QueryResult{
		Columns: cols,
		Rows:    rows,
	}, nil
}

func (d *DynDB) reverseRefLoad(txid uint32, tenantId, gslug string, req *store.RevRefLoadReq) (*store.QueryResult, error) {
	rows := make([]map[string]interface{}, 0)

	err := d.txOr(txid, func(sess db.Session) error {

		sess.Collection(d.tns.Table(tenantId, gslug, req.TargetTable)).Find(
			db.Cond{req.TargetColumn: req.CurrentItem},
		).All(&rows)
		return nil
	})
	if err != nil {
		return nil, err
	}

	cols, err := d.cache.CachedColumns(tenantId, gslug, req.TargetTable)
	if err != nil {
		return nil, err
	}

	pcr := processer.New(d.vendor, cols)
	err = pcr.FromRowsDBType(rows)
	if err != nil {
		return nil, err
	}

	return &store.QueryResult{
		Columns: cols,
		Rows:    rows,
	}, nil

}
