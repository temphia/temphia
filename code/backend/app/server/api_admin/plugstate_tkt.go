package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/admin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func (a *ApiAdmin) plugStateTKT(rg *gin.RouterGroup) {

	rg.GET("/query", a.pstX(a.listPlugState))
	rg.POST("/key", a.pstX(a.addPlugState))
	rg.GET("/key/:key", a.pstX(a.getPlugState))
	rg.POST("/key/:key", a.pstX(a.updatePlugState))
	rg.DELETE("/key/:key", a.pstX(a.deletePlugState))

}

func (a *ApiAdmin) addPlugState(aclaim *claim.PlugState, ctx *gin.Context) {
	data := admin.AddPlugStateOptions{}
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

	data := admin.UpdatePlugStateOptions{}
	err := ctx.BindJSON(&data)
	if err != nil {
		a.rutil.WriteErr(ctx, err.Error())
		return
	}

	err = a.cAdmin.UpdatePlugState(aclaim, data)
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) deletePlugState(aclaim *claim.PlugState, ctx *gin.Context) {
	err := a.cAdmin.DeletePlugState(aclaim, ctx.Param("key"))
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) listPlugState(aclaim *claim.PlugState, ctx *gin.Context) {

	query := store.PkvQuery{}
	err := ctx.BindJSON(&query)
	if err != nil {
		a.rutil.WriteErr(ctx, err.Error())
		return
	}

	resp, err := a.cAdmin.ListPlugState(aclaim, &query)
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
