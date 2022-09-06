package server

import (
	"github.com/gin-gonic/gin"
	apiadmin "github.com/temphia/temphia/code/core/backend/app/server/api_admin"
	"github.com/temphia/temphia/code/core/backend/app/server/ginlogger"
	"github.com/temphia/temphia/code/core/backend/app/server/notz"
	"github.com/temphia/temphia/code/core/backend/controllers"
	"github.com/temphia/temphia/code/core/backend/controllers/authed"
	"github.com/temphia/temphia/code/core/backend/controllers/basic"
	"github.com/temphia/temphia/code/core/backend/controllers/cabinet"
	"github.com/temphia/temphia/code/core/backend/controllers/dtable"
	"github.com/temphia/temphia/code/core/backend/controllers/operator"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
)

type Options struct {
	App               xtypes.App
	GinEngine         *gin.Engine
	StaticHosts       map[string]string
	ResolveHostTenant func(host string) string
	RootHost          string
	TenantHostBase    string
}

type Server struct {
	app       xtypes.App
	ginEngine *gin.Engine
	admin     apiadmin.ApiAdmin
	log       logx.Service
	signer    service.Signer
	notz      notz.Notz
	data      xtypes.DataBox

	// controllers

	cOperator *operator.Controller
	cAuth     *authed.Controller
	cBasic    *basic.Controller
	cDtable   *dtable.Controller
	cCabinet  *cabinet.Controller
}

func New(opts Options) *Server {
	deps := opts.App.GetDeps()

	ctrls := controllers.New(opts.App)
	logsvc := deps.LogService().(logx.Service)
	signer := deps.Signer().(service.Signer)

	return &Server{
		app:       opts.App,
		ginEngine: opts.GinEngine,
		admin:     apiadmin.New(ctrls.AdminController()),
		log:       logsvc,
		signer:    signer,
		notz: notz.New(notz.NotzOptions{
			App:               opts.App,
			StaticHosts:       opts.StaticHosts,
			ResolveHostTenant: opts.ResolveHostTenant,
			RootHost:          opts.RootHost,
			TenantHostBase:    opts.TenantHostBase,
		}),
		data: opts.App.Data(),

		// controllers

		cOperator: nil,
		cAuth:     nil,
		cBasic:    nil,
		cDtable:   nil,
		cCabinet:  nil,
	}
}

func (s *Server) BuildRoutes() {

	if s.ginEngine == nil {

		s.ginEngine = gin.New()
		gin.SetMode(gin.DebugMode)

		s.ginEngine.Use(
			ginlogger.Logger(s.log.GetAppLogger(), "GIN_APP"),
			gin.Recovery(),
		)
	}

	s.buildRoutes()
}
