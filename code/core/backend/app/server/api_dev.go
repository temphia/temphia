package server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (s *Server) devAPI(rg *gin.RouterGroup) {

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
