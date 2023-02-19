package main

import (
	"github.com/alecthomas/kong"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/tools/bdev"
)

type CLI struct {
	Demo       struct{} `cmd:"" help:"demo cli stuff."`
	Bdev       bdev.CLI `cmd:"" help:"bprint dev stuff."`
	BprintFile string
}

func main() {
	cli := &CLI{}
	ctx := kong.Parse(cli)

	pp.Println(ctx.Run(&bdev.UpperScope{
		BprintFile: cli.BprintFile,
		Ctx:        ctx,
	}))

}
