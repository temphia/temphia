package demo

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/temphia/temphia/code/backend/libx/xutils"
)

type CLI struct {
	Run struct {
	} `cmd:"" help:"Run server."`

	ClearLock struct {
	} `cmd:"" help:"Clear Postgres Lock"`

	Reset struct {
	} `cmd:"" help:"Reset server data"`

	ctx *kong.Context
}

func NewCLI() *CLI {
	cli := &CLI{}
	ctx := kong.Parse(cli)
	cli.ctx = ctx

	return cli
}

func (c *CLI) Execute() error {

	switch c.ctx.Command() {
	case "run":

		os.Chdir("cmd/demo/")
		xutils.CreateIfNotExits("temphia-data/files")
		xutils.CreateIfNotExits("temphia-data/logs")
		xutils.CreateIfNotExits("temphia-data/pgdata")

		if xutils.FileExists("temphia-data/pgdata/data", "postmaster.pid") {
			fmt.Println("looks like another demo instance is running or last one did not close you might have to clear lock with 'temphia-demo clear-lock'")
		}

		err := RunDemo()

		if xutils.FileExists("temphia-data/pgdata/data", "postmaster.pid") {
			fmt.Println("Looks like db did not close properly, might need to clear lock 'temphia-demo clear-lock'")
		}

		return err

	case "clear-lock":
		return ClearLock()
	case "reset":
		return Reset()
	default:
		panic("Command not found" + c.ctx.Command())
	}
}
