package engine

import (
	"github.com/rs/zerolog"

	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/engine/runtime"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type Engine struct {
	app     xtypes.App
	runtime etypes.Runtime
	signer  service.Signer
	syncer  store.SyncDB

	pacman   service.Pacman
	logger   zerolog.Logger
	builders map[string]etypes.ExecutorBuilder
}

func New(_app xtypes.App, logger zerolog.Logger) *Engine {

	return &Engine{
		app:     _app,
		runtime: nil,
		signer:  nil,
		syncer:  nil,
		pacman:  nil,
		logger:  logger,
	}

}

func (e *Engine) Run() error {
	return e.run()
}

func (e *Engine) GetRuntime() etypes.Runtime {
	return e.runtime
}

func (e *Engine) Execute(opts etypes.Execution) ([]byte, error) {
	return e.execute(opts)
}

func (e *Engine) ServePlugFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return e.servePlugFile(tenantId, plugId, agentId, file)
}

func (e *Engine) ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return e.serveExecutorFile(tenantId, plugId, agentId, file)
}

// private

func (e *Engine) run() error {
	deps := e.app.GetDeps()

	e.runtime = runtime.New(e.app, e.logger)
	e.signer = deps.Signer().(service.Signer)
	e.syncer = deps.CoreHub().(store.SyncDB)
	e.pacman = deps.Pacman().(service.Pacman)

	reg := deps.Registry().(*registry.Registry)

	bfuncs := reg.GetExecutors()
	e.builders = make(map[string]etypes.ExecutorBuilder)

	for k, ebf := range bfuncs {
		bf, err := ebf(e.app)
		if err != nil {
			panic(err)
		}

		e.builders[k] = bf
	}

	mfuncs := reg.GetExecModules()
	modules := make(map[string]etypes.ModuleBuilder)

	for k, mbf := range mfuncs {
		bf, err := mbf(e.app)
		if err != nil {
			panic(err)
		}
		modules[k] = bf
	}

	return e.runtime.Run(e.builders, modules)
}
