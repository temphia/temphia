package registry

import (
	"github.com/temphia/temphia/code/backend/app/adapter"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/store"
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

func SetStoreBuilders(name string, builder store.Builder) {
	G.SetStoreBuilder(name, builder)
}

func SetExecutor(name string, builder etypes.BuilderFactory) {
	G.SetExecutor(name, builder)
}

func SetExecModule(name string, builder etypes.ModuleBuilderFunc) {
	G.SetExecModule(name, builder)
}

func SetAdapterBuilder(name string, rb adapter.Builder) {
	G.SetAapterBuilder(name, rb)
}
