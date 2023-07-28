package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/app/server/API/tickets"
	"github.com/temphia/temphia/code/backend/app/server/middleware"
	"github.com/temphia/temphia/code/backend/controllers/admin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/service"
)

type Options struct {
	Admin      *admin.Controller
	MiddleWare *middleware.Middleware
	//	Notz       *notz.Notz
	Signer    service.Signer
	TicketAPI *tickets.TicketAPI
}

type ApiAdmin struct {
	rutil      httpx.Rutil
	cAdmin     *admin.Controller
	middleware *middleware.Middleware
	//	notz       *notz.Notz
	signer    service.Signer
	ticketAPI *tickets.TicketAPI
}

func New(opts Options) ApiAdmin {
	return ApiAdmin{
		rutil:      httpx.Rutil{},
		cAdmin:     opts.Admin,
		middleware: opts.MiddleWare,
		//		notz:       opts.Notz,
		signer:    opts.Signer,
		ticketAPI: opts.TicketAPI,
	}
}

func (a *ApiAdmin) API(rg *gin.RouterGroup) {
	a.bprintAPI(rg.Group("/bprint"))
	a.dataAPI(rg.Group("/data"))

	a.userAPI(rg.Group("/user"))
	a.userGroupAPI(rg.Group("/ugroup"))

	a.ticketAPI.User(rg.Group("/user_tkt"))

	a.repoAPI(rg.Group("/repo"))
	a.plugAPI(rg.Group("/plug"))
	a.resourceAPI(rg.Group("/resource"))
	a.tenantAPI(rg.Group("/tenant"))
	a.TargetAPI(rg.Group("/target"))
	a.LensAPI(rg.Group("/lens"))
	a.adapterEditorAPI(rg.Group("/adapter_editor"))

	a.ticketAPI.PlugState(rg.Group("/plug_state"))

}

func (a *ApiAdmin) X(fn func(ctx httpx.Request)) func(*gin.Context) {
	return a.middleware.LoggedX(fn)
}
