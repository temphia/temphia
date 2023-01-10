package main

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/libx/xutils"
	_ "github.com/temphia/temphia/code/distro/common"
	"github.com/temphia/temphia/code/distro/demo"
)

func main() {
	pp.Println("Hello wold @ demo")
	os.Chdir("cmd/demo/")

	xutils.CreateIfNotExits("temphia-data/files")
	xutils.CreateIfNotExits("temphia-data/logs")
	xutils.CreateIfNotExits("temphia-data/pgdata")

	pp.Println(demo.Main())

}
