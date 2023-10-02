package enginehub

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"

	//	"github.com/temphia/temphia/code/backend/engine/invokers/bundled"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (e *EngineHub) execute(tenantId, action string, ctx *gin.Context) {

	payload, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	eclaim, err := e.signer.ParseExecutor(tenantId, ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	pp.Println("@here_payload_before", string(payload))

	out, err := e.engine.RPXecute(etypes.Execution{
		TenantId: tenantId,
		PlugId:   eclaim.PlugId,
		AgentId:  eclaim.AgentId,
		Action:   action,
		Payload:  payload,
		//		Invoker:  bundled.NewWeb(ctx, nil, eclaim),
	})

	if err != nil {
		pp.Println("@here_err_after", string(payload))
		fmt.Println("@exec_err", err)

		ctx.Writer.WriteHeader(http.StatusBadRequest)
		ctx.Writer.WriteString(err.Error())
		return
	}

	ctx.Writer.Write(out)
}

func (e *EngineHub) executeDev(dclaim *claim.UserContext, plug, agent, action string, body []byte) ([]byte, error) {
	// fixme => ability to send arbitary invoker type or dev invoker type?

	return e.engine.RPXecute(etypes.Execution{
		TenantId: dclaim.TenantId,
		PlugId:   plug,
		AgentId:  agent,
		Action:   action,
		Payload:  body,
		//		Invoker:  bundled.NewAdmin(dclaim),
	})
}
