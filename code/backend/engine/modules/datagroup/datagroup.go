package datagroup

import (
	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service"

	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type DatagroupModule struct {
	tenantId string
	source   string
	group    string
	dynsrc   dyndb.DynSource
	binder   bindx.Bindings
	res      *entities.Resource
	modipc   *modipc.ModIPC
}

func (d *DatagroupModule) Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error) {
	if method == "ticket" {
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

		if err != nil {
			return nil, err
		}

		return lazydata.NewAnyData(tok), nil
	}

	return d.modipc.Handle(method, args)
}

func (d *DatagroupModule) Close() error {
	d.dynsrc = nil
	d.res = nil
	return nil
}
