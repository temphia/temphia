package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) serveAssets() func(c *gin.Context) {
	return func(c *gin.Context) {
		fs := http.FS(s.app.Data().AssetAdapter())
		c.FileFromFS(c.Param("file"), fs)
	}
}

func (s *Server) publicFile() func(c *gin.Context) {
	return func(c *gin.Context) {}
}

func (s *Server) systemAssets(rg *gin.RouterGroup) {

}

/*

	ns_assets/icon_user_john11.png
	ns_assets/icon_plug_xyz.png
	ns_assets/icon_plug_agent_xyz_mno.png
	ns_assets/icon_ugroup_super_admin.png
	ns_assets/icon_tenant.png
	ns_assets/authed_background.png


*/
