package server

import (
	"net"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"

	apiadmin "github.com/temphia/temphia/code/core/backend/app/server/api_admin"
	"github.com/temphia/temphia/code/core/backend/app/server/middleware"
	"github.com/temphia/temphia/code/core/backend/app/server/notz"
	"github.com/temphia/temphia/code/core/backend/controllers"
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
	"github.com/temphia/temphia/code/core/backend/xtypes/logx"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

type Options struct {
	App               xtypes.App
	GinEngine         *gin.Engine
	StaticHosts       map[string]string
	ResolveHostTenant func(host string) string
	RootController    *controllers.RootController
	RootHost          string
	TenantHostBase    string
	Port              string
}

type Server struct {
	app       xtypes.App
	ginEngine *gin.Engine
	admin     apiadmin.ApiAdmin
	log       logx.Service
	signer    service.Signer
	notz      *notz.Notz
	data      xtypes.DataBox
	port      string

	middleware *middleware.Middleware

	listener net.Listener

	// controllers

	cOperator            *operator.Controller
	cAuth                *authed.Controller
	cBasic               *basic.Controller
	cUser                *user.Controller
	cData                *data.Controller
	cCabinet             *cabinet.Controller
	cRepo                *repo.Controller
	cEngine              *engine.Controller
	cDev                 *dev.Controller
	cSockd               *sockd.Controller
	sockdConnIdGenerator *snowflake.Node
}

func New(opts Options) *Server {
	deps := opts.App.GetDeps()

	logsvc := deps.LogService().(logx.Service)
	signer := deps.Signer().(service.Signer)

	mware := &middleware.Middleware{
		Signer: signer,
		Logger: logsvc.GetServiceLogger("routes"),
	}

	root := opts.RootController

	plane := deps.ControlPlane().(xplane.ControlPlane)

	nz := notz.New(notz.NotzOptions{
		App:               opts.App,
		StaticHosts:       opts.StaticHosts,
		ResolveHostTenant: opts.ResolveHostTenant,
		RootHost:          opts.RootHost,
		TenantHostBase:    opts.TenantHostBase,
	})

	return &Server{
		ginEngine: opts.GinEngine,
		admin: apiadmin.New(apiadmin.Options{
			Admin:      root.AdminController(),
			MiddleWare: mware,
			Notz:       nz,
			Signer:     signer,
		}),
		log:    logsvc,
		signer: signer,
		port:   opts.Port,
		notz:   nz,
		data:   opts.App.Data(),

		middleware: mware,

		// controllers

		cOperator:            root.OperatorController(),
		cAuth:                root.AuthController(),
		cBasic:               root.BasicController(),
		cData:                root.DtableController(),
		cCabinet:             root.CabinetController(),
		cRepo:                root.RepoController(),
		cEngine:              root.EngineController(),
		cDev:                 root.DevController(),
		cUser:                root.UserController(),
		cSockd:               root.SockdController(),
		app:                  opts.App,
		sockdConnIdGenerator: plane.GetIdService().NewNode("temphia.sockd"),
	}
}

func (s *Server) Listen() error {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		panic(err)
	}
	s.listener = listener

	return s.ginEngine.RunListener(listener)
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func (s *Server) BuildRoutes() {

	if s.ginEngine == nil {

		s.ginEngine = gin.New()
		gin.SetMode(gin.DebugMode)

		s.ginEngine.Use(
			s.middleware.Log,
			gin.Recovery(),
		)
	}

	s.buildRoutes()
}
