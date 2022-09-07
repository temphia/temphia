package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) engineAPI(rg *gin.RouterGroup) {

	rg.POST("/launch/data", s.launchData)
	rg.POST("/launch/user", s.launchUser)
	rg.POST("/launch/authd", s.launchAuthd)
	rg.POST("/launch/widget", s.launchWidget)
	rg.POST("/launch/admin", s.launchAdmin)
	rg.GET("/launch/domain", s.launchDomain)

	rg.GET("/plug/:pid/agent/:aid/serve/:file", s.agentServeFile)
	rg.GET("/plug/:pid/agent/:aid/executor/:eid/:file", s.executorFile)
	rg.POST("/plug/:pid/agent/:aid/execute/:action", s.execute)

}

func (s *Server) execute(ctx *gin.Context)        {}
func (s *Server) agentServeFile(ctx *gin.Context) {}
func (s *Server) executorFile(ctx *gin.Context)   {}

func (s *Server) launchData(ctx *gin.Context)   {}
func (s *Server) launchUser(ctx *gin.Context)   {}
func (s *Server) launchAuthd(ctx *gin.Context)  {}
func (s *Server) launchWidget(ctx *gin.Context) {}
func (s *Server) launchAdmin(ctx *gin.Context)  {}
func (s *Server) launchDomain(ctx *gin.Context) {}
