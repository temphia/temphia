package server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/controllers/engine"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (s *Server) engineAPI(rg *gin.RouterGroup) {

	rg.POST("/launch/target", s.X(s.launchTarget))
	rg.POST("/launch/admin", s.X(s.launchAdmin))
	rg.POST("/launch/auth", s.launchAuth)

	rg.POST("/execute/:action", s.execute)

	rg.GET("/plug/:pid/agent/:aid/serve/:file", s.agentServeFile)
	rg.GET("/plug/:pid/agent/:aid/executor/:eid/:file", s.executorFile)

}

func (s *Server) execute(ctx *gin.Context) {
	s.cEngine.Execute(ctx.Param("tenant_id"), ctx.Param("action"), ctx)
}

func (s *Server) agentServeFile(ctx *gin.Context) {
	s.cEngine.ServeAgentFile(ctx.Param("tenant_id"), ctx.Param("pid"), ctx.Param("aid"), ctx.Param("file"))
}

func (s *Server) executorFile(ctx *gin.Context) {
	s.cEngine.ServeExecutorFile(ctx.Param("tenant_id"), ctx.Param("pid"), ctx.Param("aid"), ctx.Param("file"))
}

// launch

func (s *Server) launchTarget(ctx httpx.Request) {

	data := engine.TargetLaunchData{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		return
	}

	out, err := s.cEngine.LaunchTarget(ctx.Session, data)
	out.BaseURL = httpx.BaseURL(ctx.Http.Request.Host, ctx.Session.TenentId)
	httpx.WriteJSON(ctx.Http, out, err)
}

func (s *Server) launchAdmin(ctx httpx.Request) {
	data := engine.AdminLaunchData{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		return
	}

	out, err := s.cEngine.LaunchAdmin(ctx.Session, data)
	httpx.WriteJSON(ctx.Http, out, err)

}

func (s *Server) launchAuth(ctx *gin.Context) {
	data := engine.AuthLaunchData{}

	err := ctx.BindJSON(data)
	if err != nil {
		return
	}

	out, err := s.cEngine.LaunchAuth(data)
	httpx.WriteJSON(ctx, out, err)

}

// engine/exec sockd

// func (s *Server) sockdRoomWS(ctx *gin.Context) {}
// func (s *Server) sockdRoomUpdateWS(ctx *gin.Context) {}
