package easypage

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/tidwall/gjson"
)

func (s *EasyPage) handle(ctx httpx.Context) {
	// path := "index"

	path := strings.TrimLeft(ctx.Http.Request.URL.Path, "/")
	if path == "" {
		path = "index"
	}

	data, ok := s.pageCache[path]
	if ok {
		ctx.Http.Data(http.StatusOK, "text/html", data)
		return
	}

	out, err := s.fetch(path)
	if err != nil {
		return
	}

	ctx.Http.Data(http.StatusOK, "text/html", out)
}

func (s *EasyPage) fetch(path string) ([]byte, error) {
	val, err := s.ahandle.KvGet(pageKey(path))
	if err != nil {
		return nil, err
	}

	htmlArray := gjson.Get(val, "gen_html.0").String()

	html := (gjson.Get(htmlArray, "html").String())
	css := (gjson.Get(htmlArray, "css").String())

	var buf bytes.Buffer

	buf.WriteString(`
	<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1" />
	<script src="https://cdn.tailwindcss.com"></script>
	
    <script defer src="fixme.js"></script>
	<link rel="stylesheet" href="fixme.css" />
  </head>	
	`)

	buf.WriteString(html)
	buf.WriteString("<style>")
	buf.WriteString(css)
	buf.WriteString("</style>")
	buf.WriteString("</html>")

	out := buf.Bytes()

	s.pLock.Lock()
	s.pageCache[path] = out
	s.pLock.Unlock()

	return out, nil
}
