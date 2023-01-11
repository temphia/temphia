package server

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

var (
	DevPushMaxSize int64 = 100 << 20 // ~ 100 mb
)

func (s *Server) devAPI(rg *gin.RouterGroup) {

	rg.GET("/bprint/file", s.DevBprintFileList)
	rg.POST("/bprint/file", s.DevBprintFilePush)
	rg.GET("/bprint/file/:file", s.DevBprintFileGet)
	rg.DELETE("/bprint/file", s.DevBprintFileDel)

	rg.GET("/exec/watch/plug/:pid", s.DevExecWatch)
	rg.POST("/exec/reset/plug/:pid", s.DevExecReset)
	rg.POST("/exec/run/plug/:pid/agent/:aid/:action", s.DevExecRun)

	rg.POST("/modify", s.DevModifyPlug)
	rg.POST("/modify/agent/:aid", s.DevModifyAgent)

}

func (s *Server) DevBprintFileList(ctx *gin.Context) {}
func (s *Server) DevBprintFileDel(ctx *gin.Context)  {}
func (s *Server) DevBprintFileGet(ctx *gin.Context)  {}
func (s *Server) DevBprintFilePush(ctx *gin.Context) {
	dclaim, err := s.parseDevTkt(ctx)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	mreader, err := ctx.Request.MultipartReader()
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	form, err := mreader.ReadForm(DevPushMaxSize)
	if err != nil {
		return
	}

	files := make(map[string]io.Reader, len(form.File))

	for _, fv := range form.File["files"] {
		file, err := fv.Open()
		if err != nil {
			return
		}

		files[fv.Filename] = file
	}

	err = s.cDev.DevPushFiles(dclaim, files)
	httpx.WriteFinal(ctx, err)
}

func (s *Server) DevExecWatch(ctx *gin.Context) {

	/*

		conn, err := transports.NewConnWS(ctx, r.sessman.SessionId())
		if err != nil {
			httpx.WriteErr(ctx.Http, err)
			return
		}

		tkt, err := r.parseDevTkt(ctx)
		if err != nil {
			httpx.WriteErr(ctx.Http, err)
			return
		}

		agents := ctx.QueryArray("agents")
		plugId := ctx.Param("pid")

		err = r.sockdhub.AddDevConn(sockdhub.DevConnOptions{
			TenantId: tkt.TenantId,
			UserId:   tkt.UserId,
			PlugId:   plugId,
			AgentId:  agents[0],
			Conn:     conn,
		})
		if err != nil {
			httpx.WriteErr(ctx.Http, err)
			return
		}
	*/

}

func (s *Server) DevExecReset(ctx *gin.Context) {

}

func (s *Server) DevExecRun(ctx *gin.Context) {
	tkt, err := s.parseDevTkt(ctx)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	plugId := ctx.Param("pid")
	agentId := ctx.Param("aid")
	action := ctx.Param("action")

	data, err := ctx.GetRawData()
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	out, err := s.cEngine.ExecuteDev(tkt, plugId, agentId, action, data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	ctx.Writer.Write(out)
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
		httpx.WriteErr(ctx, err)
		return
	}
	ctx.Writer.Write(out)
}
