package api_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

func (s *Server) CabinetAPI(rg *gin.RouterGroup) {

	rg.GET("/:source/", s.X(s.listRootFolder))
	rg.GET("/:source/:folder", s.X(s.listFolder))
	rg.POST("/:source/:folder", s.X(s.newFolder))
	rg.GET("/:source/:folder/file/:fname", s.X(s.getFile))
	rg.POST("/:source/:folder/file/:fname", s.X(s.uploadFile))
	rg.DELETE("/:source/:folder/file/:fname", s.X(s.deleteFile))
	rg.GET("/:source/:folder/preview/:fname", s.getFilePreview)
}

func (s *Server) newFolder(ctx httpx.Request) {

	httpx.WriteFinal(
		ctx.Http,
		s.cCabinet.AddFolder(ctx.Session,
			ctx.MustParam("source"),
			ctx.MustParam("folder")),
	)
}

func (s *Server) uploadFile(ctx httpx.Request) {
	bytes, err := httpx.ReadForm(ctx.Http)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	err = s.cCabinet.AddBlob(ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("folder"),
		ctx.MustParam("fname"),
		bytes)
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) listRootFolder(ctx httpx.Request) {
	folders, err := s.cCabinet.ListRoot(ctx.Session,
		ctx.MustParam("source"))
	httpx.WriteJSON(ctx.Http, folders, err)
}

func (s *Server) listFolder(ctx httpx.Request) {
	files, err := s.cCabinet.ListFolder(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("folder"))
	httpx.WriteJSON(ctx.Http, files, err)
}

func (s *Server) getFile(ctx httpx.Request) {

	bytes, err := s.cCabinet.GetBlob(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("folder"),
		ctx.MustParam("fname"))
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	ctx.Http.Writer.WriteHeader(http.StatusOK)
	ctx.Http.Writer.Write(bytes)
}

func (s *Server) deleteFile(ctx httpx.Request) {
	err := s.cCabinet.DeleteBlob(
		ctx.Session,
		ctx.MustParam("source"),
		ctx.MustParam("folder"),
		ctx.MustParam("fname"))

	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) getFilePreview(ctx *gin.Context) {

	uclaim, err := s.signer.ParseSession(ctx.Param("tenant_id"), ctx.Query("token"))
	if err != nil {
		httpx.UnAuthorized(ctx)
		return
	}

	bytes, err := s.cCabinet.GetBlob(
		uclaim,
		ctx.Param("source"),
		ctx.Param("folder"),
		ctx.Param("fname"))
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteBinary(ctx, bytes)
}
