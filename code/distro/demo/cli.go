package demo

import "github.com/alecthomas/kong"

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
		return RunDemo()
	case "clear-lock":
		return ClearLock()
	case "reset":
		return Reset()
	default:
		panic("Command not found" + c.ctx.Command())
	}
}
