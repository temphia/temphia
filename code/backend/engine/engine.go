package engine

import (
	"github.com/rs/zerolog"

	"github.com/temphia/temphia/code/backend/engine/runtime"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var (
	_ etypes.Engine = (*Engine)(nil)
)

type Engine struct {
	app     xtypes.App
	runtime etypes.Runtime
	signer  service.Signer
	syncer  store.SyncDB

	pacman       repox.Hub
	logger       zerolog.Logger
	execbuilders map[string]etypes.ExecutorBuilder
	modBuilders  map[string]etypes.ModuleBuilder
}

func New(_app xtypes.App, logger zerolog.Logger) *Engine {

	return &Engine{
		app:          _app,
		runtime:      nil,
		signer:       nil,
		syncer:       nil,
		pacman:       nil,
		execbuilders: nil,
		modBuilders:  nil,
		logger:       logger,
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

func (e *Engine) ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return e.serveAgentFile(tenantId, plugId, agentId, file)
}

func (e *Engine) ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return e.serveExecutorFile(tenantId, plugId, agentId, file)
}

func (e *Engine) ListExecutors() []string {

	keys := make([]string, 0, len(e.execbuilders))
	for k := range e.execbuilders {
		keys = append(keys, k)
	}

	return keys
}

func (e *Engine) ListModules() []string {

	keys := make([]string, 0, len(e.modBuilders))
	for k := range e.modBuilders {
		keys = append(keys, k)
	}

	return keys
}

// private

func (e *Engine) run() error {
	deps := e.app.GetDeps()

	e.runtime = runtime.New(e.app, e.logger)
	e.signer = deps.Signer().(service.Signer)
	e.syncer = deps.CoreHub().(store.SyncDB)
	e.pacman = deps.RepoHub().(repox.Hub)

	e.execbuilders = e.app.GetGlobalVar().Get("executors").(map[string]etypes.ExecutorBuilder)
	e.modBuilders = e.app.GetGlobalVar().Get("modules").(map[string]etypes.ModuleBuilder)

	return e.runtime.Run(e.execbuilders, e.modBuilders)
}
