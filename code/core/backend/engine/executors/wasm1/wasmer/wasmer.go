package wasmer

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/event"
	"github.com/wasmerio/wasmer-go/wasmer"
)

type Executor struct {
	builder *Builder
	module  *wasmer.Module

	// bindings cache
	bindings   bindx.Bindings
	bindPluKV  bindx.PlugKV
	bindSockd  bindx.Sockd
	bindUser   bindx.User
	bindCab    bindx.Cabinet
	bindSelf   bindx.Self
	bindNcache bindx.NodeCache
	bindNet    bindx.Net
}

func (e *Executor) Process(req *event.Request) (*event.Response, error) {

	return nil, nil
}
