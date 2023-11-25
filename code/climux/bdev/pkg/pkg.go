package pkg

import (
	"fmt"

	"github.com/temphia/temphia/code/climux/bdev/core"
	"github.com/temphia/temphia/code/tools/repobuild/builder"
)

type PkgCLI struct {
	Zip struct {
		OutFile string
	} `cmd:"" help:"Zip and package bashed on brpint.toml."`
}

func (c *PkgCLI) Run(ctx core.BdevContext) error {

	mf, err := ctx.ReadConfig()
	if err != nil {
		return err
	}

	if c.Zip.OutFile == "" {
		c.Zip.OutFile = fmt.Sprintf("build/%s.zip", mf.Slug)
	}

	return builder.ZipIt(mf, c.Zip.OutFile)
}
