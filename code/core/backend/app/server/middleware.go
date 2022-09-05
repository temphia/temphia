package server

import (
	"mime"
	"reflect"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

// X is a auth middleware
func (s *Server) X(fn func(ctx httpx.Request)) func(*gin.Context) {

	logger := s.log.GetServiceLogger("routes")

	return func(c *gin.Context) {

		// time.Sleep(time.Duration(rand.Int()%5) * time.Second)

		tenantId := c.Param("tenant_id")
		sessToken := c.GetHeader("Authorization")
		sclaim, err := s.signer.ParseSession(tenantId, sessToken)
		if err != nil {
			logger.Error().
				Err(err).
				Str("tenant_id", tenantId).
				Msg(logid.RoutesSessionParseErr)
			return
		}

		if sclaim.Type != claim.CTypeSession {
			logger.Warn().
				Str("tenant_id", tenantId).
				Str("client_ip", c.ClientIP()).
				Interface("data", sclaim).
				Msg(logid.RoutesWrongSessionClaim)
		}

		sclaim.TenentId = tenantId

		c.Header("X-Clacks-Overhead", "Aaron Swartz")

		logger.Info().
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

func (s *Server) asFile(data []byte, ext string) func(ctx *gin.Context) {
	exmime := mime.TypeByExtension(ext)
	clen := strconv.FormatInt(int64(len(data)), 10)

	// fixme => also set etag ?

	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", exmime)
		ctx.Header("Cache-Control", `public, max-age=86400`)
		ctx.Header("Content-Length", clen)
		ctx.Writer.Write(data)
	}
}

// private

func funcName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
