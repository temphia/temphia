package server

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/controllers/sockd"
	"github.com/temphia/temphia/code/backend/services/sockdhub/transports"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

var (
	DevPushMaxSize int64 = 100 << 20 // ~ 100 mb
)

func (s *Server) devAPI(rg *gin.RouterGroup) {

	rg.GET("/bprint/file", s.devX(s.DevBprintFileList))
	rg.POST("/bprint/file", s.devX(s.DevBprintFilePush))
	rg.GET("/bprint/file/:file", s.devX(s.DevBprintFileGet))
	rg.DELETE("/bprint/file/:file", s.devX(s.DevBprintFileDel))

	rg.GET("/exec/watch/plug/:pid/agent/:aid", s.devX(s.DevExecWatch))
	rg.POST("/exec/reset/plug/:pid/agent/:aid", s.devX(s.DevExecReset))
	rg.POST("/exec/run/plug/:pid/agent/:aid/:action", s.devX(s.DevExecRun))

	rg.POST("/plug/:pid", s.devX(s.DevModifyPlug))
	rg.POST("/plug/:pid/agent/:aid", s.devX(s.DevModifyAgent))

}

func (s *Server) DevBprintFileList(dclaim *claim.PlugDevTkt, ctx *gin.Context) {
	resp, err := s.cDev.DevBprintFileList(dclaim)
	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) DevBprintFileDel(dclaim *claim.PlugDevTkt, ctx *gin.Context) {
	err := s.cDev.DevBprintFileDel(dclaim, ctx.Param("file"))
	httpx.WriteJSON(ctx, nil, err)
}

func (s *Server) DevBprintFileGet(dclaim *claim.PlugDevTkt, ctx *gin.Context) {
	out, err := s.cDev.DevBprintFileGet(dclaim, ctx.Param("file"))
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteBinary(ctx, out)
}

func (s *Server) DevBprintFilePush(dclaim *claim.PlugDevTkt, ctx *gin.Context) {
	pp.Println("@1")

	mreader, err := ctx.Request.MultipartReader()
	if err != nil {
		pp.Println(err)
		httpx.WriteErr(ctx, err)
		return
	}

	pp.Println("@2")

	form, err := mreader.ReadForm(DevPushMaxSize)
	if err != nil {
		pp.Println("@2", err)
		httpx.WriteErr(ctx, err)
		return
	}

	files := make(map[string]io.Reader, len(form.File))

	pp.Println(form.File)

	for _, fv := range form.File["files"] {
		file, err := fv.Open()
		if err != nil {
			pp.Println(err)
			httpx.WriteErr(ctx, err)
			return
		}

		files[fv.Filename] = file
	}

	pp.Println(files)

	err = s.cDev.DevPushFiles(dclaim, files)
	httpx.WriteFinal(ctx, err)
}

func (s *Server) DevExecWatch(dclaim *claim.PlugDevTkt, ctx *gin.Context) {

	conn, err := transports.NewConnWS(ctx, s.sockdConnIdGenerator.Generate().Int64())
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cSockd.AddDevConn(sockd.DevConnOptions{
		TenantId: dclaim.TenantId,
		UserId:   dclaim.UserId,
		PlugId:   ctx.Param("pid"),
		AgentId:  ctx.Param("aid"),
		Conn:     conn,
	})

	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

}

func (s *Server) DevExecReset(dclaim *claim.PlugDevTkt, ctx *gin.Context) {
	err := s.cEngine.Reset(dclaim.TenantId, ctx.Param("pid"), ctx.Param("aid"))
	httpx.WriteFinal(ctx, err)
}

func (s *Server) DevExecRun(dclaim *claim.PlugDevTkt, ctx *gin.Context) {

	plugId := ctx.Param("pid")
	agentId := ctx.Param("aid")
	action := ctx.Param("action")

	data, err := ctx.GetRawData()
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	out, err := s.cEngine.ExecuteDev(dclaim, plugId, agentId, action, data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	ctx.Writer.Write(out)
}

func (s *Server) DevModifyPlug(dclaim *claim.PlugDevTkt, ctx *gin.Context) {

	data := make(map[string]any)

	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cDev.DevModifyPlug(dclaim, ctx.Param("pid"), data)
	httpx.WriteJSON(ctx, nil, err)

}

func (s *Server) DevModifyAgent(dclaim *claim.PlugDevTkt, ctx *gin.Context) {
	data := make(map[string]any)

	err := ctx.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cDev.DevModifyAgent(dclaim, ctx.Param("pid"), ctx.Param("aid"), data)
	httpx.WriteJSON(ctx, nil, err)

}

// private

func (s *Server) devX(fn func(dclaim *claim.PlugDevTkt, ctx *gin.Context)) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {

		dc, err := s.parseDevTkt(ctx)
		if err != nil {
			httpx.UnAuthorized(ctx)
			pp.Println(err)
			return
		}

		fn(dc, ctx)
	}

}

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
		httpx.WriteErr(ctx, err)
		return
	}
	ctx.Writer.Write(out)
}
