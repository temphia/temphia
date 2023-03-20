package datagroup

import (
	"strconv"
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

const (
	MethodNewRow      = "new_row"
	MethodGetRow      = "get_row"
	MethodUpdateRow   = "update_row"
	MethodDeleteRows  = "delete_rows"
	MethodSimpleQuery = "simple_query"
	MethodTicket      = "ticket"
)

type DGModule struct {
	binder   etypes.ExecutorBinder
	res      *entities.Resource
	dynsrc   dyndb.DynSource
	tenantId string
	group    string
}

func (d *DGModule) IPC(method string, path string, args xtypes.LazyData) (xtypes.LazyData, error) {

	if MethodTicket == method {
		app := d.binder.GetApp().(xtypes.App)
		signer := app.GetDeps().Signer().(service.Signer)

		uctx := d.binder.InvokerGet().ContextUser()

		tok, err := signer.SignData(d.tenantId, &claim.Data{
			TenantId:   d.tenantId,
			Type:       claim.CTypeData,
			UserID:     uctx.Id,
			UserGroup:  uctx.Group,
			SessionID:  uctx.SessionId,
			DeviceId:   uctx.DeviceId,
			DataSource: d.dynsrc.Name(),
			DataGroup:  d.group,
			DataTables: []string{"*"},
			IsExec:     true,
		})

		return d.response(tok, err)
	}

	txid, table, rowid := d.extractPath(path)

	dhub := d.dynsrc.GetDataTableHub(d.tenantId, d.group)

	switch method {
	case MethodNewRow:
		data := make(map[string]any)
		err := args.AsObject(&data)
		if err != nil {
			return nil, err
		}

		modctx := dyndb.ModCtx{
			TableName: table,
		}

		uctx := d.binder.InvokerGet().ContextUser()
		if uctx != nil {
			modctx.UserId = uctx.Id
		}

		return d.response(dhub.NewRow((txid), dyndb.NewRowReq{
			TenantId: d.tenantId,
			Group:    d.group,
			Table:    table,
			Data:     data,
			ModCtx:   modctx,
		}))

	case MethodGetRow:
		return d.response(dhub.GetRow(txid, dyndb.GetRowReq{
			TenantId:  d.tenantId,
			Group:     d.group,
			Table:     table,
			Id:        rowid,
			SkipCache: true,
		}))

	case MethodUpdateRow:
		data := make(map[string]any)
		err := args.AsObject(&data)
		if err != nil {
			return nil, err
		}

		modctx := dyndb.ModCtx{
			TableName: table,
		}
		uctx := d.binder.InvokerGet().ContextUser()
		if uctx != nil {
			modctx.UserId = uctx.Id
		}

		return d.response(dhub.UpdateRow(txid, dyndb.UpdateRowReq{
			TenantId: d.tenantId,
			Id:       rowid,
			Version:  0,
			Group:    d.group,
			Table:    table,
			Data:     data,
			ModCtx:   modctx,
		}))

	case MethodDeleteRows:
		return d.response(nil, dhub.DeleteRow(txid, dyndb.DeleteRowReq{
			TenantId: d.tenantId,
			Group:    d.group,
			Table:    table,
			Id:       rowid,
		}))
	case MethodSimpleQuery:
		req := dyndb.SimpleQueryReq{}
		err := args.AsObject(&req)
		if err != nil {
			return nil, err
		}
		return d.response(dhub.SimpleQuery(txid, req))

	}

	return nil, easyerr.NotFound()

}

func (d *DGModule) Close() error {
	d.dynsrc = nil
	d.res = nil
	return nil
}

// private

func (d *DGModule) extractPath(path string) (uint32, string, int64) {

	/*
		<txid><table_slug><row_id>
		0/table1/12
		0/table2
	*/

	contents := strings.Split(path, "/")
	if len(contents) != 2 || len(contents) != 3 {
		panic("err invalid path")
	}

	txid := uint32(0)
	if contents[0] != "0" {
		_txid, err := strconv.ParseUint(contents[0], 10, 32)
		if err != nil {
			panic(err)
		}
		txid = uint32(_txid)
	}

	rowid := int64(0)
	if len(contents) == 3 {
		_rowid, err := strconv.ParseInt(contents[2], 10, 64)
		if err != nil {
			panic(err)
		}

		rowid = _rowid
	}

	return txid, contents[1], rowid
}

func (d *DGModule) response(data any, err error) (xtypes.LazyData, error) {
	if err != nil {
		return nil, err
	}

	return lazydata.NewAnyData(data), nil
}
