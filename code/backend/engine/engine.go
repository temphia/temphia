package engine2

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/engine/eutils/ecache"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var _ etypes.Engine = (*Engine)(nil)

type Engine struct {
	app    xtypes.App
	signer service.Signer
	syncer store.SyncDB

	ecache etypes.Ecache

	pacman       xpacman.Pacman
	logger       zerolog.Logger
	execbuilders map[string]etypes.ExecutorBuilder
	modBuilders  map[string]etypes.ModuleBuilder
}

func New(_app xtypes.App, logger zerolog.Logger) *Engine {

	return &Engine{
		app:          _app,
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

func (e *Engine) GetCache() etypes.Ecache {

	return nil
}
func (e *Engine) RPXecute(options etypes.Execution) ([]byte, error) {
	return nil, nil
}
func (e *Engine) WebRawXecute(rw http.ResponseWriter, req *http.Request) {

}

func (e *Engine) SetRemoteOption(opt etypes.RemoteOptions) {

}

func (e *Engine) ResetAgent(tenantId, plugId, agentId string) error { return nil }
func (e *Engine) ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return nil, nil
}

func (e *Engine) ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return nil, nil
}

func (e *Engine) RemotePerform(opt etypes.Remote) ([]byte, error) {
	return nil, nil
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

	ch := deps.CoreHub().(store.CoreHub)

	e.signer = deps.Signer().(service.Signer)
	e.syncer = ch
	e.pacman = deps.RepoHub().(xpacman.Pacman)

	e.execbuilders = e.app.GetGlobalVar().Get("executors").(map[string]etypes.ExecutorBuilder)
	e.modBuilders = e.app.GetGlobalVar().Get("modules").(map[string]etypes.ModuleBuilder)

	e.ecache = ecache.New(ch)

	// e.runtime.Run(e.execbuilders, e.modBuilders)

	return nil
}
