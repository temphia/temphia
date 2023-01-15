package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (a *ApiAdmin) adapterEditorAPI(rg *gin.RouterGroup) {

	rg.POST("/", a.adapterX(a.adapterSelfUpdate))
	rg.POST("/reset", a.adapterX(a.adapterReset))

	rg.GET("/app", a.adapterX(a.adapterListApps))
	rg.POST("/app", a.adapterX(a.adapterNewApp))
	rg.GET("/app/:id", a.adapterX(a.adapterGetApp))
	rg.POST("/app/:id", a.adapterX(a.adapterUpdateApp))
	rg.DELETE("/app/:id", a.adapterX(a.adapterDeleteApp))

	rg.GET("/hook", a.adapterX(a.adapterListHooks))
	rg.POST("/hook", a.adapterX(a.adapterNewHook))
	rg.GET("/hook/:id", a.adapterX(a.adapterGetHook))
	rg.POST("/hook/:id", a.adapterX(a.adapterUpdateHook))
	rg.DELETE("/hook/:id", a.adapterX(a.adapterDeleteHook))

	rg.POST("/action/:name", a.adapterX(a.adapterPreformAction))

}

func (a *ApiAdmin) adapterSelfUpdate(aclaim *claim.AdapterEditor, ctx *gin.Context) {}
func (a *ApiAdmin) adapterReset(aclaim *claim.AdapterEditor, ctx *gin.Context)      {}

func (a *ApiAdmin) adapterListApps(aclaim *claim.AdapterEditor, ctx *gin.Context)  {}
func (a *ApiAdmin) adapterNewApp(aclaim *claim.AdapterEditor, ctx *gin.Context)    {}
func (a *ApiAdmin) adapterGetApp(aclaim *claim.AdapterEditor, ctx *gin.Context)    {}
func (a *ApiAdmin) adapterUpdateApp(aclaim *claim.AdapterEditor, ctx *gin.Context) {}
func (a *ApiAdmin) adapterDeleteApp(aclaim *claim.AdapterEditor, ctx *gin.Context) {}

func (a *ApiAdmin) adapterListHooks(aclaim *claim.AdapterEditor, ctx *gin.Context)  {}
func (a *ApiAdmin) adapterNewHook(aclaim *claim.AdapterEditor, ctx *gin.Context)    {}
func (a *ApiAdmin) adapterGetHook(aclaim *claim.AdapterEditor, ctx *gin.Context)    {}
func (a *ApiAdmin) adapterUpdateHook(aclaim *claim.AdapterEditor, ctx *gin.Context) {}
func (a *ApiAdmin) adapterDeleteHook(aclaim *claim.AdapterEditor, ctx *gin.Context) {}

func (a *ApiAdmin) adapterPreformAction(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	a.notz.PreformEditorAction(aclaim.TenentId, ctx.Param("name"), aclaim.AdapterId, ctx)
}

// private

func (a *ApiAdmin) adapterX(fn func(aclaim *claim.AdapterEditor, ctx *gin.Context)) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		aclaim, err := a.signer.ParseAdapterEditor(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
		if err != nil {
			httpx.UnAuthorized(ctx)
			return
		}

		fn(aclaim, ctx)
	}
}
