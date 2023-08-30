package server

import (
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

// during dev we just proxy to dev vite server running otherwise serve files from build folder
func (s *Server) pages(z *gin.RouterGroup) {
	rfunc := s.pagesRoutes()

	z.GET("/pages", rfunc)
	z.GET("/pages/*files", rfunc)

}

func (s *Server) pagesRoutes() gin.HandlerFunc {
	var proxy *httputil.ReverseProxy
	pserver := os.Getenv("TEMPHIA_DEV_PAGES_SERVER")
	if pserver != "" {
		url, err := url.Parse(pserver)
		if err != nil {
			panic(err)
		}
		pp.Println("@using_dev_proxy", pserver)

		proxy = httputil.NewSingleHostReverseProxy(url)
		return func(ctx *gin.Context) {
			pp.Println("[PROXY]", ctx.Request.URL.String())
			proxy.ServeHTTP(ctx.Writer, ctx.Request)
		}

	}
	pp.Println("@not_using_dev_proxy")
	return func(ctx *gin.Context) {
		// fixme => serve build folder
		pp.Println("[SERVE]")
	}
}