package adapter

import (
	"strconv"

	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz/adapter"
)

var _ etypes.ModuleBuilder = (*AdapterModBuilder)(nil)

func NewBuilder(app any) (etypes.ModuleBuilder, error) {
	return &AdapterModBuilder{}, nil
}

type AdapterModBuilder struct{}

func (p *AdapterModBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	app := opts.Binder.GetApp().(xtypes.App)
	invoker := opts.Binder.GetInvoker()

	pp.Println("@id", opts.Resource.Target)

	id, err := strconv.ParseInt(opts.Resource.Target, 10, 64)
	if err != nil {
		return nil, err
	}

	return New(opts.Resource.TenantId, id, app.GetServer().GetAdapterHub().(adapter.AdapterHub), invoker), nil
}

func New(tenantId string, id int64, ahub adapter.AdapterHub, ib bindx.Invoker) *AdapterMod {

	bm := &AdapterMod{
		adapterHub: ahub,
		adapterId:  id,
		inBinder:   ib,
	}

	return bm

}
