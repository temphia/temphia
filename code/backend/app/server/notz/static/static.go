package static

import (
	"fmt"
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

/*

	"/" => index.html [root]
	"/about" => about.html [direct]
	"contrib/*" => contrib1/* [sub wildcard]
	"sub/*" => sub.html [sub to one file]
	"*" => not_found.html [not found]
	"*" => * [all to wildcard]
	"sub3/*" => * [all sub2 to wildcard]


*/

func ExtractPath(path string, agent *entities.Agent) (string, string) {

	file := ""

	if path == "/" {
		file = "index.html"
	} else {
		file = path
	}

	// exact match
	_file, ok := agent.WebFiles[path]
	if ok {
		return "", _file
	}

	// sub/* => mnop/*
	// sub/* => xyz.html
	for key, val := range agent.WebFiles {

		if key == "*" || !strings.HasSuffix(key, "*") {
			continue
		}

		prefix := strings.TrimSuffix(key, "*")
		if !strings.HasPrefix(strings.TrimPrefix(file, "/"), prefix) {
			continue
		}

		if !strings.HasSuffix(val, "*") {
			return prefix, val
		}

		val = strings.TrimSuffix(val, "*")

		paths := strings.Split(fmt.Sprintf("%s/%s", strings.TrimPrefix(prefix, file), val), "/")
		return strings.Join(paths[:len(paths)-1], "/"), paths[len(paths)-1]
	}

	// not found match

	_file, ok = agent.WebFiles["*"]
	if ok {
		if _file == "*" {
			return "", file

		}

		if strings.Contains(_file, "*") {
			return strings.Replace(_file, "*", "", 1), file
		}

		return "", file
	}

	return "", ""

}
