package server

import (
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"

	apiroot "github.com/temphia/temphia/code/backend/app/server/API"
	apiadmin "github.com/temphia/temphia/code/backend/app/server/API/admin"
	apiauth "github.com/temphia/temphia/code/backend/app/server/API/auth"
	apidata "github.com/temphia/temphia/code/backend/app/server/API/data"
	apiself "github.com/temphia/temphia/code/backend/app/server/API/self"
	"github.com/temphia/temphia/code/backend/app/server/notz"

	"github.com/temphia/temphia/code/backend/app/server/API/middleware"
	"github.com/temphia/temphia/code/backend/app/server/API/tickets"
	"github.com/temphia/temphia/code/backend/controllers"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
	"github.com/temphia/temphia/code/backend/xtypes/xserver"
)

var _ xserver.Server = (*Server)(nil)

type Options struct {
	RootDomain     string
	RunnerDomain   string
	LocalSocket    string
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
	cabhub store.CabinetHub

	listener   net.Listener
	ldListener net.Listener

	notz xnotz.Notz

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
		opts:     opts,
		duckMode: true,
		log:      logsvc,
		signer:   signer,

		notz: notz.New(opts.App),

		listener: nil,

		middleware: mware,
		admin: apiadmin.New(apiadmin.Options{
			Admin:      root.AdminController(),
			MiddleWare: mware,
			Signer:     signer,
			TicketAPI:  tktapi,
		}),
		authserver: apiauth.New(root.AuthController(), signer, opts.App.TenantId()),
		apiself:    apiself.New(signer, mware, nil, root, node),
		apidata:    apidata.New(mware, root.DtableController()),
		ticketsAPI: tickets.New(mware, root),
		apiroot:    apiroot.New(signer, mware, root, node),
		cabhub:     deps.Cabinet().(store.CabinetHub),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	pp.Println("@ServeHTTP", req.URL.String())
	pp.Println(req.URL.Path)

	if s.duckMode {
		pp.Println("@Ducked", req.URL.String())
		return
	}

	if strings.HasPrefix(req.URL.Path, "/z/") {
		pp.Println("@z")
		s.opts.GinEngine.ServeHTTP(w, req)
		return
	}

	host := strings.Split(req.Host, ":")[0]
	if strings.HasSuffix(host, s.opts.RunnerDomain) && strings.Contains(host, "-n-") {
		prefix := strings.Replace(host, fmt.Sprintf(".%s", s.opts.RunnerDomain), "", 1)

		ids := strings.Split(prefix, "-n-")

		if len(ids) == 2 {
			pp.Println("runner_prefix", ids)
			s.notz.HandleAgent(xnotz.Context{
				Writer:   w,
				Request:  req,
				TenantId: "",
				PlugId:   ids[0],
				AgentId:  ids[1],
			})

			return
		}

	}

	s.opts.GinEngine.ServeHTTP(w, req)
}

func (s *Server) Listen() error {

	err := s.BuildRoutes()
	if err != nil {
		return err
	}

	err = s.notz.Start()
	if err != nil {
		return err
	}

	err = s.localdoor()
	if err != nil {
		return err
	}

	pp.Println("@port", s.opts.Port)

	listener, err := net.Listen("tcp", s.opts.Port)
	if err != nil {
		panic(err)
	}

	s.listener = listener

	s.duckMode = false
	return http.Serve(listener, s)
}

func (s *Server) Close() error {
	if s.ldListener != nil {
		s.ldListener.Close()
	}

	if s.listener != nil {
		return s.listener.Close()
	}

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

func (s *Server) localdoor() error {

	os.Remove(s.opts.LocalSocket)

	l, err := net.Listen("unix", s.opts.LocalSocket)
	if err != nil {
		return err
	}

	s.ldListener = l

	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				log.Fatal("accept error:", err)
			}

			func(c net.Conn) {
				defer c.Close()

				buf := make([]byte, 512)
				nr, err := c.Read(buf)
				if err != nil {
					return
				}

				data := buf[0:nr]
				println("Server got:", string(data))
				_, err = c.Write(data)
				if err != nil {
					log.Fatal("Write: ", err)
				}

			}(c)

		}

	}()

	return nil

}
