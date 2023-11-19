package bdev

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/temphia/temphia/code/climux"
)

func init() {

	climux.Register(&climux.Action{
		Name: "bdev",
		Help: "app/bprint development related actions",
		Func: RunCLI,
	})

}

func RunCLI(cctx climux.Context) error {

	os.Args = cctx.Args

	cli := &CLI{}
	ctx := kong.Parse(cli)
	cli.ctx = ctx

	return cli.Run()
}

func RunCLI2(cctx climux.Context) error {

	os.Args = cctx.Args

	cli := &BdevCLI{}
	//ctx := kong.Parse(cli)

	return cli.Run()
}
