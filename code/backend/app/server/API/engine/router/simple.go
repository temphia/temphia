package router

import (
	"fmt"
	"strings"
)

func SimpleRoutePick(config *RouteConfig, path, method string) *RouteResponse {

	var selected *RouteItem
	var spath []string

	for i := range config.Items {

		croute := &config.Items[i]

		pick := func() {
			selected = croute
			spath = strings.Split(croute.Path, "/")
		}

		if method != croute.Method || method != "ANY" {
			continue
		}

		if !strings.HasPrefix(path, croute.Path) {
			continue
		}

		if selected == nil {
			pick()
			continue
		}

		slen := len(spath)
		clen := len(strings.Split(croute.Path, "/"))

		if slen < clen {
			pick()
		}

	}

	if selected != nil && selected.File == "" {
		paths := strings.Split(path, "/")
		selected.File = paths[len(paths)-1]

		if selected.ApppendHTML && !strings.Contains(selected.File, ".") {
			selected.File = fmt.Sprintf("%s.html", selected.File)
		}
	}

	resp := selected.Copy()
	return resp
}
