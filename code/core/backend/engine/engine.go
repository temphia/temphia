package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/engine/runtime"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type Engine struct {
	app        xtypes.App
	runtime    etypes.Runtime
	signer     service.Signer
	syncer     store.SyncDB
	AssetStore xtypes.DataBox
	pacman     service.Pacman
	logger     zerolog.Logger
	builders   map[string]etypes.ExecutorBuilder
}

func New(_app xtypes.App, logger zerolog.Logger) *Engine {

	return &Engine{
		app:        _app,
		runtime:    nil,
		signer:     nil,
		syncer:     nil,
		AssetStore: _app.Data(),
		pacman:     nil,
		logger:     logger,
	}

}

func (e *Engine) Run() error {

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

func (e *Engine) GetRuntime() etypes.Runtime {
	return e.runtime
}

func (e *Engine) ServerLaunchExec(tenantId, plugId, agentId, mode string, arg interface{}, resp interface{}) error {
	return e.serverLaunchExec(tenantId, plugId, agentId, mode, arg, resp)
}

func (e *Engine) ClientLaunchExec(tenantId, plugId, agentId, mode string, ctx *gin.Context) {
	e.clientLaunchExec(tenantId, plugId, agentId, mode, ctx)
}

func (e *Engine) ExecAction(tenantId, plugId, agentId, action string, ctx *gin.Context) {
	e.plugAction(tenantId, plugId, agentId, action, ctx)
}

func (e *Engine) ServePlugFile(tenantId, plugId, agentId, file string, ctx *gin.Context) {
	e.servePlugFile(tenantId, plugId, agentId, file, ctx)
}

func (e *Engine) ServeExecutorFile(tenantId, plugId, agentId, loader string, ctx *gin.Context) {
	e.serveExecutorFile(tenantId, plugId, agentId, loader, ctx)
}
