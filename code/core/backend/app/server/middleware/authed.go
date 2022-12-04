package middleware

import (
	"reflect"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

func (m *Middleware) LoggedX(fn func(ctx httpx.Request)) func(*gin.Context) {

	return func(c *gin.Context) {

		// time.Sleep(time.Duration(rand.Int()%5) * time.Second)

		tenantId := c.Param("tenant_id")
		sessToken := c.GetHeader("Authorization")
		sclaim, err := m.Signer.ParseSession(tenantId, sessToken)
		if err != nil {
			m.Logger.Error().
				Err(err).
				Str("tenant_id", tenantId).
				Msg(logid.RoutesSessionParseErr)
			return
		}

		if sclaim.Type != claim.CTypeSession {
			m.Logger.Warn().
				Str("tenant_id", tenantId).
				Str("client_ip", c.ClientIP()).
				Interface("data", sclaim).
				Msg(logid.RoutesWrongSessionClaim)
		}

		sclaim.TenentId = tenantId

		c.Header("X-Clacks-Overhead", "Aaron Swartz")

		m.Logger.Info().
			Str("tenant_id", tenantId).
			Str("user_group", sclaim.UserGroup).
			Str("user_id", sclaim.UserID).
			Str("device_id", sclaim.DeviceId).
			Int64("session_id", sclaim.SessionID).
			Str("handler_method", funcName(fn)).Msg(logid.RoutesSessionParsed)

		fn(httpx.Request{
			Id:      0,
			Http:    c,
			Session: sclaim,
		})
	}
}

func (m *Middleware) DataX(fn func(uclaim *claim.Data, ctx *gin.Context)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		uclaim, err := m.Signer.ParseData(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
		if err != nil {
			return
		}

		fn(uclaim, ctx)
	}
}

func (m *Middleware) FolderX(fn func(uclaim *claim.Folder, ctx *gin.Context)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		uclaim, err := m.Signer.ParseFolder(ctx.Param("tenant_id"), ctx.Param("ticket"))
		if err != nil {
			return
		}

		fn(uclaim, ctx)

	}
}

func (m *Middleware) ExecutorX(fn func(uclaim *claim.Executor, ctx *gin.Context)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		uclaim, err := m.Signer.ParseExecutor(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
		if err != nil {
			return
		}

		fn(uclaim, ctx)
	}
}

// private

func funcName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
