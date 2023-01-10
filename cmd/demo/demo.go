package main

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/libx/xutils"
	_ "github.com/temphia/temphia/code/distro/common"
	"github.com/temphia/temphia/code/distro/demo"
)

func main() {

	fmt.Println("Starting temphia demo")
	os.Chdir("cmd/demo/")
	xutils.CreateIfNotExits("temphia-data/files")
	xutils.CreateIfNotExits("temphia-data/logs")
	xutils.CreateIfNotExits("temphia-data/pgdata")

	democli := demo.NewCLI()

	if xutils.FileExists("temphia-data/pgdata/data", "postmaster.pid") {
		fmt.Println("looks like another demo instance is running or last one did not close you might have to clear lock with 'temphia-demo clear-lock'")
	}

	pp.Println(democli.Execute())

	if xutils.FileExists("temphia-data/pgdata/data", "postmaster.pid") {
		fmt.Println("Looks like db did not close properly, might need to clear lock 'temphia-demo clear-lock'")
	}

}
