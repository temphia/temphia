package plugkv

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
	app := opts.Binder.GetApp().(xtypes.App)
	deps := app.GetDeps()

	return New(
		opts.Resource.TenantId,
		opts.Resource.Target,
		deps.PlugKV().(store.PlugStateKV),
	), nil
}

func New(tenantId, plugId string, kv store.PlugStateKV) *PlugStateMod {

	ps := &PlugStateMod{
		tenantId: tenantId,
		plugId:   plugId,
		pkv:      kv,
		modipc:   nil,
	}

	ps.modipc = modipc.NewModIPC(ps)

	return ps
}
