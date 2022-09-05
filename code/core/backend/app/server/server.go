package server

import (
	"github.com/gin-gonic/gin"
	apiadmin "github.com/temphia/temphia/code/core/backend/app/server/api_admin"
	"github.com/temphia/temphia/code/core/backend/app/server/ginlogger"
	"github.com/temphia/temphia/code/core/backend/controllers"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx"
)

type Options struct {
	App       xtypes.App
	GinEngine *gin.Engine
}

type Server struct {
	app       xtypes.App
	ginEngine *gin.Engine
	admin     apiadmin.ApiAdmin
}

func New(opts Options) *Server {

	ctrls := controllers.New(opts.App)

	return &Server{
		app:       opts.App,
		ginEngine: opts.GinEngine,
		admin:     apiadmin.New(ctrls.AdminController()),
	}
}

func (s *Server) BuildRoutes() {

	if s.ginEngine == nil {
		ls := s.app.GetDeps().LogService().(logx.Service)

		s.ginEngine = gin.New()
		gin.SetMode(gin.DebugMode)

		s.ginEngine.Use(
			ginlogger.Logger(ls.GetAppLogger(), "GIN_APP"),
			gin.Recovery(),
		)
	}

	s.buildRoutes()
}
