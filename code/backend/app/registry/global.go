package registry

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
)

var G *Registry

func init() {
	if G == nil {
		G = New(false)
	}

}

func SetRepoBuilder(name string, builder repox.Builder) {
	G.SetRepoBuilder(name, builder)
}

func SetStoreBuilders(name string, builder StoreBuilder) {
	G.SetStoreBuilder(name, builder)
}

func SetExecutor(name string, builder etypes.BuilderFactory) {
	G.SetExecutor(name, builder)
}

func SetExecModule(name string, builder etypes.ModuleBuilderFunc) {
	G.SetExecModule(name, builder)
}

func SetDynamicScript(name string, script func(ns string, ctx any) error) {
	G.SetDynamicScript(name, script)
}

func SetAdapterBuilder(name string, rb httpx.Builder) {
	G.SetAapterBuilder(name, rb)
}
