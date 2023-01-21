package deps

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Deps struct {
	App            xtypes.App
	Corehub        store.CoreHub
	CabinetHub     store.CabinetHub
	Sockd          sockdx.SockdCore
	Pacman         repox.Hub
	LoggerBase     zerolog.Logger
	NodeCache      service.NodeCache
	PlugKV         store.PlugStateKV
	Runtime        etypes.Runtime
	ModuleBuilders map[string]etypes.ModuleBuilder
	ExecBuilders   map[string]etypes.ExecutorBuilder
	Signer         service.Signer
}
