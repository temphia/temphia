package easypage

import (
	"fmt"
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/tidwall/gjson"
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
		return processPost(path, val)
	default:
		return processPage(path, val)
	}

}

// private

func (s *EasyPage) build() {}
