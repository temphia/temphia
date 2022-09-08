package server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/controllers/engine"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (s *Server) engineAPI(rg *gin.RouterGroup) {

	rg.POST("/launch/data", s.X(s.launchData))
	rg.POST("/launch/user", s.X(s.launchUser))
	rg.POST("/launch/admin", s.X(s.launchAdmin))
	rg.POST("/launch/domain", s.X(s.launchDomain))

	rg.POST("/launch/authd", s.launchAuthd)
	rg.POST("/launch/widget", s.launchWidget)

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

func (s *Server) launchData(ctx httpx.Request) {
	data := engine.Data{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		return
	}

	out, err := s.cEngine.LaunchData(ctx.Session, data)
	httpx.WriteJSON(ctx.Http, out, err)
}

func (s *Server) launchUser(ctx httpx.Request) {

	data := engine.Data{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		return
	}

	out, err := s.cEngine.LaunchData(ctx.Session, data)
	httpx.WriteJSON(ctx.Http, out, err)
}

func (s *Server) launchAdmin(ctx httpx.Request) {
	data := engine.Admin{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		return
	}

	out, err := s.cEngine.LaunchAdmin(ctx.Session, data)
	httpx.WriteJSON(ctx.Http, out, err)
}

func (s *Server) launchDomain(ctx httpx.Request) {
	data := engine.Domain{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		return
	}

	out, err := s.cEngine.LaunchDomain(ctx.Session, data)
	httpx.WriteJSON(ctx.Http, out, err)
}

func (s *Server) launchAuthd(ctx *gin.Context)  {}
func (s *Server) launchWidget(ctx *gin.Context) {}
