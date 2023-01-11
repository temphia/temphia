package goja

import (
	"github.com/dop251/goja"
	sdkdist "github.com/temphia/temphia/code/core/backend/engine/executors/javascript1/sdk/dist"
)

var libesplug *goja.Program

func init() {
	program, err := goja.Compile("libesplug", sdkdist.LibESPlug, true)
	if err != nil {
		panic(err)
	}
	libesplug = program
}
