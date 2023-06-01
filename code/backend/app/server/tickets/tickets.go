package tickets

import (
	"github.com/temphia/temphia/code/backend/app/server/middleware"
	"github.com/temphia/temphia/code/backend/controllers"
	"github.com/temphia/temphia/code/backend/controllers/admin"
	"github.com/temphia/temphia/code/backend/controllers/cabinet"
)

type TicketAPI struct {
	middleware *middleware.Middleware
	cAdmin     *admin.Controller
	cCabinet   *cabinet.Controller
}

func New(middleware *middleware.Middleware, root *controllers.RootController) *TicketAPI {
	return &TicketAPI{
		middleware: middleware,
		cAdmin:     root.AdminController(),
		cCabinet:   root.CabinetController(),
	}
}
