package datagroup

import (
	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

var _ etypes.ModuleBuilder = (*DGModuleBuilder)(nil)

func NewBuilder(app any) (etypes.ModuleBuilder, error) {
	return &DGModuleBuilder{}, nil
}

type DGModuleBuilder struct{}

func (d *DGModuleBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {

	deps := opts.Binder.GetApp().(xtypes.App).GetDeps()

	target, err := opts.Resource.SplitTarget(2)
	if err != nil {
		return nil, err
	}

	dynhub := deps.DataHub().(dyndb.DataHub)

	dgmod := &DatagroupModule{
		binder:   opts.Binder,
		res:      opts.Resource,
		dynsrc:   dynhub,
		group:    target[1],
		tenantId: opts.Resource.TenantId,
		source:   target[0],
		modipc:   nil,
	}

	dgmod.modipc = modipc.NewModIPC(dgmod)
	return dgmod, nil
}