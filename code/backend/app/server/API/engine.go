package api_server

import (
	_ "embed"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/controllers/sockd"
	"github.com/temphia/temphia/code/backend/services/sockd/transports"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (s *Server) EngineAPI(rg *gin.RouterGroup) {

	// launch
	rg.POST("/launch/target", s.X(s.launchTarget))
	rg.POST("/launch/agent", s.X(s.launchAgent))
	rg.POST("/launch/editor", s.X(s.launchEditor))

	rg.POST("/reset", s.X(s.reset))

	// rg.GET("/launch/domain_target/:pid/:aid", s.domainTargetLaunch())

	// execute action
	rg.POST("/execute/:action", s.execute)
	rg.OPTIONS("/execute/:action", func(ctx *gin.Context) {
		pp.Println("@iframe_cors")

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		ctx.AbortWithStatus(204)
	})

	// s.app.Data().AssetAdapter("build")
	fs := http.FS(nil) // fixme

	rg.GET("/plug/:pid/agent/:aid/launcher/:file", func(ctx *gin.Context) {
		ctx.FileFromFS(ctx.Param("file"), fs)
	})

	// serve file
	rg.GET("/plug/:pid/agent/:aid/serve/*file", s.agentServeFile)
	rg.GET("/plug/:pid/agent/:aid/executor/:eid/*file", s.executorFile)

	// engine sockd
	rg.GET("/ws", s.sockdRoomWS)
	rg.POST("/ws/update", func(ctx *gin.Context) {})
}

func (s *Server) execute(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	s.cEngine.Execute(ctx.Param("tenant_id"), ctx.Param("action"), ctx)
}

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

	resp, err := s.cEngine.LaunchAgent(ctx.Session, opts.PlugId, opts.AgentId)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) launchTarget(ctx httpx.Request) {
	opts := &launchOptions{}
	resp, err := s.cEngine.LaunchTarget(ctx.Session, opts.TargetId)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) launchEditor(ctx httpx.Request) {
	opts := &launchOptions{}
	resp, err := s.cEngine.LaunchEditor(ctx.Session, opts.PlugId, opts.AgentId)
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
		httpx.UnAuthorized(ctx.Http)
		return
	}

	s.cEngine.Reset(ctx.Session.TenantId, req.PlugId, req.AgentId)
	httpx.WriteFinal(ctx.Http, err)
}

// engine/exec sockd

func (s *Server) sockdRoomWS(ctx *gin.Context) {

	if !ctx.IsWebsocket() {
		return
	}

	tenantId := ctx.Param("tenant_id")

	dclaim, err := s.signer.ParseSockdTkt(tenantId, ctx.Query("ticket"))
	if err != nil {
		httpx.UnAuthorized(ctx)
		return
	}

	conn, err := transports.NewConnWS(ctx, s.idNode.Generate().Int64())
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cSockd.AddPlugConn(sockd.PlugConnOptions{
		TenantId: tenantId,
		UserId:   dclaim.UserId,
		GroupId:  "",
		DeviceId: dclaim.DeviceId,
		Plug:     "",
		Conn:     conn,
		Room:     dclaim.Room,
	})

	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

}
