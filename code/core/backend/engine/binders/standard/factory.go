package standard

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/core/backend/engine/binders/standard/deps"
	"github.com/temphia/temphia/code/core/backend/engine/binders/standard/handle"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type Factory struct {
	deps deps.Deps
}

func NewFactory(app xtypes.App, logger zerolog.Logger, modules map[string]etypes.ModuleBuilder) Factory {

	appdeps := app.GetDeps()

	return Factory{
		deps: deps.Deps{
			App:            app,
			Sockd:          appdeps.Sockd().(sockdx.SockdCore),
			Pacman:         appdeps.Pacman().(service.Pacman),
			Corehub:        appdeps.CoreHub().(store.CoreHub),
			CabinetHub:     appdeps.Cabinet().(store.CabinetHub),
			NodeCache:      appdeps.NodeCache().(service.NodeCache),
			PlugKV:         appdeps.PlugKV().(store.PlugStateKV),
			LoggerBase:     logger,
			ModuleBuilders: modules,
		},
	}
}

type BinderOptions struct {
	Namespace string
	PlugId    string
	AgentId   string
	BprintId  string
	Epoch     int64
}

func (bf Factory) New(opts BinderOptions) *Binder {

	handle := handle.New(opts.Namespace,
		opts.PlugId,
		opts.AgentId,
		opts.BprintId,
		&bf.deps,
	)

	b := &Binder{
		Handle:       &handle,
		ReuseCounter: -1,
		Epoch:        opts.Epoch,
	}

	return b
}
