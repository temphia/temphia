package main

import (
	"github.com/k0kubun/pp"
	_ "github.com/temphia/temphia/code/distro/common"
	"github.com/temphia/temphia/code/distro/demo"
)

func main() {
	pp.Println("Hello wold @ demo")

	pp.Println(demo.Main())

}
