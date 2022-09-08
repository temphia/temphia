package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (s *Server) cabinetAPI(rg *gin.RouterGroup) {

	rg.GET("/", s.X(s.ListRootFolder))
	rg.GET("/:folder", s.X(s.ListFolder))
	rg.POST("/:folder", s.X(s.NewFolder))
	rg.GET("/:folder/file/:fname", s.X(s.GetFile))
	rg.POST("/:folder/file/:fname", s.X(s.UploadFile))
	rg.DELETE("/:folder/file/:fname", s.X(s.DeleteFile))
	rg.GET("/:folder/preview/:fname", s.X(s.GetFilePreview))

}

func (s *Server) ListCabinetSources(ctx httpx.Request) {

	sources, err := s.cBasic.ListCabinetSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, sources, err)
}

func (s *Server) NewFolder(ctx httpx.Request) {
	httpx.WriteFinal(
		ctx.Http,
		s.cCabinet.AddFolder(ctx.Session, ctx.Http.Param("folder")),
	)
}

func (s *Server) UploadFile(ctx httpx.Request) {
	bytes, err := httpx.ReadForm(ctx.Http)
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	err = s.cCabinet.AddBlob(ctx.Session, ctx.Http.Param("folder"), ctx.Http.Param("fname"), bytes)
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) ListRootFolder(ctx httpx.Request) {
	folders, err := s.cCabinet.ListRoot(ctx.Session)
	httpx.WriteJSON(ctx.Http, folders, err)
}

func (s *Server) ListFolder(ctx httpx.Request) {
	files, err := s.cCabinet.ListFolder(ctx.Session, ctx.Http.Param("folder"))
	httpx.WriteJSON(ctx.Http, files, err)
}

func (s *Server) GetFile(ctx httpx.Request) {

	bytes, err := s.cCabinet.GetBlob(ctx.Session, ctx.Http.Param("folder"), ctx.Http.Param("fname"))
	if err != nil {
		httpx.WriteErr(ctx.Http, err.Error())
		return
	}

	ctx.Http.Writer.WriteHeader(http.StatusOK)
	ctx.Http.Writer.Write(bytes)
}

func (s *Server) DeleteFile(ctx httpx.Request) {
	err := s.cCabinet.DeleteBlob(ctx.Session, ctx.Http.Param("folder"), ctx.Http.Param("fname"))
	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) GetFolderTicket(ctx httpx.Request) {

	resp, err := s.cCabinet.NewFolderTicket(ctx.Session, ctx.Http.Param("folder"))

	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) GetFilePreview(ctx httpx.Request) {

	// bytes, err := b.blobfs.GetBlobPreview(c.Request.Context(), cm.TenentId, c.Param("folder"), c.Param("fname"))
	// if err != nil {
	// 	apiutils.WriteErr(c, err.Error())
	// 	return
	// }
	// c.Writer.Header().Add("Content-Type", "application/octet-stream")
	// c.Writer.Write(bytes)
}
