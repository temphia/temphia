package data

import (
	"github.com/k0kubun/pp"
	"gitlab.com/mr_balloon/golib"
)

var (
	defaultDataBox *DataBox
)

func lazyInit() {
	var frontendBuild = "code/core/frontend/public"
	if exits, _ := golib.FileExists(frontendBuild); !exits {
		frontendBuild = ""
	}

	pp.Println("Using overlay assets from", frontendBuild)

	defaultDataBox = &DataBox{
		dataOverlay:        "",
		buildAssetsOverlay: frontendBuild,
		fs:                 dataDir,
	}
}
