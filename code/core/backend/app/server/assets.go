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
