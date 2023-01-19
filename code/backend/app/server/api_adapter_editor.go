package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) adapterEditorAPI(rg *gin.RouterGroup) {

	// /z/api/:tenant_id/v2/adapter_editor/serve/:did/:file

	rg.GET("/serve/:did/:file", func(ctx *gin.Context) {
		tenantId := ctx.Param("tenant_id")
		file := ctx.Param("file")

		did, err := strconv.ParseInt(ctx.Param("did"), 10, 64)
		if err != nil {
			return
		}

		s.notz.ServeEditorFile(tenantId, file, did, ctx)
	})

}
