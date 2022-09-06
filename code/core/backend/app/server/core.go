package server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (s *Server) X(fn func(ctx httpx.Request)) func(*gin.Context) {
	return s.middleware.Authed(fn)
}

func (s *Server) asFile(data []byte, ext string) func(ctx *gin.Context) {
	return s.middleware.AsFile(data, ext)
}
