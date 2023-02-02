package dyndb

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/processer"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/upper/db/v4"
)

func (d *DynDB) newRow(txid uint32, req store.NewRowReq) (int64, error) {
	var id int64

	req.Data[store.KeyVersion] = 1

	req.ModCtx.TableName = d.tns.Table(req.TenantId, req.Group, req.Table)

	modsig, err := req.ModCtx.JSON()
	if err != nil {
		return 0, err
	}

	req.Data[store.KeyModSig] = string(modsig)

	err = d.beforeDBRow(req.TenantId, req.Group, req.Table, req.Data)
	if err != nil {
		return 0, err
	}

	pp.Println("@inserting =>", req)

	err = d.txOr(txid, func(sess db.Session) error {
		tbl := sess.Collection(d.tns.Table(req.TenantId, req.Group, req.Table))

		pp.Println("@@inserting..", req)

		ir, err := tbl.Insert(req.Data)
		if err != nil {
			return err
		}
		id = ir.ID().(int64)
		return nil
	})
	return id, err
}

type Key struct {
	ID int64 `db:"__id"`
}

func (d *DynDB) NewBatchRows(txid uint32, req store.NewBatchRowReq) ([]int64, error) {

	keys := make(map[string]struct{})

	for _, data := range req.Data {
		for k := range data {
			keys[k] = struct{}{}
		}
	}

	for _, data := range req.Data {
		for k := range keys {
			_, ok := data[k]
			if !ok {
				data[k] = nil
			}
		}
	}

	keyMap := make([]Key, 0)
	err := d.txOr(txid, func(sess db.Session) error {
		inserter := sess.SQL().InsertInto(
			d.tns.Table(req.TenantId, req.Group, req.Table),
		).Returning("__id").Batch(len(req.Data))

		for _, data := range req.Data {
			inserter.Values(data)
		}

		keyMap := make([]Key, len(req.Data))
		inserter.NextResult(&keyMap)

		inserter.Done()
		return inserter.Err()
	})
	if err != nil {
		return nil, err
	}

	ids := make([]int64, 0, len(keyMap))
	for _, k := range keyMap {
		ids = append(ids, (k.ID))
	}

	return ids, nil
}

func (d *DynDB) deleteRows(txid uint32, req store.DeleteRowReq) error {
	return d.txOr(txid, func(sess db.Session) error {
		tbl := sess.Collection(d.tns.Table(req.TenantId, req.Group, req.Table))
		return tbl.Find(store.KeyPrimary, req.Id).Delete()
	})
}

func (d *DynDB) updateRow(txid uint32, req store.UpdateRowReq) (map[string]interface{}, error) {
	var data map[string]interface{}

	req.ModCtx.TableName = d.tns.Table(req.TenantId, req.Group, req.Table)
	modsig, err := req.ModCtx.JSON()
	if err != nil {
		return nil, err
	}
	req.Data[store.KeyModSig] = string(modsig)

	err = d.beforeDBRow(req.TenantId, req.Group, req.Table, req.Data)
	if err != nil {
		return nil, err
	}

	err = d.txOr(txid, func(sess db.Session) error {
		var query db.Updater
		if req.Version == -1 {
			query = sess.SQL().
				Update(d.tns.Table(req.TenantId, req.Group, req.Table)).
				Where("__id = ?", req.Id).
				Set(req.Data).
				Amend(func(queryIn string) string {
					// fixme => amend set verison = version + 1
					pp.Println(queryIn)

					return queryIn + " RETURNING *"
				})

		} else {
			req.Data[store.KeyVersion] = req.Version + 1
			query = sess.SQL().
				Update(d.tns.Table(req.TenantId, req.Group, req.Table)).
				Where("__id = ?", req.Id).
				And("__version = ?", req.Version).
				Set(req.Data).
				Amend(func(queryIn string) string {
					return queryIn + " RETURNING *"
				})
		}

		r, err := sess.SQL().Query(query.String(), query.Arguments()...)
		if err != nil {
			return err
		}
		data, err = dbutils.GetScan(r)
		return err
	})

	if err != nil {
		return nil, err
	}

	pcr := d.processer(req.TenantId, req.Group, req.Table)
	err = pcr.FromRowDBType(data)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (d *DynDB) beforeDBRow(tenantId, group, table string, data map[string]interface{}) error {
	cols, err := d.cache.CachedColumns(tenantId, group, table)
	if err != nil {
		return err
	}

	pcr := processer.New(d.vendor, cols)
	return pcr.ToRowDBType(data)
}
