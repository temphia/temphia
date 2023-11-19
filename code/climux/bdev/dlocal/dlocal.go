package dlocal

import "github.com/temphia/temphia/code/climux/bdev/core"

/*


	fixme impl => xinstancer.go ( xinstancer.MigTypeAddColumn)



*/

type DLocalCLI struct {
	Init struct{} `cmd:"" help:"init dydb migration with step to run in app.json"`
}

func (c *DLocalCLI) Run(ctx core.BdevContext) error { return nil }
