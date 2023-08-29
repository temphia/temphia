package apiadmin

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
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

func (a *ApiAdmin) adapterSelfUpdate(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	data := make(map[string]any)
	err := ctx.BindJSON(&data)
	if err != nil {
		a.rutil.WriteErr(ctx, err.Error())
		return
	}

	err = a.cAdmin.AdapterSelfUpdate(aclaim, data)
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) adapterReset(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	// a.notz.Reset(aclaim.TenantId, aclaim.AdapterId)
	time.Sleep(time.Second * 5)
}

func (a *ApiAdmin) adapterListApps(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	resp, err := a.cAdmin.AdapterListApps(aclaim)
	a.rutil.WriteJSON(ctx, resp, err)
}

func (a *ApiAdmin) adapterNewApp(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	data := &entities.TargetApp{}
	err := ctx.BindJSON(data)
	if err != nil {
		a.rutil.WriteErr(ctx, err.Error())
		return
	}

	err = a.cAdmin.AdapterNewApp(aclaim, data)
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) adapterGetApp(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	resp, err := a.cAdmin.AdapterGetApp(aclaim, id)
	a.rutil.WriteJSON(ctx, resp, err)
}

func (a *ApiAdmin) adapterUpdateApp(aclaim *claim.AdapterEditor, ctx *gin.Context) {

	data := make(map[string]any)
	err := ctx.BindJSON(&data)
	if err != nil {
		a.rutil.WriteErr(ctx, err.Error())
		return
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err = a.cAdmin.AdapterUpdateApp(aclaim, id, data)
	a.rutil.WriteFinal(ctx, err)

}

func (a *ApiAdmin) adapterDeleteApp(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := a.cAdmin.AdapterDeleteApp(aclaim, id)
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) adapterListHooks(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	data := &entities.TargetHook{}
	err := ctx.BindJSON(data)
	if err != nil {
		a.rutil.WriteErr(ctx, err.Error())
		return
	}

	err = a.cAdmin.AdapterNewHook(aclaim, data)
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) adapterNewHook(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	data := &entities.TargetHook{}
	err := ctx.BindJSON(data)
	if err != nil {
		a.rutil.WriteErr(ctx, err.Error())
		return
	}

	err = a.cAdmin.AdapterNewHook(aclaim, data)
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) adapterGetHook(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	resp, err := a.cAdmin.AdapterGetHook(aclaim, id)
	a.rutil.WriteJSON(ctx, resp, err)
}

func (a *ApiAdmin) adapterUpdateHook(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	data := make(map[string]any)
	err := ctx.BindJSON(&data)
	if err != nil {
		a.rutil.WriteErr(ctx, err.Error())
		return
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err = a.cAdmin.AdapterUpdateHook(aclaim, id, data)
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) adapterDeleteHook(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := a.cAdmin.AdapterDeleteHook(aclaim, id)
	a.rutil.WriteFinal(ctx, err)
}

func (a *ApiAdmin) adapterPreformAction(aclaim *claim.AdapterEditor, ctx *gin.Context) {
	/*	out, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			httpx.WriteErr(ctx, err)
			return
		}

		uctx := &claim.UserContext{
			UserID:    aclaim.UserID,
			UserGroup: aclaim.UserGroup,
			SessionID: aclaim.SessionID,
			DeviceId:  aclaim.DeviceId,
		}

		resp, err := a.notz.PreformEditorAction(httpx.AdapterEditorContext{
			Id:   aclaim.AdapterId,
			User: uctx,
			Name: ctx.Param("name"),
			Data: out,
		})

		httpx.WriteJSON(ctx, resp, err)
	*/

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
