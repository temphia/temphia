package bdev

import (
	"github.com/temphia/temphia/code/climux/bdev/agent"
	"github.com/temphia/temphia/code/climux/bdev/resource/cfolder"
	"github.com/temphia/temphia/code/climux/bdev/resource/dyndb"
	"github.com/temphia/temphia/code/climux/bdev/resource/socket"

	"github.com/temphia/temphia/code/climux/bdev/dlocal"

	"github.com/temphia/temphia/code/climux/bdev/pkg"
	"github.com/temphia/temphia/code/climux/bdev/resource"
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
