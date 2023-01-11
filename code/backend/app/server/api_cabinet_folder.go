package server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (s *Server) folderTktAPI(rg *gin.RouterGroup) {

	rg.GET("/:ticket/", s.fx(s.folderTktList))
	rg.GET("/:ticket/:name", s.fx(s.folderTktFile))
	rg.GET("/:ticket/:name/preview", s.fx(s.folderTktPreview))
	rg.POST("/:ticket/:name", s.fx(s.folderTktUpload))
	rg.DELETE("/:ticket/:name", s.fx(s.folderTktDelete))

}

func (s *Server) folderTktList(uclaim *claim.Folder, ctx *gin.Context) {

	resp, err := s.cCabinet.TicketList(uclaim)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteJSON(ctx, resp, err)
}

func (s *Server) folderTktFile(uclaim *claim.Folder, ctx *gin.Context) {

	file := ctx.Param("name")

	out, err := s.cCabinet.TicketFile(uclaim, file)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteBinary(ctx, out)
}

func (s *Server) folderTktPreview(uclaim *claim.Folder, ctx *gin.Context) {
	file := ctx.Param("name")

	out, err := s.cCabinet.TicketFile(uclaim, file)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteBinary(ctx, out)
}

func (s *Server) folderTktUpload(uclaim *claim.Folder, ctx *gin.Context) {

	file := ctx.Param("name")

	out, err := httpx.ReadForm(ctx)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cCabinet.TicketUpload(uclaim, file, out)
	httpx.WriteFinal(ctx, err)
}

func (s *Server) folderTktDelete(uclaim *claim.Folder, ctx *gin.Context) {}

// utils

func (s *Server) fx(fn func(uclaim *claim.Folder, ctx *gin.Context)) func(ctx *gin.Context) {
	return s.middleware.FolderX(fn)
}
