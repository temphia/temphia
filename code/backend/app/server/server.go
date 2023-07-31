package server

import (
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"

	apiroot "github.com/temphia/temphia/code/backend/app/server/API"
	apiadmin "github.com/temphia/temphia/code/backend/app/server/API/admin"
	apiauth "github.com/temphia/temphia/code/backend/app/server/API/auth"
	apidata "github.com/temphia/temphia/code/backend/app/server/API/data"
	apiself "github.com/temphia/temphia/code/backend/app/server/API/self"
	"github.com/temphia/temphia/code/backend/app/server/agent"

	"github.com/temphia/temphia/code/backend/app/server/API/middleware"
	"github.com/temphia/temphia/code/backend/app/server/API/tickets"
	"github.com/temphia/temphia/code/backend/controllers"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

var _ xtypes.Server = (*Server)(nil)

type Options struct {
	RootDomain     string
	RunnerDomain   string
	TenantId       string
	App            xtypes.App
	GinEngine      *gin.Engine
	RootController *controllers.RootController
	Port           string
	BuildFS        fs.FS
}

type Server struct {
	opts     Options
	duckMode bool

	log    logx.Service
	signer service.Signer
	notz   httpx.AdapterHub
	cabhub store.CabinetHub

	listener net.Listener

	agentServer *agent.AgentServer

	middleware *middleware.Middleware

	admin      apiadmin.ApiAdmin
	authserver *apiauth.Auth
	apiself    *apiself.Self
	apidata    *apidata.Data
	apiroot    *apiroot.Server
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
		opts:        opts,
		duckMode:    true,
		log:         logsvc,
		signer:      signer,
		notz:        nil,
		listener:    nil,
		agentServer: agent.New(nil, opts.App),

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
		apiroot:    apiroot.New(signer, mware, root, node),
		cabhub:     deps.Cabinet().(store.CabinetHub),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if s.duckMode {
		return
	}

	if strings.HasPrefix(req.URL.Path, "/z/") {
		s.opts.GinEngine.ServeHTTP(w, req)
		return
	}

	host := strings.Split(req.Host, ":")[0]
	if strings.HasSuffix(host, s.opts.RunnerDomain) {
		prefix := strings.Replace(host, fmt.Sprintf(".%s", s.opts.RunnerDomain), "", 1)

		ids := strings.Split(prefix, "-n-")

		pp.Println("runner_prefix", ids)

		s.agentServer.Render(agent.Context{
			Writer:   w,
			Request:  req,
			TenantId: "",
			PlugId:   ids[0],
			AgentId:  ids[1],
		})

		return
	}

	s.opts.GinEngine.ServeHTTP(w, req)
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
