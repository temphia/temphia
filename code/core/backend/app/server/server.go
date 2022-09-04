package server

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/controllers"
	"github.com/temphia/temphia/code/core/backend/xtypes"
)

type Options struct {
	App       xtypes.App
	GinEngine *gin.Engine
}

type Server struct {
	app       xtypes.App
	ginEngine *gin.Engine
}

func New(opts Options) *Server {

	ctrls := controllers.New(opts.App, nil)

	pp.Println(ctrls)

	return &Server{
		app:       opts.App,
		ginEngine: nil,
	}
}

func (s *Server) BuildRoutes() {
	s.buildRoutes()
}
