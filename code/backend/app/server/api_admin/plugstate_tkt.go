package apiadmin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (a *ApiAdmin) plugStateTKT(rg *gin.RouterGroup) {

	rg.GET("/", a.pstX(a.listPlugState))
	rg.POST("/", a.pstX(a.addPlugState))
	rg.GET("/:key", a.pstX(a.getPlugState))
	rg.POST("/:key", a.pstX(a.updatePlugState))
	rg.DELETE("/:key", a.pstX(a.deletePlugState))

}

func (a *ApiAdmin) addPlugState(aclaim *claim.PlugState, ctx *gin.Context) {
	data := entities.PlugKV{}

	err := ctx.BindJSON(&data)
	if err != nil {
		a.rutil.WriteErr(ctx, err.Error())
		return
	}

	err = a.cAdmin.AddPlugState(aclaim, data)
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) getPlugState(aclaim *claim.PlugState, ctx *gin.Context) {
	resp, err := a.cAdmin.GetPlugState(aclaim, ctx.Param("key"))
	a.rutil.WriteJSON(ctx, resp, err)
}

func (a *ApiAdmin) updatePlugState(aclaim *claim.PlugState, ctx *gin.Context) {
	var value string

	err := ctx.BindJSON(&value)
	if err != nil {
		a.rutil.WriteErr(ctx, err.Error())
		return
	}

	err = a.cAdmin.UpdatePlugState(aclaim, ctx.Param("key"), value)
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) deletePlugState(aclaim *claim.PlugState, ctx *gin.Context) {
	err := a.cAdmin.DeletePlugState(aclaim, ctx.Param("key"))
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) listPlugState(aclaim *claim.PlugState, ctx *gin.Context) {

	page, _ := strconv.ParseInt(ctx.Param("page"), 10, 64)
	pcount, _ := strconv.ParseInt(ctx.Param("page_count"), 10, 64)

	resp, err := a.cAdmin.ListPlugState(aclaim, int(page), int(pcount), ctx.Query("key_cursor"))
	a.rutil.WriteJSON(ctx, resp, err)
}

// private

func (a *ApiAdmin) pstX(fn func(aclaim *claim.PlugState, ctx *gin.Context)) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		aclaim, err := a.signer.ParsePlugState(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
		if err != nil {
			httpx.UnAuthorized(ctx)
			return
		}

		fn(aclaim, ctx)
	}
}
