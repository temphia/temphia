package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/app/server/middleware"
	"github.com/temphia/temphia/code/core/backend/controllers/admin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

type Options struct {
	Admin      *admin.Controller
	MiddleWare *middleware.Middleware
}

type ApiAdmin struct {
	rutil      httpx.Rutil
	cAdmin     *admin.Controller
	middleware *middleware.Middleware
}

func New(opts Options) ApiAdmin {
	return ApiAdmin{
		rutil:      httpx.Rutil{},
		cAdmin:     opts.Admin,
		middleware: opts.MiddleWare,
	}
}

func (a *ApiAdmin) API(rg *gin.RouterGroup) {
	a.bprintAPI(rg.Group("/bprint"))
	a.dataAPI(rg.Group("/data"))

	a.userAPI(rg.Group("/user"))
	a.userGroupAPI(rg.Group("/ugroup"))
	a.userTktAPI(rg.Group("/user_tkt"))

	a.repoAPI(rg.Group("/repo"))
	a.plugAPI(rg.Group("/plug"))
	a.resourceAPI(rg.Group("/resource"))
	a.tenantAPI(rg.Group("/tenant"))
	a.checkAPI(rg.Group("/check"))
	a.TargetAPI(rg.Group("/target"))

}

func (a *ApiAdmin) X(fn func(ctx httpx.Request)) func(*gin.Context) {
	return a.middleware.Authed(fn)
}
