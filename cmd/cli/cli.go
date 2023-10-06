package cli

import (
	"fmt"
	"os"

	"github.com/temphia/temphia/code/distro"
)

const helpText = `
Temphia is a platform for apps.

Usage:

        temphia <command> [arguments]

The commands are:
		{{ }}
        version     print Temphia version
        help        this help page

Use "temphia <command> help" for more information about a command.


`

func Run() {
	if len(os.Args) > 1 {
		if os.Args[1] == "help" || os.Args[1] == "--help" {
			fmt.Print(helpText)
			return
		}
	}

	err := distro.RunAppCLI(os.Args)
	if err != nil {
		panic(err)
	}
}
