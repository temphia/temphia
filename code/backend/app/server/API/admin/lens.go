package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/admin"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (a *ApiAdmin) LensAPI(rg *gin.RouterGroup) {
	rg.POST("/query", a.X(a.LensQuery))
}

func (a *ApiAdmin) LensQuery(ctx httpx.Request) {
	query := admin.LogQuery{}

	err := ctx.Http.BindJSON(&query)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := a.cAdmin.LensQuery(ctx.Session, query)
	a.rutil.WriteJSON(ctx.Http, resp, err)
}
