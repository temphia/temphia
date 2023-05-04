package usergroup

import (
	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var _ etypes.ModuleBuilder = (*UGBuilder)(nil)

func NewBuilder(app any) (etypes.ModuleBuilder, error) {
	return &UGBuilder{
		app: app.(xtypes.App),
	}, nil
}

type UGBuilder struct {
	app xtypes.App
}

func (u UGBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	umod := &UserGroupModule{
		coreHub:  u.app.GetDeps().CoreHub().(store.CoreHub),
		group:    opts.Resource.Target,
		tenantId: opts.Resource.TenantId,
		bindings: opts.Binder,
	}

	umod.modipc = modipc.NewModIPC(umod)

	return umod, nil
}
