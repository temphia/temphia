package binder

import (
	"github.com/rs/zerolog"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type FactoryOptions struct {
	App          xtypes.App
	Logger       zerolog.Logger
	Modules      map[string]etypes.ModuleBuilder
	ExecBuilders map[string]etypes.ExecutorBuilder
	Runtime      etypes.Runtime
}

type Factory struct {
	App            xtypes.App
	Corehub        store.CoreHub
	CabinetHub     store.CabinetHub
	Sockd          sockdx.SockdCore
	Pacman         xpacman.Pacman
	LoggerBase     zerolog.Logger
	NodeCache      service.NodeCache
	PlugKV         store.PlugStateKV
	Runtime        etypes.Runtime
	ModuleBuilders map[string]etypes.ModuleBuilder
	ExecBuilders   map[string]etypes.ExecutorBuilder
	Signer         service.Signer
}

func NewFactory(opts FactoryOptions) Factory {

	appdeps := opts.App.GetDeps()

	return Factory{

		App:            opts.App,
		Sockd:          appdeps.SockdHub().(sockdx.Hub).GetSockd(),
		Pacman:         appdeps.RepoHub().(xpacman.Pacman),
		Corehub:        appdeps.CoreHub().(store.CoreHub),
		CabinetHub:     appdeps.Cabinet().(store.CabinetHub),
		PlugKV:         appdeps.PlugKV().(store.PlugStateKV),
		Runtime:        opts.Runtime,
		ExecBuilders:   opts.ExecBuilders,
		LoggerBase:     opts.Logger,
		ModuleBuilders: opts.Modules,
		Signer:         appdeps.Signer().(service.Signer),
	}
}

type BinderOptions struct {
	Namespace string
	PlugId    string
	AgentId   string
	BprintId  string
	Epoch     int64
}

func (bf *Factory) New(opts BinderOptions) *Binder {

	b := &Binder{
		Deps: bf,

		Namespace: opts.Namespace,
		PlugId:    opts.PlugId,
		AgentId:   opts.AgentId,
		BprintId:  opts.BprintId,

		ReuseCounter: -1,
		Epoch:        opts.Epoch,
	}

	return b
}
