package engine

import (
	"io"
	"net/http"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type Controller struct {
	engine  etypes.Engine
	signer  service.Signer
	corehub store.CoreHub
	idgen   *snowflake.Node
}

func New(engine etypes.Engine, signer service.Signer, corehub store.CoreHub) *Controller {

	idgen, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	return &Controller{
		engine:  engine,
		signer:  signer,
		corehub: corehub,
		idgen:   idgen,
	}
}

func (c *Controller) Execute(tenantId, action string, ctx *gin.Context) {

	payload, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	eclaim, err := c.signer.ParseExecutor(tenantId, ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	pp.Println("@here_payload_before", string(payload))

	out, err := c.engine.Execute(etypes.Execution{
		TenantId: tenantId,
		PlugId:   eclaim.PlugId,
		AgentId:  eclaim.AgentId,
		Action:   action,
		Payload:  payload,
		Invoker:  nil, //web.NewWeb(ctx, eclaim),
	})

	if err != nil {
		pp.Println("@here_err_after", string(payload))
		ctx.AbortWithError(http.StatusBadRequest, err)
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
