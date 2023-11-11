package self

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (s *Self) selfSysAPI(rg *gin.RouterGroup) {

	rg.GET("/cabinet", s.X(s.ListCabinetSources))
	rg.GET("/datatable", s.X(s.ListDtableSources))
	rg.GET("/repo", s.X(s.ListRepoSources))
	rg.GET("/module", s.X(s.ListModules))
	rg.GET("/executor", s.X(s.ListExecutor))
	rg.GET("/invoker", s.X(s.ListInvokers))

}

func (s *Self) ListCabinetSources(ctx httpx.Request) {
	sources, err := s.cBasic.ListCabinetSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, sources, err)

}

func (s *Self) ListDtableSources(ctx httpx.Request) {
	sources, err := s.cBasic.ListDyndbSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, sources, err)
}

func (s *Self) ListRepoSources(ctx httpx.Request) {
	resp, err := s.cBasic.ListRepoSources(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Self) ListExecutor(ctx httpx.Request) {
	resp, err := s.cEngine.ListExecutors(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)
}

func (s *Self) ListModules(ctx httpx.Request) {
	resp, err := s.cEngine.ListModules(ctx.Session)
	httpx.WriteJSON(ctx.Http, resp, err)

}

func (s *Self) ListInvokers(ctx httpx.Request) {

}
