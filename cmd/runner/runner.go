package runner

import "github.com/alecthomas/kong"

type CLI struct {
	Start struct {
		Config string `help:"config file"`
	} `cmd:"" help:"Start server."`

	ActualStart struct {
		Config string `help:"config file"`
	} `cmd:"" help:"Start server but called by runner do not call directly."`

	ctx *kong.Context
}
