package dyndb

import (
	"fmt"
	"strings"

	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/stores/upperdb/dyndb/filter"
	"github.com/temphia/temphia/code/backend/stores/upperdb/dyndb/processer"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/upper/db/v4"
)

func (d *DynDB) simpleQuery(txid uint32, req dyndb.SimpleQueryReq) (*dyndb.QueryResult, error) {
	records := make([]map[string]interface{}, 0)

	conds, err := filter.Transform(d.vendor, req.FilterConds)
	if err != nil {
		return nil, err
	}

	err = d.txOr(txid, func(sess db.Session) error {

		pp.Println(req)

		selects := make([]interface{}, 0, len(req.Selects))

		for _, s := range req.Selects {
			selects = append(selects, s)
		}

		orderBy := dyndb.KeyPrimary
		if req.OrderBy != "" {
			orderBy = req.OrderBy
		}

		if req.Desc {
			orderBy = fmt.Sprintf("%s DESC", req.OrderBy)
		}

		err = sess.SQL().
			Select(selects...).
			From(d.tns.Table(req.TenantId, req.Group, req.Table)).
			Where(conds).
			OrderBy(orderBy).
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

	resp := &dyndb.QueryResult{
		Count:   int64(len(records)),
		Page:    req.Page,
		Rows:    records,
		Columns: cols,
	}

	return resp, nil
}

func (d *DynDB) processer(tenantId, group, table string) processer.Processer {
	cols, err := d.cache.CachedColumns(tenantId, group, table)
	if err != nil {
		return nil
	}
	return processer.New(d.vendor, cols)
}

func (d *DynDB) getRow(txid uint32, req dyndb.GetRowReq) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := d.txOr(txid, func(sess db.Session) error {
		tbl := sess.Collection(d.tns.Table(req.TenantId, req.Group, req.Table))
		return tbl.Find(dyndb.KeyPrimary, req.Id).One(&data)
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

func (d *DynDB) joinQuery(txid uint32, req dyndb.JoinReq) (*dyndb.JoinResult, error) {

	records := make([]map[string]any, 0)

	cond1, err := filter.TransformWithPrefix(d.vendor, req.ParentFilters, "parent.")
	if err != nil {
		return nil, err
	}

	cond2, err := filter.TransformWithPrefix(d.vendor, req.ChildFilters, "child.")
	if err != nil {
		return nil, err
	}

	for k, v := range cond2 {
		cond1[k] = v
	}

	orderby := ""
	if req.Desc {
		orderby = fmt.Sprintf("-parent.%s", req.OrderBy)
	} else {
		orderby = fmt.Sprintf("parent.%s", req.OrderBy)
	}

	parent := d.tns.Table(req.TenantId, req.Group, req.Parent)
	child := d.tns.Table(req.TenantId, req.Group, req.Child)

	err = d.txOr(txid, func(sess db.Session) error {
		return sess.SQL().
			Select().
			From(fmt.Sprintf("%s AS parent", parent)).
			Where(cond1).
			Join(fmt.Sprintf("%s As child", child)).
			On(fmt.Sprintf("parent.%s = child.%s", req.OnParent, req.OnChild)).
			OrderBy(orderby).
			All(&records)
	})

	if err != nil {
		return nil, err
	}

	pcols, err := d.cache.CachedColumns(req.TenantId, req.Group, req.Parent)
	if err != nil {
		return nil, err
	}

	ccols, err := d.cache.CachedColumns(req.TenantId, req.Group, req.Child)
	if err != nil {
		return nil, err
	}

	return &dyndb.JoinResult{
		Rows:       records,
		ParentCols: pcols,
		ChildCols:  ccols,
	}, nil
}

func (d *DynDB) multiJoinQuery(txid uint32, req dyndb.MultiJoinReq) (*dyndb.MultiJoinResult, error) {

	records := make([]map[string]any, 0)

	cond1, err := filter.TransformWithPrefix(d.vendor, req.ParentFilters, "parent.")
	if err != nil {
		return nil, err
	}

	for idx, fg := range req.Fragments {
		fname := fmt.Sprintf("inner%d", idx)
		cond, err := filter.TransformWithPrefix(d.vendor, fg.Filters, fmt.Sprintf("%s.", fname))
		if err != nil {
			return nil, err
		}

		for k, v := range cond {
			cond1[k] = v
		}
	}

	parent := d.tns.Table(req.TenantId, req.Group, req.Parent)

	orderby := ""
	if req.Desc {
		orderby = fmt.Sprintf("-parent.%s", req.OrderBy)
	} else {
		orderby = fmt.Sprintf("parent.%s", req.OrderBy)
	}

	err = d.txOr(txid, func(sess db.Session) error {
		sqlq := sess.SQL().
			Select().
			From(fmt.Sprintf("%s AS parent", parent)).
			Where(cond1)

		for idx, fg := range req.Fragments {
			fname := fmt.Sprintf("inner%d", idx)
			fgtable := d.tns.Table(req.TenantId, req.Group, fg.Name)

			sqlq = sqlq.Join(fmt.Sprintf("%s As %s", fgtable, fname)).
				On(fmt.Sprintf("parent.%s = %s.%s", req.OnParent, fname, fg.OnColumn))
		}

		return sqlq.OrderBy(orderby).All(&records)
	})

	if err != nil {
		return nil, err
	}

	pcols, err := d.cache.CachedColumns(req.TenantId, req.Group, req.Parent)
	if err != nil {
		return nil, err
	}

	return &dyndb.MultiJoinResult{
		Rows:    records,
		Columns: pcols,
	}, nil

}

func (d *DynDB) ftsQuery(txid uint32, req dyndb.FTSQueryReq) (*dyndb.QueryResult, error) {

	records := make([]map[string]any, 0)
	err := d.txOr(txid, func(sess db.Session) error {
		conds, err := filter.Transform(d.vendor, req.Filters)
		if err != nil {
			return err
		}

		searchTerm := req.SearchTerm
		if !req.UsePattern {
			searchTerm = "%" + searchTerm + "%"
		}

		likeQ := db.Cond{fmt.Sprintf("%s LIKE", req.SearchColumn): searchTerm}

		return sess.SQL().
			Select().
			From(d.tns.Table(req.TenantId, req.Group, req.Table)).
			Where(conds).
			And(likeQ).
			Paginate(uint(req.Count)).
			Page(uint(req.Page + 1)).
			All(&records)
	})

	if err != nil {
		return nil, err
	}

	cols, err := d.cache.CachedColumns(req.TenantId, req.Group, req.Table)
	if err != nil {
		return nil, err
	}

	return &dyndb.QueryResult{
		Count:   int64(len(records)),
		Page:    req.Page,
		Rows:    records,
		Columns: cols,
	}, err
}

func (d *DynDB) templateQuery(txid uint32, req dyndb.TemplateQueryReq) (*dyndb.QueryResult, error) {
	return nil, nil
}

func (d *DynDB) RefResolve(txid uint32, tenantId, gslug string, req *dyndb.RefResolveReq) (*dyndb.QueryResult, error) {

	rows := make([]map[string]interface{}, 0)

	err := d.txOr(txid, func(sess db.Session) error {
		key := ""

		switch req.Type {
		case dyndb.RefHardPriId, dyndb.RefSoftPriId:
			key = fmt.Sprintf("%s IN", dyndb.KeyPrimary)
		case dyndb.RefSoftText:
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

	return &dyndb.QueryResult{
		Rows:    rows,
		Columns: cols,
	}, nil
}

func (d *DynDB) refLoad(txid uint32, tenantId, gslug string, req *dyndb.RefLoadReq) (*dyndb.QueryResult, error) {
	rows := make([]map[string]interface{}, 0)

	err := d.txOr(txid, func(sess db.Session) error {
		sess.Collection(d.tns.Table(tenantId, gslug, req.Target)).Find(
			db.Cond{
				fmt.Sprintf("%s >", dyndb.KeyPrimary): req.CursorRowId,
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

	return &dyndb.QueryResult{
		Rows:    rows,
		Count:   int64(len(rows)),
		Page:    0,
		Columns: cols,
	}, nil
}

func (d *DynDB) reverseRefLoad(txid uint32, tenantId, gslug string, req *dyndb.RevRefLoadReq) (*dyndb.QueryResult, error) {
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

	return &dyndb.QueryResult{
		Columns: cols,
		Rows:    rows,
	}, nil

}

func (d *DynDB) sqlQueryRaw(txid uint32, tenantId, group, qstr string) (any, error) {
	var rs []map[string]any

	err := d.txOr(txid, func(sess db.Session) error {
		rows, err := sess.SQL().Query(qstr)
		if err != nil {
			return err
		}

		rs, err = dbutils.SelectScan(rows)
		return err
	})

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (d *DynDB) sqlQueryScopped(txid uint32, tenantId, group, qstr string) (any, error) {
	qstr = strings.TrimSpace(removeSQLComments(qstr))

	if strings.HasPrefix(qstr, "FETCH syetem_tables") {
		return d.ListTables(tenantId, group)
	} else if strings.HasPrefix(qstr, "FETCH syetem_columns") {
		tname, err := extractTableName(qstr)
		if err != nil {
			return nil, err
		}
		return d.ListColumns(tenantId, group, tname)
	}

	result, err := d.hsql.Transform(tenantId, group, nil, qstr)
	if err != nil {
		return nil, err
	}

	var rs []map[string]any

	err = d.txOr(txid, func(sess db.Session) error {

		rows, err := sess.SQL().Query(result.TransformedQuery)
		if err != nil {
			return err
		}
		rs, err = dbutils.SelectScan(rows)
		return err
	})

	if err != nil {
		return nil, err
	}

	return rs, nil
}
