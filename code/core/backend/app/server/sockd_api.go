package server

import "github.com/gin-gonic/gin"

func (s *Server) sockdAPI(rg *gin.RouterGroup) {

	rg.GET("/user/ws", func(ctx *gin.Context) {})
	rg.GET("/data/ws", func(ctx *gin.Context) {})
	rg.GET("/data/update", func(ctx *gin.Context) {})

	rg.GET("/room/ws", func(ctx *gin.Context) {})
	rg.GET("/room/update", func(ctx *gin.Context) {})

}
