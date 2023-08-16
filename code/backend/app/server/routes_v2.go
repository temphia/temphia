package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/temphia/temphia/code/backend/app/server/templates"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz/httpx"
)

func (s *Server) buildRoutes() {

	ginEngine := s.opts.GinEngine

	z := ginEngine.Group("/z")
	s.zRoutes(z)
	ginEngine.NoRoute(s.noRoute)
}

func (s *Server) zRoutes(z *gin.RouterGroup) {

	z.GET("/", s.asFile(templates.Root, "html"))
	z.GET("/portal", s.asFile(templates.Portal, "html"))

	z.GET("/auth", s.AuthIndex)
	s.assets(z.Group("/assets"))

	s.API(z.Group("/api/:tenant_id/v2/"))
}

func (s *Server) API(rg *gin.RouterGroup) {
	s.admin.API(rg.Group("/admin"))

	s.authserver.API(rg.Group("/auth"))

	s.apidata.API(rg.Group("/data"))

	s.apiroot.DataWSAPI(rg.Group("/data_ws"))

	s.apiroot.CabinetAPI(rg.Group("/cabinet"))

	s.ticketsAPI.Folder(rg.Group("/folder"))

	s.apiroot.DevAPI(rg.Group("/dev"))

	s.apiroot.EngineAPI(rg.Group("/engine"))

	s.apiself.API(rg.Group("/self"))

	s.apiroot.RepoAPI(rg.Group("/repo"))

	s.apiroot.AdapterEditorAPI(rg.Group("/adapter_editor"))

}

func (s *Server) noRoute(ctx *gin.Context) {

	if strings.HasPrefix(ctx.Request.URL.Path, "/z/") {
		pparts := strings.Split(ctx.Request.URL.Path, "/")

		switch pparts[2] {
		case "portal":
			ctx.Redirect(http.StatusFound, "/z/portal")
			return
		case "auth":
			ctx.Redirect(http.StatusFound, "/z/auth")
			return
		default:
			httpx.NotFound(ctx)
			return
		}
	}

	// s.notz.Serve(ctx)
}

func (s *Server) asFile(data []byte, ext string) func(ctx *gin.Context) {
	return s.middleware.AsFile(data, ext)
}

// fixme  =>  /z/extension/<name> ?
