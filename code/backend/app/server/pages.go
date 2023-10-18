package server

import (
	"io"
	"io/fs"
	"net/http"
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

	BuildFolder, err := fs.Sub(s.opts.BuildFS, "build")
	if err != nil {
		panic(err)
	}

	bfs := http.FS(BuildFolder)

	return func(ctx *gin.Context) {
		path := strings.TrimSuffix(strings.TrimPrefix(ctx.Request.URL.Path, "/z/pages/"), "/")

		pitems := strings.Split(path, "/")
		lastpath := pitems[len(pitems)-1]

		if !strings.Contains(lastpath, ".") {
			path = path + ".html"
		}

		pp.Println("@FILE ==>", path)

		hf, err := bfs.Open(path)
		if err != nil {
			pp.Println("@open_err", err.Error())
			return
		}

		out, err := io.ReadAll(hf)
		if err != nil {
			pp.Println("@read_err", err.Error())
			return
		}

		httpx.WriteFile(path, out, ctx)
	}
}
