package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) assets(rg *gin.RouterGroup) {

	rg.GET("/static/:file", s.serveStaticAssets())
	rg.GET("/build/:file", s.serveBuildAssets())
	rg.GET("/lib/:file", s.serveLibAssets())
	rg.GET("/public/:file", s.servePublicAssets())
	rg.GET("/system/:otype/:file", s.serveSystemAssets())
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
	return func(c *gin.Context) {
		/*

			ns_assets/icon_user_john11.png
			ns_assets/icon_plug_xyz.png
			ns_assets/icon_plug_agent_xyz_mno.png
			ns_assets/icon_ugroup_super_admin.png
			ns_assets/icon_tenant.png
			ns_assets/authed_background.png



		*/

	}
}

// private utils

func (s *Server) assetServe(folder string) func(c *gin.Context) {
	fs := http.FS(s.app.Data().AssetAdapter(folder))
	return func(c *gin.Context) {
		c.FileFromFS(c.Param("file"), fs)
	}
}
