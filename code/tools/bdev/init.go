package bdev

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/temphia/temphia/code/distro/climux"
)

func init() {

	climux.Register(&climux.CliAction{
		Name: "bdev",
		Help: "bprint development related actions",
		Func: RunCLI,
	})

}

func RunCLI(args []string) error {

	os.Args = args

	cli := &CLI{}
	ctx := kong.Parse(cli)
	cli.ctx = ctx

	return cli.Execute()
}
