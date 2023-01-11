package web

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/engine/invokers"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/invoker"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

type WebRequest struct {
	rctx  *gin.Context
	claim *claim.Executor
}

func NewWeb(ctx *gin.Context, eclaim *claim.Executor) *WebRequest {
	return &WebRequest{
		rctx:  ctx,
		claim: eclaim,
	}
}

func (r *WebRequest) Handle(method string, data xtypes.LazyData) (xtypes.LazyData, error) {

	if strings.HasPrefix(method, "http.") {
		return r.webHandle(method, data)
	} else if strings.HasPrefix(method, "data.") {
		return r.dataHandle(method, data)
	}

	return nil, invokers.ErrInvokerActionNotImplemented
}

func (r *WebRequest) Name() string {
	return r.claim.ExecType
}

func (r *WebRequest) User() *invoker.User {

	return &invoker.User{
		Id:        r.claim.UserId,
		Group:     r.claim.UserGroup,
		SessionId: r.claim.SessionId,
	}

}
