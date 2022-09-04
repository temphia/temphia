package dyndb

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type DyndbBuilder struct {
}

func (DyndbBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	return New(opts), nil
}

func (DyndbBuilder) Init(app interface{}) error {
	return nil
}

func New(opts etypes.ModuleOptions) *DyndbModule {
	deps := opts.Binder.GetApp().(xtypes.App).GetDeps()

	dynhub := deps.DynHub().(store.DynHub).GetSource("default", opts.Resource.TenantId) // fixme => get source from resource

	return &DyndbModule{
		binder: opts.Binder,
		res:    opts.Resource,
		dynsrc: dynhub,
		group:  "",
		table:  "",
	}
}
