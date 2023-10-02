package enginehub

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	engine "github.com/temphia/temphia/code/backend/engine"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/launch"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

var (
	_ etypes.EngineHub = (*EngineHub)(nil)
)

type EngineHub struct {
	engine  etypes.Engine
	signer  service.Signer
	corehub store.CoreHub
	idgen   *snowflake.Node
	logger  *zerolog.Logger

	app xtypes.App
}

func New(app xtypes.App, logsvc logx.Service) *EngineHub {

	engine := engine.New(app, *logsvc.GetEngineLogger())

	return &EngineHub{
		engine: engine,
		app:    app,
	}

}

func (e *EngineHub) Start() error {
	deps := e.app.GetDeps()

	logger := deps.LogService().(logx.Service).GetServiceLogger("enginehub")

	e.signer = deps.Signer().(service.Signer)
	e.corehub = deps.CoreHub().(store.CoreHub)
	e.idgen = deps.ControlPlane().(xplane.ControlPlane).GetIdService().NewNode("engine")
	e.logger = &logger

	return e.engine.Run()
}

func (e *EngineHub) GetEngine() etypes.Engine {
	return e.engine
}

func (e *EngineHub) GetCache() etypes.Ecache {
	return e.engine.GetCache()
}

func (e *EngineHub) GetExecutorBuilder(name string) etypes.ExecutorBuilder {
	// fixme => impl
	return nil
}

func (e *EngineHub) LaunchAgent(uclaim *claim.Session, plugId, agentId string) (*launch.Response, error) {
	return e.launchAgent(uclaim, plugId, agentId)
}

func (e *EngineHub) LaunchTarget(uclaim *claim.Session, targetId int64) (*launch.Response, error) {
	return e.launchTarget(uclaim, targetId)
}

func (e *EngineHub) LaunchEditor(uclaim *claim.Session, plugId, agentId string) (*launch.Response, error) {
	return e.launchEditor(uclaim, plugId, agentId)
}

func (e *EngineHub) Execute(tenantId, action string, ctx *gin.Context) {
	e.execute(tenantId, action, ctx)
}

func (e *EngineHub) ExecuteDev(dclaim *claim.UserContext, plug, agent, action string, body []byte) ([]byte, error) {
	return e.executeDev(dclaim, plug, agent, action, body)
}

func (e *EngineHub) Reset(tenantId, plugId, agentId string) error {
	// e.engine.GetRuntime().ResetAgents(tenantId, plugId, []string{agentId})
	return nil
}

func (e *EngineHub) ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return e.engine.ServeAgentFile(tenantId, plugId, agentId, file)
}

func (e *EngineHub) ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return e.engine.ServeExecutorFile(tenantId, plugId, agentId, file)
}

func (e *EngineHub) ListExecutors() ([]string, error) {
	execs := e.engine.ListExecutors()
	return execs, nil
}

func (e *EngineHub) ListModules() ([]string, error) {
	mods := e.engine.ListModules()
	return mods, nil
}

func (e *EngineHub) RunStartupHooks(tenants []string, minwait time.Duration) {
	e.runStartupHooks(tenants, minwait)
}

// func (e *EngineHub) RunDataTableHook() {}
// func (e *EngineHub) RunDataTableSheet() {}
// func (e *EngineHub) RunUserEvent() {}
