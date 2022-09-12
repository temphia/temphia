package engine

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/engine/invokers/web"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
)

type Controller struct {
	engine etypes.Engine
	signer service.Signer
}

func New(engine etypes.Engine, signer service.Signer) *Controller {
	return &Controller{
		engine: engine,
		signer: signer,
	}
}

func (c *Controller) Execute(tenantId, action string, ctx *gin.Context) {

	payload, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return
	}

	eclaim, err := c.signer.ParseExecutor(tenantId, ctx.GetHeader("Authorization"))
	if err != nil {
		return
	}

	out, err := c.engine.Execute(etypes.Execution{
		TenantId: tenantId,
		PlugId:   eclaim.PlugId,
		AgentId:  eclaim.AgentId,
		Action:   action,
		Payload:  payload,
		Invoker:  web.NewWeb(ctx, eclaim),
	})
	if err != nil {
		return
	}

	ctx.Writer.Write(out)
}

func (c *Controller) ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return c.engine.ServeAgentFile(tenantId, plugId, agentId, file)
}

func (c *Controller) ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error) {
	return c.engine.ServeExecutorFile(tenantId, plugId, agentId, file)
}

func (c *Controller) LaunchAdmin(uclaim *claim.Session, data Admin) (string, error) {
	return c.launchAdmin(uclaim, data)
}

func (c *Controller) LaunchDev(uclaim *claim.Session, data Admin) (string, error) {
	//return c.launchAdmin(uclaim, data)
	return "", nil
}

func (c *Controller) LaunchData(uclaim *claim.Session, data Data) (string, error) {
	return c.launchData(uclaim, data)
}

func (c *Controller) LaunchUser(uclaim *claim.Session, data User) (string, error) {
	return c.launchUser(uclaim, data)
}

func (c *Controller) LaunchAuthd(uclaim *claim.Session, data Authd) (string, error) {
	return c.launchAuthd(uclaim, data)
}

func (c *Controller) LaunchWidget(uclaim *claim.Session, data Widget) (string, error) {
	return c.launchWidget(uclaim, data)
}

func (c *Controller) LaunchDomain(uclaim *claim.Session, data Domain) (string, error) {
	// DomainEditor
	return c.launchDomain(uclaim, data)

}
