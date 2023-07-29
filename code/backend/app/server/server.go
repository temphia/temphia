package server

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"

	apiroot "github.com/temphia/temphia/code/backend/app/server/API"
	apiadmin "github.com/temphia/temphia/code/backend/app/server/API/admin"
	apiauth "github.com/temphia/temphia/code/backend/app/server/API/auth"
	apidata "github.com/temphia/temphia/code/backend/app/server/API/data"
	apiself "github.com/temphia/temphia/code/backend/app/server/API/self"

	"github.com/temphia/temphia/code/backend/app/server/API/middleware"
	"github.com/temphia/temphia/code/backend/app/server/API/tickets"
	"github.com/temphia/temphia/code/backend/controllers"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

var _ xtypes.Server = (*Server)(nil)

type Options struct {
	RootDomain     string
	RunnerDomain   string
	App            xtypes.App
	GinEngine      *gin.Engine
	RootController *controllers.RootController
	Port           string
}

type Server struct {
	opts     Options
	duckMode bool

	log    logx.Service
	signer service.Signer
	notz   httpx.AdapterHub

	listener net.Listener

	middleware *middleware.Middleware

	admin      apiadmin.ApiAdmin
	authserver *apiauth.Auth
	apiself    *apiself.Self
	apidata    *apidata.Data
	apiroot    apiroot.Server
	ticketsAPI *tickets.TicketAPI
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

	tktapi := tickets.New(mware, root)

	node := plane.GetIdService().NewNode("temphia.sockd")

	return &Server{
		opts:       opts,
		duckMode:   true,
		log:        logsvc,
		signer:     signer,
		notz:       nil,
		listener:   nil,
		middleware: mware,
		admin: apiadmin.New(apiadmin.Options{
			Admin:      root.AdminController(),
			MiddleWare: mware,
			Signer:     signer,
			TicketAPI:  tktapi,
		}),
		authserver: apiauth.New(root.AuthController(), signer),
		apiself:    apiself.New(signer, mware, nil, root, node),
		apidata:    apidata.New(mware, root.DtableController()),
		ticketsAPI: tickets.New(mware, root),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if s.duckMode {
		return
	}

}

func (s *Server) Listen() error {

	listener, err := net.Listen("tcp", s.opts.Port)
	if err != nil {
		panic(err)
	}

	s.listener = listener

	return http.Serve(listener, s)
}

func (s *Server) Close() error {
	return nil
}

func (s *Server) BuildRoutes() error {

	if s.opts.GinEngine == nil {

		s.opts.GinEngine = gin.New()
		gin.SetMode(gin.DebugMode)

		s.opts.GinEngine.Use(
			s.middleware.Log,
			gin.Recovery(),
		)
	}

	s.buildRoutes()

	return nil

}

func (s *Server) GetAdapterHub() any { return nil }
