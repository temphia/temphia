package pstate

import (
	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var _ etypes.ModuleBuilder = (*PStateBuilder)(nil)

func NewBuilder(app any) (etypes.ModuleBuilder, error) {
	return &PStateBuilder{}, nil
}

type PStateBuilder struct{}

func (p *PStateBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	return New(
		opts.Resource.TenantId,
		opts.Resource.Target,
		opts.Binder.GetApp().(xtypes.App),
	), nil
}

func New(tenantId, plugId string, app xtypes.App) *PStateMod {
	deps := app.GetDeps()

	ps := &PStateMod{
		tenantId: tenantId,
		plugId:   plugId,
		pkv:      deps.PlugKV().(store.PlugStateKV),
		modipc:   nil,
	}

	ps.modipc = modipc.NewModIPC(ps)

	return ps
}
