package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/controllers/engine"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

func (s *Server) engineAPI(rg *gin.RouterGroup) {

	// launch
	rg.POST("/launch/target", s.X(s.launchTarget))
	rg.POST("/launch/admin", s.X(s.launchAdmin))
	rg.POST("/launch/auth", s.launchAuth)

	// execute action
	rg.POST("/execute/:action", s.execute)

	fs := http.FS(s.app.Data().AssetAdapter("build"))
	rg.GET("/plug/:pid/agent/:aid/launcher/:file", func(ctx *gin.Context) {
		ctx.FileFromFS(ctx.Param("file"), fs)
	})

	// serve file
	rg.GET("/plug/:pid/agent/:aid/serve/:file", s.agentServeFile)
	rg.GET("/plug/:pid/agent/:aid/executor/:eid/:file", s.executorFile)

	// engine sockd
	rg.GET("/ws", func(ctx *gin.Context) {})
	rg.POST("/ws/update", func(ctx *gin.Context) {})
}

func (s *Server) execute(ctx *gin.Context) {
	s.cEngine.Execute(ctx.Param("tenant_id"), ctx.Param("action"), ctx)
}

func (s *Server) agentServeFile(ctx *gin.Context) {

	out, err := s.cEngine.ServeAgentFile(ctx.Param("tenant_id"), ctx.Param("pid"), ctx.Param("aid"), ctx.Param("file"))
	if err != nil {
		pp.Println("@err/server_agent", err)
		return
	}

	httpx.WriteFile(ctx.Param("file"), out, ctx)
}

func (s *Server) executorFile(ctx *gin.Context) {
	out, err := s.cEngine.ServeExecutorFile(ctx.Param("tenant_id"), ctx.Param("pid"), ctx.Param("aid"), ctx.Param("file"))
	if err != nil {
		pp.Println("@err/executor_file", err)
		return
	}
	httpx.WriteFile(ctx.Param("file"), out, ctx)
}

// launch

func (s *Server) launchTarget(ctx httpx.Request) {

	data := engine.TargetLaunchData{}

	err := ctx.Http.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	out, err := s.cEngine.LaunchTarget(ctx.Session, data)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	out.BaseURL = httpx.BaseURL(ctx.Http.Request.Host, ctx.Session.TenantId)
	httpx.WriteJSON(ctx.Http, out, err)
}

func (s *Server) launchAdmin(ctx httpx.Request) {
	data := engine.AdminLaunchData{}

	err := ctx.Http.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	out, err := s.cEngine.LaunchAdmin(ctx.Session, data)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	out.BaseURL = httpx.BaseURL(ctx.Http.Request.Host, ctx.Session.TenantId)
	httpx.WriteJSON(ctx.Http, out, err)

}

func (s *Server) launchAuth(ctx *gin.Context) {
	data := engine.AuthLaunchData{}

	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	out, err := s.cEngine.LaunchAuth(data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteJSON(ctx, out, err)

}

// engine/exec sockd

// func (s *Server) sockdRoomWS(ctx *gin.Context) {}
// func (s *Server) sockdRoomUpdateWS(ctx *gin.Context) {}
