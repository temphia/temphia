package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) folderTktAPI(rg *gin.RouterGroup) {

	// rg.GET("/:ticket/", s.folderTktList)
	// rg.GET("/:ticket/:name", s.folderTktFile)
	// rg.GET("/:ticket/:name/preview", s.folderTktPreview)
	// rg.POST("/:ticket/:name", s.folderTktUpload)
	// rg.DELETE("/:ticket/:name", s.folderTktDelete)

}

// func (s *Server) folderTktList(ctx *gin.Context) {
// 	ticket := ctx.Param("ticket")
// 	tenantId := ctx.Param("tenant_id")

// 	ct, err := s.signer.ParseFolderTkt(tenantId, ticket)
// 	if err != nil {
// 		httpx.WriteErr(ctx, err)
// 		return
// 	}

// 	resp, err := s.cCabinet.TicketList(tenantId, ct)
// 	if err != nil {
// 		httpx.WriteErr(ctx, err)
// 		return
// 	}

// 	httpx.WriteJSON(ctx, resp, err)
// }

// func (s *Server) folderTktFile(ctx *gin.Context) {
// 	ticket := ctx.Param("ticket")
// 	tenantId := ctx.Param("tenant_id")
// 	file := ctx.Param("name")

// 	ct, err := s.signer.ParseFolderTkt(tenantId, ticket)
// 	if err != nil {
// 		httpx.WriteErr(ctx, err)
// 		return
// 	}

// 	out, err := s.cCabinet.TicketFile(tenantId, file, ct)
// 	if err != nil {
// 		httpx.WriteErr(ctx, err)
// 		return
// 	}

// 	httpx.WriteBinary(ctx, out)
// }

// func (s *Server) folderTktPreview(ctx *gin.Context) {
// 	ticket := ctx.Param("ticket")
// 	tenantId := ctx.Param("tenant_id")
// 	file := ctx.Param("name")

// 	ct, err := s.signer.ParseFolderTkt(tenantId, ticket)
// 	if err != nil {
// 		httpx.WriteErr(ctx, err)
// 		return
// 	}

// 	out, err := s.cCabinet.TicketFile(tenantId, file, ct)
// 	if err != nil {
// 		httpx.WriteErr(ctx, err)
// 		return
// 	}

// 	httpx.WriteBinary(ctx, out)
// }

// func (s *Server) folderTktUpload(ctx *gin.Context) {
// 	ticket := ctx.Param("ticket")
// 	tenantId := ctx.Param("tenant_id")
// 	file := ctx.Param("name")

// 	ct, err := s.signer.ParseFolderTkt(tenantId, ticket)
// 	if err != nil {
// 		httpx.WriteErr(ctx, err)
// 		return
// 	}

// 	out, err := httpx.ReadForm(ctx)
// 	if err != nil {
// 		httpx.WriteErr(ctx, err)
// 		return
// 	}

// 	err = s.cCabinet.TicketUpload(tenantId, file, out, ct)
// 	httpx.WriteFinal(ctx, err)
// }

// func (s *Server) folderTktDelete(ctx *gin.Context) {

// }
