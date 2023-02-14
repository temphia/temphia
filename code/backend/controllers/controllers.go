package controllers

import (
	"github.com/temphia/temphia/code/backend/controllers/admin"
	"github.com/temphia/temphia/code/backend/controllers/authed"
	"github.com/temphia/temphia/code/backend/controllers/basic"
	"github.com/temphia/temphia/code/backend/controllers/cabinet"
	"github.com/temphia/temphia/code/backend/controllers/data"
	"github.com/temphia/temphia/code/backend/controllers/dev"
	"github.com/temphia/temphia/code/backend/controllers/engine"
	"github.com/temphia/temphia/code/backend/controllers/operator"
	"github.com/temphia/temphia/code/backend/controllers/repo"
	"github.com/temphia/temphia/code/backend/controllers/sockd"
	"github.com/temphia/temphia/code/backend/controllers/user"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
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
	seq := cplane.GetIdService()

	corehub := deps.CoreHub().(store.CoreHub)
	pacman := deps.RepoHub().(repox.Hub)
	signer := deps.Signer().(service.Signer)
	cab := deps.Cabinet().(store.CabinetHub)
	dynhub := deps.DataHub().(dyndb.DataHub)
	egine := deps.Engine().(etypes.Engine)
	logservice := deps.LogService().(logx.Service)
	pstate := deps.PlugKV().(store.PlugStateKV)

	return &RootController{
		cAdmin:   admin.New(pacman, cplane, corehub, signer, dynhub, cab, pstate, logservice.GetLogProxy()),
		cAuth:    authed.New(corehub, signer, seq),
		cBasic:   basic.New(corehub, cab, dynhub, pacman, signer),
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
		cEngine: engine.New(egine, signer, corehub),
		cUser:   user.New(corehub),
		cSockd:  sockd.New(deps.SockdHub().(sockdx.Hub).GetSockd()),
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
