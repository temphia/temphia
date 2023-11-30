package server

import (
	"fmt"
	"io/fs"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/app/data"
	"github.com/temphia/temphia/code/backend/xtypes/store/fdatautil"
)

func (s *Server) assets(rg *gin.RouterGroup) {

	assetFS, err := fs.Sub(data.DataFS, "assets")
	if err != nil {
		panic(err)
	}

	libsFS, err := fs.Sub(data.DataFS, "libs")
	if err != nil {
		panic(err)
	}

	rg.StaticFS("/static/", http.FS(assetFS))
	rg.StaticFS("/build/", http.FS(s.opts.BuildFS))
	rg.StaticFS("/lib/", http.FS(libsFS))

	rg.GET("/public/:file", s.servePublicAssets)
	rg.GET("/system/:tenant_id/:otype/:name", s.serveSystemAssets)
}

// bashed on public folder in tenant's root cabinet source
func (s *Server) servePublicAssets(c *gin.Context) {

	file := c.Param("file")

	data, err := s.cabhub.GetFile(c.Request.Context(), c.Param("tenant_id"), path.Join("public", file))
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	fdatautil.WriteAndClose(c.Writer, file, data)
}

func (s *Server) serveSystemAssets(c *gin.Context) {

	file := fmt.Sprintf("%s_%s.png", c.Param("otype"), c.Param("name"))

	data, err := s.cabhub.GetFile(c.Request.Context(), c.Param("tenant_id"), path.Join("system", file))
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	fdatautil.WriteAndClose(c.Writer, file, data)

}
