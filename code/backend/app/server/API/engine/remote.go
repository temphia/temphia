package engineapi

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/engine"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (s *EngineAPI) remoteEngineAPI(rg *gin.RouterGroup) {

	e := s.cEngine.GetEngine()

	h := e.GetRemoteHandler().(*engine.RemoteHandler)

	rg.POST("/register", s.RX(s.setRemoteOption(e)))
	rg.POST("/bindx/log", s.RX(h.Log))
	rg.POST("/bindx/lazylog", s.RX(h.LazyLog))
	rg.GET("/bindx/self_file", s.RX(h.GetSelfFile))
	rg.GET("/bindx/resource", s.RX(h.ListResources))
	rg.GET("/bindx/resource/:name", s.RX(h.GetResource))

	rg.GET("/bindx/link/in", s.RX(h.InLinks))
	rg.GET("/bindx/link/out", s.RX(h.OutLinks))
	rg.POST("/bindx/link/exec", s.RX(h.LinkExec))

	rg.POST("/bindx/module", s.RX(h.NewModule))
	rg.POST("/bindx/module/ticket", s.RX(h.ModuleTicket))
	rg.POST("/bindx/module/:mid/exec", s.RX(h.ModuleExec))

}

func (s *EngineAPI) setRemoteOption(e etypes.Engine) func(aclaim *claim.RemoteExec, ctx *gin.Context) {

	return func(aclaim *claim.RemoteExec, ctx *gin.Context) {

		opts := etypes.RemoteOptions{}
		err := ctx.BindJSON(&opts)
		if err != nil {
			httpx.WriteErr(ctx, err)
			return
		}

		e.SetRemoteOption(opts)
	}
}

func (s *EngineAPI) RX(fn func(aclaim *claim.RemoteExec, ctx *gin.Context)) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		aclaim, err := s.signer.ParseRemoteExec(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
		if err != nil {
			httpx.UnAuthorized(ctx)
			return
		}

		fn(aclaim, ctx)
	}
}
