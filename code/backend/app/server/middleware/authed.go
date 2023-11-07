package middleware

import (
	"reflect"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
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

			httpx.UnAuthorized(c)
			return
		}

		if sclaim.Type != claim.CTypeSession {
			m.Logger.Warn().
				Str("tenant_id", tenantId).
				Str("client_ip", c.ClientIP()).
				Interface("data", sclaim).
				Msg(logid.RoutesWrongSessionClaim)
		}

		sclaim.TenantId = tenantId

		c.Header("X-Clacks-Overhead", "Aaron Swartz")

		m.Logger.Info().
			Str("tenant_id", tenantId).
			Str("user_group", sclaim.UserGroup).
			Str("user_id", sclaim.UserID).
			Int64("device_id", sclaim.DeviceId).
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
			httpx.UnAuthorized(ctx)
			return
		}

		fn(uclaim, ctx)
	}
}

func (m *Middleware) FolderX(fn func(uclaim *claim.Folder, ctx *gin.Context)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		uclaim, err := m.Signer.ParseFolder(ctx.Param("tenant_id"), ctx.Param("ticket"))
		if err != nil {
			httpx.UnAuthorized(ctx)
			return
		}

		fn(uclaim, ctx)

	}
}

func (m *Middleware) ExecutorX(fn func(uclaim *claim.Executor, ctx *gin.Context)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		uclaim, err := m.Signer.ParseExecutor(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
		if err != nil {
			httpx.UnAuthorized(ctx)
			return
		}

		fn(uclaim, ctx)
	}
}

func (m *Middleware) PSX(fn func(aclaim *claim.PlugState, ctx *gin.Context)) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		aclaim, err := m.Signer.ParsePlugState(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
		if err != nil {
			httpx.UnAuthorized(ctx)
			return
		}

		fn(aclaim, ctx)
	}
}

func (m *Middleware) UX(fn func(uclaim *claim.UserMgmtTkt, http *gin.Context)) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		umClaim, err := m.Signer.ParseUserMgmtTkt(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
		if err != nil {
			httpx.UnAuthorized(ctx)
			return
		}

		fn(umClaim, ctx)
	}
}

func (m *Middleware) BprintX(fn func(uclaim *claim.BprintTkt, http *gin.Context)) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		umClaim, err := m.Signer.ParseBprintTkt(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
		if err != nil {
			httpx.UnAuthorized(ctx)
			return
		}

		fn(umClaim, ctx)
	}
}

// private

func funcName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
