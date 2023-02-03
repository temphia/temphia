package dyndb

import (
	"strconv"
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type DyndbModule struct {
	binder etypes.ExecutorBinder
	res    *entities.Resource
	dynsrc store.DynSource
	group  string
	table  string
}

func (d *DyndbModule) IPC(method string, path string, args xtypes.LazyData) (xtypes.LazyData, error) {

	txid, table, rowid := d.extractPath(path)

	switch method {
	case "new_row":
		return d.response(d.dynsrc.NewRow((txid), store.NewRowReq{
			TenantId: "",
			Group:    d.group,
			Table:    table,
			Data:     nil, // fixme
		}))

	case "get_row":
		return d.response(d.dynsrc.GetRow(txid, store.GetRowReq{
			TenantId:  "",
			Group:     d.group,
			Table:     table,
			Id:        rowid,
			SkipCache: true,
		}))

	case "update_row":
		return d.response(d.dynsrc.UpdateRow(txid, store.UpdateRowReq{
			TenantId: "",
			Id:       rowid,
			Version:  0,
			Group:    d.group,
			Table:    table,
			Data:     nil,            // fixme
			ModCtx:   store.ModCtx{}, // fixme
		}))

	case "delete_rows":
		return d.response(nil, d.dynsrc.DeleteRow(txid, store.DeleteRowReq{
			TenantId: "",
			Group:    d.group,
			Table:    table,
			Id:       rowid,
		}))
	case "simple_query":
		req := store.SimpleQueryReq{}
		err := args.AsObject(&req)
		if err != nil {
			return nil, err
		}
		return d.response(d.dynsrc.SimpleQuery(txid, req))
	default:
		return nil, easyerr.NotFound()
	}

}

func (d *DyndbModule) Close() error {
	d.binder = nil
	d.dynsrc = nil
	d.res = nil
	return nil
}

// private

func (d *DyndbModule) extractPath(path string) (uint32, string, int64) {

	contents := strings.Split(path, "/")
	txid := uint32(0)
	if contents[0] != "0" {
		_txid, err := strconv.ParseUint(contents[0], 10, 32)
		if err != nil {
			panic(err)
		}
		txid = uint32(_txid)
	}

	if d.table == "" {
		// "0/<table>/<rowid>"
		// "0/<table>"

		switch len(contents) {
		case 3:
			rowid, err := strconv.ParseInt(contents[3], 10, 64)
			if err != nil {
				panic(err)
			}

			return txid, contents[1], rowid
		case 2:

			return txid, contents[1], 0
		default:
			panic("invalid path")
		}

	}

	// "0/<rowid>"
	// "0"

	rowid := int64(0)

	if len(contents) == 2 {
		_rowid, err := strconv.ParseInt(contents[0], 10, 64)
		if err != nil {
			panic(err)
		}
		rowid = _rowid
	}

	return txid, d.table, rowid

}

func (d *DyndbModule) response(data any, err error) (xtypes.LazyData, error) {

	return nil, nil
}
