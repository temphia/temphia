package server

import (
	"embed"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

// during dev we just proxy to dev vite server running otherwise serve files from build folder
func (s *Server) pages(z *gin.RouterGroup) {
	rfunc := s.pagesRoutes()

	z.GET("/pages", rfunc)
	z.GET("/pages/*files", rfunc)

}

const NoPreBuildFiles = false

func (s *Server) pagesRoutes() gin.HandlerFunc {
	var proxy *httputil.ReverseProxy
	pserver := os.Getenv(xtypes.EnvDevPageServer)

	if s.opts.BuildFS == nil {
		panic("BUILD FS")
	}

	if pserver != "" && !NoPreBuildFiles {
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

		ppath := strings.TrimSuffix(strings.TrimPrefix(ctx.Request.URL.Path, "/z/pages"), "/")

		if ppath == "" {
			ppath = "index.html"
		}

		pitems := strings.Split(ppath, "/")
		lastpath := pitems[len(pitems)-1]

		if !strings.Contains(lastpath, ".") {
			ppath = ppath + ".html"
		}

		pp.Println("@FILE ==>", ppath)

		out, err := bfs.ReadFile(path.Join("build", ppath))
		if err != nil {
			pp.Println("@open_err", err.Error())
			return
		}

		httpx.WriteFile(ppath, out, ctx)
	}
}
