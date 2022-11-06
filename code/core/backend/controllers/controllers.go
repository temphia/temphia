package controllers

import (
	"github.com/temphia/temphia/code/core/backend/controllers/admin"
	"github.com/temphia/temphia/code/core/backend/controllers/authed"
	"github.com/temphia/temphia/code/core/backend/controllers/basic"
	"github.com/temphia/temphia/code/core/backend/controllers/cabinet"
	"github.com/temphia/temphia/code/core/backend/controllers/data"
	"github.com/temphia/temphia/code/core/backend/controllers/dev"
	"github.com/temphia/temphia/code/core/backend/controllers/engine"
	"github.com/temphia/temphia/code/core/backend/controllers/operator"
	"github.com/temphia/temphia/code/core/backend/controllers/repo"
	"github.com/temphia/temphia/code/core/backend/controllers/sockd"
	"github.com/temphia/temphia/code/core/backend/controllers/user"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
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
	cDtable   *data.Controller
	cDev      *dev.Controller
	cEngine   *engine.Controller
	cRepo     *repo.Controller
	cSockd    *sockd.Controller
	cOperator *operator.Controller
	cUser     *user.Controller
}

type Options struct {
	App              xtypes.App
	OperatorUser     string
	OperatorPassword string
}

func New(opts Options) *RootController {

	deps := opts.App.GetDeps()

	cplane := deps.ControlPlane().(xplane.ControlPlane)
	seq := cplane.GetSequencer()

	corehub := deps.CoreHub().(store.CoreHub)
	pacman := deps.RepoHub().(repox.Hub)
	signer := deps.Signer().(service.Signer)
	cab := deps.Cabinet().(store.CabinetHub)
	dynhub := deps.DataHub().(store.DataHub)
	egine := deps.Engine().(etypes.Engine)

	return &RootController{
		cAdmin:   admin.New(pacman, cplane, corehub, signer, dynhub, cab),
		cAuth:    authed.New(corehub, signer, seq),
		cBasic:   basic.New(corehub, cab, dynhub, pacman),
		cCabinet: cabinet.New(cab, signer),
		cDtable:  data.New(dynhub, cab, signer),
		cRepo:    repo.New(pacman),
		cOperator: operator.New(
			corehub,
			signer,
			opts.App,
			opts.OperatorUser,
			opts.OperatorPassword),
		cDev:    dev.New(pacman, corehub),
		cEngine: engine.New(egine, signer),
		cUser:   user.New(corehub),
		//		cSockd:  sockd.New(deps.SockdHub().(sockdx.Hub).GetSockd()),
	}
}

func (c *RootController) AdminController() *admin.Controller       { return c.cAdmin }
func (c *RootController) AuthController() *authed.Controller       { return c.cAuth }
func (c *RootController) BasicController() *basic.Controller       { return c.cBasic }
func (c *RootController) CabinetController() *cabinet.Controller   { return c.cCabinet }
func (c *RootController) DtableController() *data.Controller       { return c.cDtable }
func (c *RootController) OperatorController() *operator.Controller { return c.cOperator }

func (c *RootController) SockdController() *sockd.Controller   { return c.cSockd }
func (c *RootController) RepoController() *repo.Controller     { return c.cRepo }
func (c *RootController) EngineController() *engine.Controller { return c.cEngine }
func (c *RootController) DevController() *dev.Controller       { return c.cDev }
func (c *RootController) UserController() *user.Controller     { return c.cUser }
