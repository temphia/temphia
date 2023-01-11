package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func (s *Server) assets(rg *gin.RouterGroup) {

	rg.GET("/static/:file", s.serveStaticAssets())
	rg.GET("/build/:file", s.serveBuildAssets())
	rg.GET("/lib/*file", s.serveLibAssets())
	rg.GET("/public/:file", s.servePublicAssets())
	rg.GET("/system/:tenant_id/:otype/:name", s.serveSystemAssets())
}

// static assets that are "static" not build files
func (s *Server) serveStaticAssets() func(c *gin.Context) {
	return s.assetServe("static")
}

func (s *Server) serveBuildAssets() func(c *gin.Context) {
	return s.assetServe("build")
}

func (s *Server) serveLibAssets() func(c *gin.Context) {
	return s.assetServe("lib")
}

// bashed on public folder in tenant's root cabinet source
func (s *Server) servePublicAssets() func(c *gin.Context) {
	return func(c *gin.Context) {
		s.notz.ServePublic(c, c.Param("file"))
	}
}

func (s *Server) serveSystemAssets() func(c *gin.Context) {

	deps := s.app.GetDeps()
	cab := deps.Cabinet().(store.CabinetHub)

	return func(c *gin.Context) {
		csource := cab.Default(c.Param("tenant_id"))
		out, err := csource.GetBlob(c.Request.Context(), "system", fmt.Sprintf("%s_%s.png", c.Param("otype"), c.Param("name")))
		if err != nil {
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}

		httpx.WriteBinary(c, out)
	}
}

// private utils

func (s *Server) assetServe(folder string) func(c *gin.Context) {
	fs := http.FS(s.app.Data().AssetAdapter(folder))
	return func(c *gin.Context) {
		c.FileFromFS(c.Param("file"), fs)
	}
}
