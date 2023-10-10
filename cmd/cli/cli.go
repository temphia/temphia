package cli

import (
	"fmt"
	"os"
	"text/template"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/distro/climux"
)

const HelpTemplate = `
Temphia is a platform for apps.

Usage:

        temphia <command> [arguments]

The commands are:{{range $index, $element := .}}
    {{$index}}	{{$element.Help}}{{end}}
    version     print Temphia version
    help        this help page

Use "temphia <command> help" for more information about a command.


`

func Run() {

	if len(os.Args) > 1 {
		if os.Args[1] == "help" || os.Args[1] == "--help" {
			PrintHelpText()
			return
		}

		if os.Args[1] == "version" || os.Args[1] == "--version" {
			fmt.Printf("temphia %s", xtypes.Version)
			return
		}

	}

	clis := climux.GetRegistry()
	acli, ok := clis[os.Args[1]]
	if ok {
		return
	}

	err := acli.Func(os.Args[1:])
	if err != nil {
		panic(err)
	}
}

func PrintHelpText() {
	clis := climux.GetRegistry()

	tpl, err := template.New("").Parse(HelpTemplate)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, clis)
	if err != nil {
		panic(err)
	}

}
