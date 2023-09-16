package lsock

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func (l *LSock) apiRPCAction(ctx *gin.Context) {

	lclaim, err := l.signer.ParseLSock(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	l.sLock.Lock()
	sub := l.subs[lclaim.SID]
	l.sLock.Unlock()

	out, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sub.Handle(ctx.Param("action"), out)
}

func (l *LSock) apiWS(ctx *gin.Context) {
	if !ctx.IsWebsocket() {
		return
	}

	tok, _ := ctx.GetQuery("token")
	lclaim, err := l.signer.ParseLSock(tok)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	l.sLock.Lock()
	sub := l.subs[lclaim.SID]
	l.sLock.Unlock()

	conn, _, _, err := ws.UpgradeHTTP(ctx.Request, ctx.Writer)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	lws := &lsockWs{
		parent: l,
		conn:   conn,
		sub:    sub,
		wChan:  make(chan []byte),
	}

	l.wLock.Lock()
	l.wsconns[lclaim.SID] = lws
	l.wLock.Unlock()

	go l.handleWS(lws)

}

func (l *LSock) apiRegister(ctx *gin.Context) {
	lclaim, err := l.signer.ParseLSock(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	out, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	l.notz.RegisterLocalAddr(lclaim.Plug, lclaim.Agent, string(out))
}

// websocket

func (l *LSock) handleWS(lws *lsockWs) {

	go func() {

		for {
			msg, op, err := wsutil.ReadClientData(lws.conn)
			if err != nil {
				return
			}

			if op == ws.OpText {
				lws.sub.HandleWS(msg)
				continue
			}
		}

	}()

	go func() {

		for {
			err := wsutil.WriteServerMessage(lws.conn, ws.OpText, <-lws.wChan)
			if err != nil {
				return
			}
		}

	}()

}
