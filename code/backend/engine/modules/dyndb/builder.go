package dyndb

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type DyndbBuilder struct {
}

func (DyndbBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	return New(opts), nil
}

func (DyndbBuilder) Init(app any) error {
	return nil
}

func New(opts etypes.ModuleOptions) *DyndbModule {
	deps := opts.Binder.GetApp().(xtypes.App).GetDeps()

	dynhub := deps.DataHub().(dyndb.DataHub).GetSource("default", opts.Resource.TenantId) // fixme => get source from resource

	return &DyndbModule{
		binder:   opts.Binder,
		res:      opts.Resource,
		dynsrc:   dynhub,
		group:    "",
		table:    "",
		tenantId: opts.Resource.Target,
	}
}
