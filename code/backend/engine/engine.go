package engine

import (
	"path"
	"sync"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/engine/binder"
	"github.com/temphia/temphia/code/backend/engine/eutils/ecache"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
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

	// runtime
	running map[string]*binder.Binder
	rLock   sync.RWMutex
	rundb   runDB

	binderFactory binder.Factory
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
		running:      make(map[string]*binder.Binder),
		rLock:        sync.RWMutex{},
	}

}

func (e *Engine) Run() error {

	err := e.run()
	if err != nil {
		return err
	}

	e.binderFactory = binder.NewFactory(binder.FactoryOptions{
		App:          e.app,
		Logger:       e.logger,
		Modules:      e.modBuilders,
		ExecBuilders: e.execbuilders,
	})

	cd := e.app.GetDeps().Confd().(config.Confd)
	e.rundb = newRunDB(path.Join(cd.RootDataFolder(), "running.json"))

	return nil
}

func (e *Engine) GetCache() etypes.Ecache {

	return e.ecache
}

func (e *Engine) RPXecute(opts etypes.RPXecuteOptions) ([]byte, error) {
	return e.rPXecute(opts)
}

func (e *Engine) WebRawXecute(opts etypes.WebRawXecuteOptions) {
	e.webRawXecute(opts)
}

func (e *Engine) SetRemoteOption(opts etypes.RemoteOptions) {
	eb := e.execbuilders[opts.PlugId+opts.AgentId]
	eb.SetRemoteOptions(opts)
}

func (e *Engine) ResetAgent(tenantId, plugId, agentId string) error {

	return nil
}

func (e *Engine) ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	agent := e.ecache.GetAgent(tenantId, plugId, agentId)
	plug := e.ecache.GetPlug(tenantId, plugId)
	actualFile := agent.WebFiles[file]

	bstore := e.pacman.GetBprintFileStore()
	out, err := bstore.GetBlob(tenantId, plug.BprintId, "", actualFile)

	if err != nil {
		e.logger.Debug().
			Str("tenant_id", tenantId).
			Str("plug_id", plugId).
			Str("agent_id", agentId).
			Err(err).
			Str("file", file).
			Msg(logid.EngineServeBprintErr)
		return nil, err
	}

	return out, err

}

func (e *Engine) ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error) {

	agent := e.ecache.GetAgent(tenantId, plugId, agentId)
	eb := e.execbuilders[agent.Executor]

	return eb.ServeFile(file)
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
