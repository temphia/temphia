package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/controllers/admin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx"
)

func (a *ApiAdmin) LensAPI(rg *gin.RouterGroup) {
	rg.POST("/:index", a.X(a.LensQuery))
}

func (a *ApiAdmin) LensQuery(ctx httpx.Request) {
	query := admin.LogQuery{}

	err := ctx.Http.BindJSON(&query)
	if err != nil {
		a.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	var resp []logx.Log

	switch ctx.MustParam("index") {
	case "app":
		resp, err = a.cAdmin.LensQueryApp(ctx.Session, query)
	default:
		panic("Not implemented")
	}

	a.rutil.WriteJSON(ctx.Http, resp, err)
}
