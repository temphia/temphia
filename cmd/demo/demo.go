package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	_ "github.com/temphia/temphia/code/distro/common"
	"github.com/temphia/temphia/code/distro/demo"
)

func main() {

	fmt.Println("Starting temphia demo")

	democli := demo.NewCLI()

	pp.Println(democli.Execute())

}
