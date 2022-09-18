package server

import (
	"net"

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
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
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
	notz      notz.Notz
	data      xtypes.DataBox
	port      string

	middleware *middleware.Middleware

	// controllers

	cOperator *operator.Controller
	cAuth     *authed.Controller
	cBasic    *basic.Controller
	cData     *data.Controller
	cCabinet  *cabinet.Controller
	cRepo     *repo.Controller
	cEngine   *engine.Controller
	cDev      *dev.Controller
}

func New(opts Options) *Server {
	deps := opts.App.GetDeps()

	logsvc := deps.LogService().(logx.Service)
	signer := deps.Signer().(service.Signer)

	mware := &middleware.Middleware{
		Signer: signer,
		Logger: logsvc.GetServiceLogger("routes"),
	}

	return &Server{
		ginEngine: opts.GinEngine,
		admin: apiadmin.New(apiadmin.Options{
			Admin:      opts.RootController.AdminController(),
			MiddleWare: mware,
		}),
		log:    logsvc,
		signer: signer,
		port:   opts.Port,
		notz: notz.New(notz.NotzOptions{
			App:               opts.App,
			StaticHosts:       opts.StaticHosts,
			ResolveHostTenant: opts.ResolveHostTenant,
			RootHost:          opts.RootHost,
			TenantHostBase:    opts.TenantHostBase,
		}),
		data: opts.App.Data(),

		middleware: mware,

		// controllers

		cOperator: nil,
		cAuth:     nil,
		cBasic:    nil,
		cData:     nil,
		cCabinet:  nil,
		cRepo:     nil,
		cEngine:   nil,
		cDev:      nil,
	}
}

func (s *Server) ListenHTTP() error {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		panic(err)
	}

	return s.ginEngine.RunListener(listener)
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
