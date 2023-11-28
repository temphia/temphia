package api_server

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/app/server/API/dev"
	engineapi "github.com/temphia/temphia/code/backend/app/server/API/engine"
	"github.com/temphia/temphia/code/backend/app/server/middleware"
	"github.com/temphia/temphia/code/backend/controllers"
	"github.com/temphia/temphia/code/backend/controllers/cabinet"
	"github.com/temphia/temphia/code/backend/controllers/engine"
	"github.com/temphia/temphia/code/backend/controllers/repo"
	"github.com/temphia/temphia/code/backend/controllers/sockd"
	"github.com/temphia/temphia/code/backend/xtypes/service"
)

type Server struct {

	// controllers
	cCabinet *cabinet.Controller
	cRepo    *repo.Controller
	cEngine  *engine.Controller

	cSockd *sockd.Controller

	middleware *middleware.Middleware
	idNode     *snowflake.Node // sockdConnIdGenerator
	signer     service.Signer

	engineAPI *engineapi.EngineAPI
	devAPI    *dev.DevAPI
}

func New(signer service.Signer, mw *middleware.Middleware, rc *controllers.RootController, idNode *snowflake.Node) *Server {
	ec := rc.EngineController()

	return &Server{
		signer:     signer,
		middleware: mw,
		cCabinet:   rc.CabinetController(),
		cRepo:      rc.RepoController(),
		cEngine:    ec,
		cSockd:     rc.SockdController(),
		idNode:     idNode,
		engineAPI:  engineapi.New(ec, mw, signer),
		devAPI:     dev.New(signer, rc.DevController()),
	}
}

func (s *Server) DevAPI(rg *gin.RouterGroup) {
	s.devAPI.DevAPI(rg)
}
