package demo

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/distro/sharedcli"
)

type CLI struct {
	StartPg struct {
	} `cmd:"" help:"Start demo server."`

	ClearLockPg struct {
	} `cmd:"" help:"Clear Postgres Lock"`

	ResetPg struct {
	} `cmd:"" help:"Reset server data"`

	Start struct {
	} `cmd:"" help:"Start demo server."`

	ctx *kong.Context
}

func NewCLI() *CLI {
	cli := &CLI{}
	ctx := kong.Parse(cli)
	cli.ctx = ctx

	return cli
}

func (c *CLI) Run(scope *sharedcli.Context) error {
	c.ctx = scope.KongContext
	return c.doExecute("demo ")
}

func (c *CLI) Execute() error {
	return c.doExecute("")
}

func (c *CLI) doExecute(prefix string) error {

	switch c.ctx.Command() {

	case prefix + "start":
		initData()
		Conf = sqliteConf
		return RunDemo()
	case prefix + "start-pg":

		initData()

		Conf = postgresConf

		if xutils.FileExists("temphia-data/pgdata/data", "postmaster.pid") {
			fmt.Println("looks like another demo instance is running or last one did not close you might have to clear lock with 'temphia-demo clear-lock'")
		}

		err := RunDemo()

		if xutils.FileExists("temphia-data/pgdata/data", "postmaster.pid") {
			fmt.Println("Looks like db did not close properly, might need to clear lock 'temphia-demo clear-lock'")
		}

		return err

	case prefix + "clear-lock-pg":
		return ClearLock()
	case prefix + "reset-pg":
		return Reset()
	default:
		panic("Command not found" + c.ctx.Command())
	}
}

func initData() {
	os.Chdir("cmd/demo/")

	xutils.CreateIfNotExits("temphia-data")
	xutils.CreateIfNotExits("temphia-data/files")
	xutils.CreateIfNotExits("temphia-data/logs")
	xutils.CreateIfNotExits("temphia-data/pgdata")
	xutils.CreateIfNotExits("temphia-data/sqlitedata")

}
