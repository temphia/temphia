package server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

var (
	DevPushMaxSize int64 = 100 << 20 // ~ 100 mb
)

func (s *Server) devAPI(rg *gin.RouterGroup) {

	// dev.GET("/bprint/file", s.routes.DevBprintFileList)
	// dev.POST("/bprint/file", s.routes.DevBprintFilePush)
	// dev.GET("/bprint/file/:file", s.routes.DevBprintFileGet)
	// dev.DELETE("/bprint/file", s.routes.DevBprintFileDel)

	// dev.GET("/exec/watch/plug/:pid", s.routes.DevExecWatch)
	// dev.POST("/exec/reset/plug/:pid", s.routes.DevExecReset)
	// dev.POST("/exec/run/plug/:pid/agent/:aid/:action", s.routes.DevExecRun)

	// dev.POST("/modify", s.routes.DevModifyPlug)
	// dev.POST("/modify/agent/:aid", s.routes.DevModifyAgent)

}

func (s *Server) DevBprintFileList(ctx *gin.Context) {}
func (s *Server) DevBprintFileDel(ctx *gin.Context)  {}
func (s *Server) DevBprintFileGet(ctx *gin.Context)  {}
func (s *Server) DevBprintFilePush(ctx *gin.Context) {
	// dclaim, err := r.parseDevTkt(ctx)
	// if err != nil {
	// 	r.rutil.WriteErr(ctx, err.Error())
	// 	return
	// }

	// mreader, err := ctx.Request.MultipartReader()
	// if err != nil {
	// 	r.rutil.WriteErr(ctx, err.Error())
	// 	return
	// }

	// form, err := mreader.ReadForm(DevPushMaxSize)
	// if err != nil {
	// 	return
	// }

	// files := make(map[string]io.Reader, len(form.File))

	// for _, fv := range form.File["files"] {
	// 	file, err := fv.Open()
	// 	if err != nil {
	// 		return
	// 	}

	// 	files[fv.Filename] = file
	// }

	// err = r.cAdmin.DevPushFiles(dclaim, files)
	// r.rutil.WriteFinal(ctx, err)
}

func (s *Server) DevExecWatch(ctx *gin.Context) {

	// conn, err := transports.NewConnWS(ctx, r.sessman.SessionId())
	// if err != nil {
	// 	r.rutil.WriteErr(ctx, err.Error())
	// 	return
	// }

	// tkt, err := r.parseDevTkt(ctx)
	// if err != nil {
	// 	r.rutil.WriteErr(ctx, err.Error())
	// 	return
	// }

	// agents := ctx.QueryArray("agents")
	// plugId := ctx.Param("pid")

	// err = r.sockdhub.AddDevConn(sockdhub.DevConnOptions{
	// 	TenantId: tkt.TenantId,
	// 	UserId:   tkt.UserId,
	// 	PlugId:   plugId,
	// 	AgentId:  agents[0],
	// 	Conn:     conn,
	// })
	// if err != nil {
	// 	r.rutil.WriteErr(ctx, err.Error())
	// 	return
	// }

}

func (s *Server) DevExecReset(ctx *gin.Context) {

}

func (s *Server) DevExecRun(ctx *gin.Context) {
	// tkt, err := r.parseDevTkt(ctx)
	// if err != nil {
	// 	r.rutil.WriteErr(ctx, err.Error())
	// 	return
	// }

	// plugId := ctx.Param("pid")
	// agentId := ctx.Param("aid")
	// action := ctx.Param("action")

	// // fixme check if tkt can execute this plug/agent

	// r.engine.ExecAction(tkt.TenantId, plugId, agentId, action, ctx)
}

func (s *Server) DevModifyPlug(ctx *gin.Context) {

}

func (s *Server) DevModifyAgent(ctx *gin.Context) {

}

// private

func (s *Server) parseDevTkt(ctx *gin.Context) (*claim.PlugDevTkt, error) {

	authtok := ctx.GetHeader("Authorization")
	if authtok == "" {
		authtok = ctx.Query("token")
	}

	return s.signer.ParsePlugDevTkt(ctx.Param("tenant_id"), authtok)
}

func (s *Server) serveInterface(ctx *gin.Context) {
	// @core/<interface_file>.json

	out, err := s.app.Data().GetIfaceFile(ctx.Param("name"))
	if err != nil {
		httpx.WriteErr(ctx, err.Error())
		return
	}
	ctx.Writer.Write(out)
}
