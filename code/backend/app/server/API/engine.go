package api_server

import (
	_ "embed"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (s *Server) EngineAPI(rg *gin.RouterGroup) {

	// launch
	rg.POST("/launch/target", s.X(s.launchTarget))
	rg.POST("/launch/agent", s.X(s.launchAgent))
	rg.POST("/launch/editor", s.X(s.launchEditor))

	rg.POST("/reset", s.X(s.reset))

	// execute action

	rg.POST("/rpx/closed/:action", s.executeClosedRPX)
	rg.POST("/rpx/open/:action", s.executeOpenRPX)
	rg.POST("/rpx/raw/:action", s.executeRawRPX)

	// serve file
	rg.GET("/plug/:pid/agent/:aid/serve/*file", s.agentServeFile)
	rg.GET("/plug/:pid/agent/:aid/executor/:eid/*file", s.executorFile)

}

// rpx execute

func (s *Server) executeOpenRPX(ctx *gin.Context) {

}

func (s *Server) executeClosedRPX(ctx *gin.Context) {
	s.cEngine.Execute(ctx.Param("tenant_id"), ctx.Param("action"), ctx)
}

func (s *Server) executeRawRPX(ctx *gin.Context) {

}

// serve

func (s *Server) agentServeFile(ctx *gin.Context) {
	file := strings.TrimPrefix(ctx.Param("file"), "/")

	out, err := s.cEngine.ServeAgentFile(ctx.Param("tenant_id"), ctx.Param("pid"), ctx.Param("aid"), file)
	if err != nil {
		pp.Println("@err/server_agent", file, err)
		return
	}

	httpx.WriteFile(file, out, ctx)
}

func (s *Server) executorFile(ctx *gin.Context) {
	file := strings.TrimPrefix(ctx.Param("file"), "/")

	out, err := s.cEngine.ServeExecutorFile(ctx.Param("tenant_id"), ctx.Param("pid"), ctx.Param("aid"), file)
	if err != nil {
		pp.Println("@err/executor_file", err)
		return
	}
	httpx.WriteFile(file, out, ctx)
}

// launch

type launchOptions struct {
	PlugId   string `json:"plug_id,omitempty"`
	AgentId  string `json:"agent_id,omitempty"`
	TargetId int64  `json:"target_id,omitempty"`
}

func (s *Server) launchAgent(ctx httpx.Request) {
	opts := &launchOptions{}

	err := ctx.Http.BindJSON(opts)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	resp, err := s.cEngine.LaunchAgent(ctx.Session, opts.PlugId, opts.AgentId)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) launchTarget(ctx httpx.Request) {
	opts := &launchOptions{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	resp, err := s.cEngine.LaunchTarget(ctx.Session, opts.TargetId)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) launchEditor(ctx httpx.Request) {
	opts := &launchOptions{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	resp, err := s.cEngine.LaunchEditor(ctx.Session, opts.PlugId, opts.AgentId)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	httpx.WriteJSON(ctx.Http, resp, err)
}

type ResetRequest struct {
	PlugId  string `json:"plug_id,omitempty"`
	AgentId string `json:"agent_id,omitempty"`
}

func (s *Server) reset(ctx httpx.Request) {
	req := ResetRequest{}

	err := ctx.Http.BindJSON(&req)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	s.cEngine.Reset(ctx.Session.TenantId, req.PlugId, req.AgentId)
	httpx.WriteFinal(ctx.Http, err)
}
