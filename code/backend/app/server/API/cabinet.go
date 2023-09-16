package api_server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (s *Server) CabinetAPI(rg *gin.RouterGroup) {

	rg.GET("/:source/", s.X(s.listRootFolder))
	rg.GET("/:source/folder/*fpath", s.X(s.listFolder))
	rg.POST("/:source/folder", s.X(s.newFolder))

	rg.GET("/:source/file/*fpath", s.X(s.getFile))
	rg.POST("/:source/file/*fpath", s.X(s.uploadFile))
	rg.DELETE("/:source/file/*fpath", s.X(s.deleteFile))
	rg.GET("/:source/preview/*fpath", s.getFilePreview)
}

type newFolder struct {
	Folder string `json:"folder,omitempty"`
}

func (s *Server) newFolder(ctx httpx.Request) {
	data := &newFolder{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	httpx.WriteFinal(
		ctx.Http,
		s.cCabinet.AddFolder(ctx.Session,
			ctx.MustParam("source"),
			data.Folder),
	)
}

func (s *Server) uploadFile(ctx httpx.Request) {
	bytes, err := httpx.ReadForm(ctx.Http)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	folder, file := fileFolderFromPath(ctx.MustParam("fpath"))

	err = s.cCabinet.AddBlob(ctx.Session,
		ctx.MustParam("source"),
		folder,
		file,
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
		ctx.MustParam("fpath"))

	pp.Println("@files", files)

	httpx.WriteJSON(ctx.Http, files, err)
}

func (s *Server) getFile(ctx httpx.Request) {

	folder, file := fileFolderFromPath(ctx.MustParam("fpath"))

	bytes, err := s.cCabinet.GetBlob(
		ctx.Session,
		ctx.MustParam("source"),
		folder,
		file)
	if err != nil {
		httpx.WriteErr(ctx.Http, err)
		return
	}

	ctx.Http.Writer.WriteHeader(http.StatusOK)
	ctx.Http.Writer.Write(bytes)
}

func (s *Server) deleteFile(ctx httpx.Request) {
	folder, file := fileFolderFromPath(ctx.MustParam("fpath"))

	err := s.cCabinet.DeleteBlob(
		ctx.Session,
		ctx.MustParam("source"),
		folder,
		file)

	httpx.WriteFinal(ctx.Http, err)
}

func (s *Server) getFilePreview(ctx *gin.Context) {
	folder, file := fileFolderFromPath(ctx.Param("fpath"))

	uclaim, err := s.signer.ParseSession(ctx.Param("tenant_id"), ctx.Query("token"))
	if err != nil {
		httpx.UnAuthorized(ctx)
		return
	}

	bytes, err := s.cCabinet.GetBlob(
		uclaim,
		ctx.Param("source"),
		folder,
		file,
	)

	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteBinary(ctx, bytes)
}

// private

func fileFolderFromPath(path string) (string, string) {
	frags := strings.Split(path, "/")
	return strings.Join(frags[:len(frags)-1], "/"), frags[len(frags)-1]
}
