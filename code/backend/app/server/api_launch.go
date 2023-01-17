package server

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func (s *Server) launcher(rg *gin.RouterGroup) {
	// fixme => x-content-security-policy: frame-ancestors 'self' https://mycourses.w3schools.com;
	// Referer: https://example/launcher/<ticket>

	rg.GET("/:launch_token", func(ctx *gin.Context) {
		pp.Println(ctx.Param("launch_token"))
	})

}
