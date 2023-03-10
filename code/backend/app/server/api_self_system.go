package server

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

func (s *Server) selfSysAPI(rg *gin.RouterGroup) {

	rg.GET("/cabinet", s.X(s.ListCabinetSources))
	rg.GET("/datatable", s.X(s.ListDtableSources))
	rg.GET("/adapter", s.X(s.ListAdapters))
	rg.GET("/repo", s.X(s.ListRepoSources))
	rg.GET("/module", s.X(s.ListModules))
	rg.GET("/executor", s.X(s.ListExecutor))
	rg.GET("/invoker", s.X(s.ListInvokers))

}

func (s *Server) ListCabinetSources(ctx httpx.Request) {
	sources, err := s.cBasic.ListCabinetSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, sources, err)

}

func (s *Server) ListDtableSources(ctx httpx.Request) {
	sources, err := s.cBasic.ListDyndbSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, sources, err)
}

func (s *Server) ListAdapters(ctx httpx.Request) {
	resp := s.notz.ListAdapters()
	httpx.WriteJSON(ctx.Http, resp, nil)
}

func (s *Server) ListRepoSources(ctx httpx.Request) {
	resp, err := s.cBasic.ListRepoSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) ListExecutor(ctx httpx.Request) {
	resp, err := s.cEngine.ListExecutors(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Server) ListModules(ctx httpx.Request) {
	resp, err := s.cEngine.ListModules(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)

}

func (s *Server) ListInvokers(ctx httpx.Request) {

}
