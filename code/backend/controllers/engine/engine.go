package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/launch"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/thoas/go-funk"
)

type Controller struct {
	enginehub etypes.EngineHub
	corehub   store.CoreHub
}

func New(enginehub etypes.EngineHub, corehub store.CoreHub) *Controller {

	return &Controller{
		enginehub: enginehub,
		corehub:   corehub,
	}
}

func (c *Controller) Execute(tenantId, action string, ctx *gin.Context) {
	c.enginehub.Execute(tenantId, action, ctx)
}

func (c *Controller) Reset(tenantId, plugId, agentId string) error {
	return c.enginehub.Reset(tenantId, plugId, agentId)
}

func (c *Controller) ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return c.enginehub.ServeAgentFile(tenantId, plugId, agentId, file)
}

func (c *Controller) ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return c.enginehub.ServeExecutorFile(tenantId, plugId, agentId, file)
}

func (c *Controller) ListExecutors(uclaim *claim.Session) ([]string, error) {
	// fixme => check perm?
	return c.enginehub.ListExecutors()

}

func (c *Controller) ListModules(uclaim *claim.Session) ([]string, error) {
	// fixme => check perm?
	return c.enginehub.ListModules()
}

// launch stuff

func (c *Controller) LaunchAgent(uclaim *claim.Session, plugId, agentId string) (*launch.Response, error) {
	return c.enginehub.LaunchAgent(uclaim, plugId, agentId)
}

func (c *Controller) LaunchTarget(uclaim *claim.Session, targetId int64) (*launch.Response, error) {
	return c.enginehub.LaunchTarget(uclaim, targetId)
}

func (c *Controller) LaunchEditor(uclaim *claim.Session, plugId, agentId string) (*launch.Response, error) {
	return c.enginehub.LaunchEditor(uclaim, plugId, agentId)
}

func (c *Controller) ExecuteDev(dclaim *claim.PlugDevTkt, plug, agent, action string, body []byte) ([]byte, error) {

	if !dclaim.AllPlugs && !funk.ContainsString(dclaim.PlugIds, plug) {
		return nil, easyerr.NotAuthorized()
	}

	return c.enginehub.ExecuteDev(&claim.UserContext{
		TenantId:  dclaim.TenantId,
		UserID:    dclaim.UserId,
		UserGroup: dclaim.UserGroup,
	}, plug, agent, action, body)

}

func (c *Controller) GetEngine() etypes.Engine {
	return c.enginehub.GetEngine()
}

func (c *Controller) RegisterLocalAddr(opts etypes.RemoteOptions) {
	c.enginehub.GetEngine().SetRemoteOption(opts)
}
