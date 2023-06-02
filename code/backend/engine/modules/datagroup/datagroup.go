package datagroup

import (
	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"

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
	return d.modipc.Handle(method, args)
}

func (d *DatagroupModule) Close() error {
	d.dynsrc = nil
	d.res = nil
	return nil
}
