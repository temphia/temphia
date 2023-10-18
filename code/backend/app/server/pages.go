package server

import (
	"embed"
	"fmt"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

// during dev we just proxy to dev vite server running otherwise serve files from build folder
func (s *Server) pages(z *gin.RouterGroup) {
	rfunc := s.pagesRoutes()

	z.GET("/pages", rfunc)
	z.GET("/pages/*files", rfunc)

}

const NoDEV = true

func (s *Server) pagesRoutes() gin.HandlerFunc {
	var proxy *httputil.ReverseProxy
	pserver := os.Getenv("TEMPHIA_DEV_PAGES_SERVER")

	if s.opts.BuildFS == nil {
		panic("BUILD FS")
	}

	if pserver != "" && !NoDEV {
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

	bfs := (s.opts.BuildFS.(embed.FS))

	return func(ctx *gin.Context) {
		path := strings.TrimSuffix(strings.TrimPrefix(ctx.Request.URL.Path, "/z/pages/"), "/")

		pitems := strings.Split(path, "/")
		lastpath := pitems[len(pitems)-1]

		if !strings.Contains(lastpath, ".") {
			path = path + ".html"
		}

		pp.Println("@FILE ==>", path)

		out, err := bfs.ReadFile(fmt.Sprintf("build/%s", path))
		if err != nil {
			pp.Println("@open_err", err.Error())
			return
		}

		httpx.WriteFile(path, out, ctx)
	}
}
