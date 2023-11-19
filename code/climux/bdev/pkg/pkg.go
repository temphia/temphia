package pkg

import "github.com/temphia/temphia/code/climux/bdev/core"

type PkgCLI struct {
	Zip struct{} `cmd:"" help:"Zip and package bashed on brpint.yaml."`
}

func (c *PkgCLI) Run(ctx core.BdevContext) error { return nil }
