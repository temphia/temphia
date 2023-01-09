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

	xutils.CreateIfNotExits("tmp/files")
	xutils.CreateIfNotExits("tmp/logs")
	xutils.CreateIfNotExits("tmp/pgdata")

	pp.Println(demo.Main())

}
