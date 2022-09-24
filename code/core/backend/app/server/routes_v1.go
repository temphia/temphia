package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/app/server/static"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

func (s *Server) buildRoutes() {
	z := s.ginEngine.Group("/z")
	s.zRoutes(z)
	s.ginEngine.NoRoute(s.noRoute)
}

func (s *Server) zRoutes(z *gin.RouterGroup) {

	z.GET("/", s.asFile(static.Root, "html"))
	z.GET("/start", s.asFile(static.Root, "html"))
	z.GET("/portal", s.asFile(static.Portal, "html"))
	z.GET("/operator", s.asFile(static.Operator, "html"))
	z.GET("/interface/:name", s.serveInterface)
	z.GET("/assets/static/:file", s.serveAssets())
	z.GET("/assets/public/:file", s.publicFile())
	s.systemAssets(z.Group("/assets/system/"))

	s.operatorAPI(z.Group("/operator/api"))

	s.API(z.Group("/api/:tenant_id/v1/"))
}

func (s *Server) API(rg *gin.RouterGroup) {
	s.admin.API(rg.Group("/admin"))
	s.authAPI(rg.Group("/auth"))
	s.dataAPI(rg.Group("/data"))
	s.cabinetAPI(rg.Group("/cabinet"))
	s.devAPI(rg.Group("/dev"))
	s.engineAPI(rg.Group("/engine"))
	s.selfAPI(rg.Group("/self"))
	s.repoAPI(rg.Group("/repo"))
	s.sockdAPI(rg.Group("/sockd"))
	s.userAPI(rg.Group("/user"))

	s.cabinetTktAPI(rg.Group("/tkt/cabinet"))
	s.dataTktAPI(rg.Group("/tkt/data"))

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
		case "operator":
			ctx.Redirect(http.StatusFound, "/z/operator")
			return
		default:
			httpx.NotFound(ctx)
			return
		}
	}

	s.notz.Serve(ctx)
}
