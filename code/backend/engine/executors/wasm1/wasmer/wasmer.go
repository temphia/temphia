package wasmer

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/wasmerio/wasmer-go/wasmer"
)

type Executor struct {
	builder  *Builder
	module   *wasmer.Module
	instance *wasmer.Instance
	extenFns map[string]wasmer.IntoExtern

	// bindings cache
	bindings  bindx.Bindings
	bindPluKV bindx.PlugKV
	bindSockd bindx.Sockd
	bindCab   bindx.Cabinet
	bindSelf  bindx.Self
	bindNet   bindx.Net
}

func (e *Executor) Process(req *event.Request) (*event.Response, error) {

	return nil, nil
}
