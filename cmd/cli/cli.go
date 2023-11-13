package cli

import (
	"fmt"
	"os"

	"github.com/temphia/temphia/code/backend/xtypes"

	"github.com/temphia/temphia/code/distro/climux"

	_ "github.com/temphia/temphia/code/distro"
	_ "github.com/temphia/temphia/code/ebrowser"
	_ "github.com/temphia/temphia/code/tools/bdev"
	_ "github.com/temphia/temphia/code/tools/repobuild"

	_ "github.com/temphia/temphia/code/backend/stores/localfs"
	_ "github.com/temphia/temphia/code/backend/stores/upperdb/vendors/sqlite"

	// repo provider

	_ "github.com/temphia/temphia/code/backend/services/pacman/provider/github"
)

func Run() {

	if len(os.Args) == 1 {
		os.Args = []string{os.Args[0], climux.DefaultCLI}
	}

	if os.Args[1] == "help" || os.Args[1] == "--help" {
		PrintHelpText()
		return
	}

	if os.Args[1] == "version" || os.Args[1] == "--version" {
		fmt.Printf("temphia %s", xtypes.Version)
		return
	}

	clis := climux.GetRegistry()
	acli, ok := clis[os.Args[1]]
	if !ok {
		fmt.Println("not found cli :", os.Args)
		os.Exit(1)
		return
	}

	err := acli.Func(climux.Context{
		Args: os.Args[1:],
		R:    clis,
	})
	if err != nil {
		panic(err)
	}
}

func PrintHelpText() {
	clis := climux.GetRegistry()

	fmt.Printf(`Temphia is a platform for apps.
Usage:
	
	temphia <command> [arguments]

The commands are:
`)

	for _, v := range clis {
		fmt.Printf("\t%s  \t\t%s\n", v.Name, v.Help)
	}

	fmt.Printf(`	version		print Temphia version
	help  		this help page

Use "temphia <command> help" for more information about a command.
`)
}
