package controllers

import (
	"github.com/temphia/temphia/code/core/backend/controllers/admin"
	"github.com/temphia/temphia/code/core/backend/controllers/authed"
	"github.com/temphia/temphia/code/core/backend/controllers/basic"
	"github.com/temphia/temphia/code/core/backend/controllers/cabinet"
	"github.com/temphia/temphia/code/core/backend/controllers/dtable"
	"github.com/temphia/temphia/code/core/backend/controllers/operator"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

// Controller is Parent controller which holds all and inits all
// actual controller, where as controller mostly sits on top of
// api/routes, it gets all claim and params parsed from routes. It
// checks permisission bashed on claim and calls lower level services
// to finish the operations [ api/routes => controller => services => db/providers]

type RootController struct {
	cAdmin    *admin.Controller
	cAuth     *authed.Controller
	cBasic    *basic.Controller
	cCabinet  *cabinet.Controller
	cDtable   *dtable.Controller
	cOperator *operator.Controller
}

func New(_app xtypes.App) *RootController {

	deps := _app.GetDeps()

	cplane := deps.ControlPlane().(xplane.ControlPlane)
	seq := cplane.GetSequencer()

	corehub := deps.CoreHub().(store.CoreHub)
	pacman := deps.Pacman().(service.Pacman)
	signer := deps.Signer().(service.Signer)
	cab := deps.Cabinet().(store.CabinetHub)
	dynhub := deps.DynHub().(store.DynHub)

	return &RootController{
		cAdmin:    admin.New(pacman, cplane, corehub, signer),
		cAuth:     authed.New(corehub, signer, seq),
		cBasic:    basic.New(corehub, cab, dynhub, pacman),
		cCabinet:  cabinet.New(cab, signer),
		cDtable:   dtable.New(dynhub, cab, signer),
		cOperator: operator.New(corehub, signer, _app, "", ""), // fixme =>
	}
}

func (c *RootController) AdminController() *admin.Controller       { return c.cAdmin }
func (c *RootController) AuthController() *authed.Controller       { return c.cAuth }
func (c *RootController) BasicController() *basic.Controller       { return c.cBasic }
func (c *RootController) CabinetController() *cabinet.Controller   { return c.cCabinet }
func (c *RootController) DtableController() *dtable.Controller     { return c.cDtable }
func (c *RootController) OperatorController() *operator.Controller { return c.cOperator }
