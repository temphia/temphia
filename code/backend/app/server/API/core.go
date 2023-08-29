package api_server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (s *Server) X(fn func(ctx httpx.Request)) func(*gin.Context) {
	return s.middleware.LoggedX(fn)
}
