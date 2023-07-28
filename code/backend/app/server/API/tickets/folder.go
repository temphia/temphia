package tickets

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (s *TicketAPI) Folder(rg *gin.RouterGroup) {

	rg.GET("/:ticket/", s.fx(s.folderTktList))
	rg.GET("/:ticket/:name", s.fx(s.folderTktFile))
	rg.GET("/:ticket/:name/preview", s.fx(s.folderTktPreview))
	rg.POST("/:ticket/:name", s.fx(s.folderTktUpload))
	rg.DELETE("/:ticket/:name", s.fx(s.folderTktDelete))

}

func (s *TicketAPI) folderTktList(uclaim *claim.Folder, ctx *gin.Context) {

	resp, err := s.cTicket.TicketList(uclaim)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteJSON(ctx, resp, err)
}

func (s *TicketAPI) folderTktFile(uclaim *claim.Folder, ctx *gin.Context) {

	file := ctx.Param("name")

	out, err := s.cTicket.TicketFile(uclaim, file)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteBinary(ctx, out)
}

func (s *TicketAPI) folderTktPreview(uclaim *claim.Folder, ctx *gin.Context) {
	file := ctx.Param("name")

	out, err := s.cTicket.TicketFile(uclaim, file)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	httpx.WriteBinary(ctx, out)
}

func (s *TicketAPI) folderTktUpload(uclaim *claim.Folder, ctx *gin.Context) {

	file := ctx.Param("name")

	out, err := httpx.ReadForm(ctx)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	err = s.cTicket.TicketUpload(uclaim, file, out)
	httpx.WriteFinal(ctx, err)
}

func (s *TicketAPI) folderTktDelete(uclaim *claim.Folder, ctx *gin.Context) {}

// utils

func (s *TicketAPI) fx(fn func(uclaim *claim.Folder, ctx *gin.Context)) func(ctx *gin.Context) {
	return s.middleware.FolderX(fn)
}
