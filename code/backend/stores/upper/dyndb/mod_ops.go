package dyndb

import (
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/filter"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/processer"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/upper/db/v4"
)

func (d *DynDB) newRow(txid uint32, req dyndb.NewRowReq) (int64, error) {
	var id int64

	req.Data[dyndb.KeyVersion] = 1

	req.ModCtx.TableName = d.tns.Table(req.TenantId, req.Group, req.Table)

	modsig, err := req.ModCtx.JSON()
	if err != nil {
		return 0, err
	}

	req.Data[dyndb.KeyModSig] = string(modsig)

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
			pp.Println("@inser_req", req)
			pp.Println("@inser_err", err)
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

func (d *DynDB) NewBatchRows(txid uint32, req dyndb.NewBatchRowReq) ([]int64, error) {

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

func (d *DynDB) deleteRow(txid uint32, req dyndb.DeleteRowReq) error {

	modctx, err := req.ModCtx.JSON()
	if err != nil {
		return err
	}

	tablename := d.tns.Table(req.TenantId, req.Group, req.Table)

	return d.txOr(txid, func(sess db.Session) error {
		if d.vendor == store.VendorSqlite {
			_, err := sess.SQL().Exec("select temphia_delete_record(?, ?, ?)", tablename, modctx, strconv.FormatInt(req.Id, 10))
			return err
		} else {
			tbl := sess.Collection(tablename)
			return tbl.Find(dyndb.KeyPrimary, req.Id).Delete()
		}

	})
}

func (d *DynDB) deleteRowBatch(txid uint32, req dyndb.DeleteRowBatchReq) ([]int64, error) {
	fcond, err := filter.Transform(req.FilterConds)
	if err != nil {
		return nil, err
	}

	modctx, err := req.ModCtx.JSON()
	if err != nil {
		return nil, err
	}

	result := make([]int64, 0)

	tablename := d.tns.Table(req.TenantId, req.Group, req.Table)

	err = d.txOr(txid, func(sess db.Session) error {

		if d.vendor == store.VendorSqlite {
			keys := make([]Key, 0)

			err = sess.Collection(tablename).Find(fcond).All(keys)
			if err != nil {
				return err
			}

			var buf strings.Builder

			for id, v := range keys {
				if id != 0 {
					buf.WriteByte(',')
				}
				buf.WriteString(strconv.FormatInt(v.ID, 10))
				result = append(result, v.ID)
			}

			_, err := sess.SQL().Exec("select temphia_delete_record(?, ?, ?)", tablename, modctx, buf.String())
			return err

		} else {
			deleter := sess.SQL().DeleteFrom(tablename).Where(fcond).Amend(func(queryIn string) (queryOut string) {
				return queryIn + " RETURNING __id"
			})

			rows, err := sess.SQL().Query(deleter.String(), deleter.Arguments()...)
			if err != nil {
				return nil
			}

			defer rows.Close()

			for rows.Next() {
				var k Key
				err := rows.Scan(&k.ID)
				if err != nil {
					return err
				}
				result = append(result, k.ID)
			}
			if err := rows.Err(); err != nil {
				return err
			}
			return nil
		}
	})

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (d *DynDB) deleteRowMulti(txid uint32, req dyndb.DeleteRowMultiReq) error {

	modctx, err := req.ModCtx.JSON()
	if err != nil {
		return err
	}

	tablename := d.tns.Table(req.TenantId, req.Group, req.Table)

	return d.txOr(txid, func(sess db.Session) error {
		if d.vendor == store.VendorSqlite {
			var buf strings.Builder

			for id, v := range req.Ids {
				if id != 0 {
					buf.WriteByte(',')
				}
				buf.WriteString(strconv.FormatInt(v, 10))
			}

			_, err := sess.SQL().Exec("select temphia_delete_record(?, ?, ?)", tablename, modctx, buf.String())
			return err
		} else {
			tbl := sess.Collection(d.tns.Table(req.TenantId, req.Group, req.Table))
			return tbl.Find(dyndb.KeyPrimary, req.Ids).Delete()
		}

	})
}

func (d *DynDB) updateRow(txid uint32, req dyndb.UpdateRowReq) (map[string]interface{}, error) {
	var data map[string]interface{}

	req.ModCtx.TableName = d.tns.Table(req.TenantId, req.Group, req.Table)
	modsig, err := req.ModCtx.JSON()
	if err != nil {
		return nil, err
	}
	req.Data[dyndb.KeyModSig] = string(modsig)

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
			req.Data[dyndb.KeyVersion] = req.Version + 1
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
