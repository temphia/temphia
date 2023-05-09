package easypage

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/tidwall/gjson"
	"github.com/yuin/goldmark"
)

func (s *EasyPage) handle(ctx httpx.Context) {
	// path := "index"

	path := ctx.Http.Request.URL.Path
	if strings.HasPrefix(path, "/image/") {
		s.handleImg(ctx, strings.Replace(path, "/image/", "", 1))
		return
	}

	path = strings.TrimLeft(ctx.Http.Request.URL.Path, "/")
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

func (s *EasyPage) handleImg(ctx httpx.Context, file string) {
	// fixme => fixme from bprint folder
}

// private

func (s *EasyPage) kvGet(string) (string, error) {
	// val, err := s.ahandle.KvGet(pageKey(path))
	// if err != nil {
	// 	return "", err
	// }
	return "", nil

}

func (s *EasyPage) fetch(path string) ([]byte, error) {

	val, err := s.kvGet(path)
	if err != nil {
		return nil, err
	}

	switch gjson.Get(val, "type").String() {
	case "post":
		return s.processPost(path, val)
	default:
		return s.processPage(path, val)
	}

}

func (s *EasyPage) processPost(path, val string) ([]byte, error) {

	md := (gjson.Get(val, "code").String())

	var buf bytes.Buffer

	buf.WriteString(`
	<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1" />
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/spcss@0.9.0">
  </head>	
  <body>
  	<div class="marked">
	`)

	if err := goldmark.Convert([]byte(md), &buf); err != nil {
		return nil, err
	}

	buf.WriteString(`
			</div>
		<body>
	</html>`)

	out := buf.Bytes()

	s.pLock.Lock()
	s.pageCache[path] = out
	s.pLock.Unlock()

	return nil, nil
}

func (s *EasyPage) processPage(path, val string) ([]byte, error) {

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
	<script src="/z/assets/lib/tailwind.js"></script>
	
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
