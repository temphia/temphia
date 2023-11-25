package bdev

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/temphia/temphia/code/climux"
	"github.com/temphia/temphia/code/climux/bdev/core"
)

func init() {

	climux.Register(&climux.Action{
		Name: "bdev",
		Help: "app/bprint development related actions",
		Func: RunCLI2,
	})

}

func RunCLI2(cctx climux.Context) error {

	os.Args = cctx.Args

	cli := &BdevCLI{}
	ctx := kong.Parse(cli)

	bctx := core.BdevContext{
		ConfigFile: cli.BprintFile,
		KongCtx:    ctx,
	}

	return ctx.Run(bctx)

}
