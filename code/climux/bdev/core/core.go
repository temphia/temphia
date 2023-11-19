package core

import (
	"github.com/alecthomas/kong"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
)

type BdevContext struct {
	ConfigFile string
	KongCtx    *kong.Context
}

func (b *BdevContext) MustGetConfig() *xpackage.Manifest {

	return nil
}
