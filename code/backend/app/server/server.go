package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers"
	"github.com/temphia/temphia/code/backend/xtypes"
)

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
}

func New(opts Options) *Server {
	return &Server{
		opts:     opts,
		duckMode: true,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if s.duckMode {
		return
	}

}

//  /z/extension/<name>
