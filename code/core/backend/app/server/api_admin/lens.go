package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/controllers/admin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (a *ApiAdmin) LensAPI(rg *gin.RouterGroup) {
	rg.GET("/app", a.X(a.LensQueryApp))
	rg.GET("/engine", a.X(a.LensQueryEngine))
	rg.GET("/site", a.X(a.LensQuerySite))
}

func (a *ApiAdmin) LensQueryApp(ctx httpx.Request) {
	query := admin.LogQuery{}

	err := ctx.Http.BindJSON(&query)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	a.cAdmin.LensQueryApp(ctx.Session, query)
}

func (a *ApiAdmin) LensQueryEngine(ctx httpx.Request) {
	query := admin.LogQuery{}

	err := ctx.Http.BindJSON(&query)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	a.cAdmin.LensQueryEngine(ctx.Session, query)

}

func (a *ApiAdmin) LensQuerySite(ctx httpx.Request) {
	query := admin.LogQuery{}

	err := ctx.Http.BindJSON(&query)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	a.cAdmin.LensQuerySite(ctx.Session, query)
}
