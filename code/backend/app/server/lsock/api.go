package lsock

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/xserver"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (l *LSock) apiRPCAction(ctx *gin.Context) {

	lclaim, err := l.signer.ParseLSock(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	l.sLock.Lock()
	sub := l.subs[lclaim.IID]
	l.sLock.Unlock()

	out, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	out, err = sub.Handle(ctx.Param("action"), out)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.Data(http.StatusOK, httpx.CtypeJSON, out)
}

func (l *LSock) apiRegister(ctx *gin.Context) {
	lclaim, err := l.signer.ParseLSock(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	info := &xserver.LSOptions{}
	err = ctx.BindJSON(info)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	info.AgentId = lclaim.Agent
	info.PlugId = lclaim.Plug

	l.rLock.Lock()
	l.remotes[lclaim.IID] = info
	l.rLock.Unlock()

}
