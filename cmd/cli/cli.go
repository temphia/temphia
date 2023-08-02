package cli

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/distro"
)

func Run() {
	pp.Println("@i_am_cli")

	err := distro.RunAppCLI(os.Args)
	if err != nil {
		panic(err)
	}
}
