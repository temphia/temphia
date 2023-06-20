package evgoja

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
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

	evl := eventloop.NewEventLoop()
	var rt *goja.Runtime

	evl.Run(func(r *goja.Runtime) {
		rt = r
	})

	return &GojaInstance{
		tenantId: tenantId,
		evLoop:   evl,
		rt:       rt,

		activeEvents: make(map[string]bindx.Bindings),
		lastBinding:  nil,
		needsClose:   false,
	}
}

func (g *GojaInstance) Init(b bindx.Bindings) error {
	g.lastBinding = b
	g.attachBindings()
	g.evLoop.Start()

	return nil
}

func (g *GojaInstance) attachBindings() {

}
