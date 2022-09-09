package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/app/server/middleware"
	"github.com/temphia/temphia/code/core/backend/controllers/admin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

type ApiAdmin struct {
	rutil      httpx.Rutil
	cAdmin     *admin.Controller
	middleware *middleware.Middleware
}

func New(cAdmin *admin.Controller, m *middleware.Middleware) ApiAdmin {
	return ApiAdmin{
		rutil:      httpx.Rutil{},
		cAdmin:     cAdmin,
		middleware: m,
	}
}

func (a *ApiAdmin) API(rg *gin.RouterGroup) {
	a.bprintAPI(rg.Group("/bprint"))
	a.dataAPI(rg.Group("/data"))

	a.userAPI(rg.Group("/user"))
	a.userGroupAPI(rg.Group("/ugroup"))

	a.repoAPI(rg.Group("/repo"))
	a.plugAPI(rg.Group("/plug"))
	a.resourceAPI(rg.Group("/resource"))
	a.tenantAPI(rg.Group("/tenant"))
	a.checkAPI(rg.Group("/check"))

}

func (a *ApiAdmin) X(fn func(ctx httpx.Request)) func(*gin.Context) {
	return a.middleware.Authed(fn)
}
