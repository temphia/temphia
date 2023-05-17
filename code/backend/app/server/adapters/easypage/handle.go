package easypage

import (
	"fmt"
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

func (s *EasyPage) handle(ctx httpx.Context) {
	path := strings.TrimLeft(ctx.Http.Request.URL.Path, "/")

	if path == "" {
		path = "index"
	} else if strings.HasPrefix(path, "/image/") {
		path = fmt.Sprintf("image_%s", strings.Replace(path, "/image/", "", 1))
	}

	s.filecache.Get(ctx.Http.Writer, path)
}
