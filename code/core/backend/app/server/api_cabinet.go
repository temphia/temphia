package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (s *Server) cabinetAPI(rg *gin.RouterGroup) {

	rg.GET("/", s.X(s.listRootFolder))
	rg.GET("/:folder", s.X(s.listFolder))
	rg.POST("/:folder", s.X(s.newFolder))
	rg.GET("/:folder/file/:fname", s.X(s.getFile))
	rg.POST("/:folder/file/:fname", s.X(s.uploadFile))
	rg.DELETE("/:folder/file/:fname", s.X(s.deleteFile))

	rg.GET("/:folder/preview/:fname", s.X(s.getFilePreview))
}

func (s *Server) newFolder(ctx httpx.Request) {
	httpx.WriteFinal(
		ctx.Http,
		s.cCabinet.AddFolder(nil, ctx.Http.Param("folder")),
	)
}

func (s *Server) uploadFile(ctx httpx.Request) {
	bytes, err := httpx.ReadForm(ctx.Http)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	err = s.cCabinet.AddBlob(ctx.Session, ctx.Http.Param("folder"), ctx.Http.Param("fname"), bytes)
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) listRootFolder(ctx httpx.Request) {
	folders, err := s.cCabinet.ListRoot(nil)
	httpx.WriteJSON(ctx.Http, folders, err)
}

func (s *Server) listFolder(ctx httpx.Request) {
	files, err := s.cCabinet.ListFolder(ctx.Session, ctx.Http.Param("folder"))
	httpx.WriteJSON(ctx.Http, files, err)
}

func (s *Server) getFile(ctx httpx.Request) {

	bytes, err := s.cCabinet.GetBlob(ctx.Session, ctx.Http.Param("folder"), ctx.Http.Param("fname"))
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	ctx.Http.Writer.WriteHeader(http.StatusOK)
	ctx.Http.Writer.Write(bytes)
}

func (s *Server) deleteFile(ctx httpx.Request) {
	err := s.cCabinet.DeleteBlob(ctx.Session, ctx.Http.Param("folder"), ctx.Http.Param("fname"))
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) getFilePreview(ctx httpx.Request) {

	// bytes, err := b.blobfs.GetBlobPreview(c.Request.Context(), cm.TenentId, c.Param("folder"), c.Param("fname"))
	// if err != nil {
	// 	apiutils.WriteErr(c, err.Error())
	// 	return
	// }
	// c.Writer.Header().Add("Content-Type", "application/octet-stream")
	// c.Writer.Write(bytes)
}
