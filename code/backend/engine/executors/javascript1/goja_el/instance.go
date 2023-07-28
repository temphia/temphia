package evgoja

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

type GojaInstance struct {
	tenantId string
	evLoop   *eventloop.EventLoop
	rt       *goja.Runtime

	lastBinding  bindx.Core
	activeEvents map[string]bindx.Bindings
	needsClose   bool
}

func NewInstance(tenantId string) *GojaInstance {
	registry := require.NewRegistry()

	gi := &GojaInstance{
		tenantId: tenantId,
		evLoop:   nil,
		rt:       nil,

		activeEvents: make(map[string]bindx.Bindings),
		lastBinding:  nil,
		needsClose:   false,
	}

	evl := eventloop.NewEventLoop()

	evl.Run(func(r *goja.Runtime) {
		gi.rt = r
		gi.evLoop = evl

		registry.Enable(r)
		registry.RegisterNativeModule("temphia", gi.temphiaBindings)
	})

	return gi
}

func (g *GojaInstance) Init(b bindx.Bindings) error {
	g.lastBinding = b

	g.evLoop.Start()

	return nil
}
