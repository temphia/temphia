package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/controllers/admin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (a *ApiAdmin) LensAPI(rg *gin.RouterGroup) {
	rg.POST("/app", a.X(a.LensQueryApp))
	rg.POST("/engine", a.X(a.LensQueryEngine))
	rg.POST("/site", a.X(a.LensQuerySite))
}

func (a *ApiAdmin) LensQueryApp(ctx httpx.Request) {
	query := admin.LogQuery{}

	err := ctx.Http.BindJSON(&query)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := a.cAdmin.LensQueryApp(ctx.Session, query)
	a.rutil.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) LensQueryEngine(ctx httpx.Request) {
	query := admin.LogQuery{}

	err := ctx.Http.BindJSON(&query)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := a.cAdmin.LensQueryEngine(ctx.Session, query)
	a.rutil.WriteJSON(ctx.Http, resp, err)
}

func (a *ApiAdmin) LensQuerySite(ctx httpx.Request) {
	query := admin.LogQuery{}

	err := ctx.Http.BindJSON(&query)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := a.cAdmin.LensQuerySite(ctx.Session, query)
	a.rutil.WriteJSON(ctx.Http, resp, err)
}
