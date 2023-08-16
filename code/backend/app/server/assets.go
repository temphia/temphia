package server

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/app/data"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz/httpx"
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

	out, err := s.cabhub.GetBlob(c.Request.Context(), c.Param("tenant_id"), "public", c.Param("file"))
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	httpx.WriteBinary(c, out)

}

func (s *Server) serveSystemAssets(c *gin.Context) {

	out, err := s.cabhub.GetBlob(c.Request.Context(), c.Param("tenant_id"), "system", fmt.Sprintf("%s_%s.png", c.Param("otype"), c.Param("name")))
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	httpx.WriteBinary(c, out)

}
