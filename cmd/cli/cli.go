package cli

import (
	"github.com/alecthomas/kong"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/distro/demo"
	"github.com/temphia/temphia/code/distro/sharedcli"
	"github.com/temphia/temphia/code/tools/bdev"

	_ "github.com/temphia/temphia/code/distro/common"
)

type CLI struct {
	Demo       demo.CLI `cmd:"" help:"demo cli stuff."`
	Bdev       bdev.CLI `cmd:"" help:"bprint dev stuff."`
	BprintFile string
	ConfigFile string
}

func RunMain() {
	cli := &CLI{}
	ctx := kong.Parse(cli)

	pp.Println(ctx.Run(&sharedcli.Context{
		BprintFile:  cli.BprintFile,
		ConfigFile:  cli.ConfigFile,
		KongContext: ctx,
	}))

}
