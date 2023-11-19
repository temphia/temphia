package bdev

import (
	"os"

	"github.com/temphia/temphia/code/climux/bdev/agent"
	"github.com/temphia/temphia/code/climux/bdev/cfolder"
	"github.com/temphia/temphia/code/climux/bdev/dlocal"
	"github.com/temphia/temphia/code/climux/bdev/dyndb"
	"github.com/temphia/temphia/code/climux/bdev/pkg"
	"github.com/temphia/temphia/code/climux/bdev/resource"
	"github.com/temphia/temphia/code/climux/bdev/socket"
)

type BdevCLI struct {
	Agent    *agent.AgentCLI       `cmd:"" help:"Agent related actions."`
	CFolder  *cfolder.CFolderCLI   `cmd:"" help:"Cabinet folder related actions."`
	DLocal   *dlocal.DLocalCLI     `cmd:"" help:"Local Dyndb related actions."`
	Dyndb    *dyndb.DyndbCLI       `cmd:"" help:"Dyndb related actions."`
	Pkg      *pkg.PkgCLI           `cmd:"" help:"Package relaed actions."`
	Resource *resource.ResourceCLI `cmd:"" help:"Resource actions."`
	Socket   *socket.SocketCLI     `cmd:"" help:"Socket actions."`

	BprintFile string
}

func (c *BdevCLI) Run() error {

	bconf := os.Getenv("TEMPHIA_BDEV_BPRINT_CONFIG")
	if bconf == "" {
		panic(".bprint.yaml not specified")
	}

	return nil
}
