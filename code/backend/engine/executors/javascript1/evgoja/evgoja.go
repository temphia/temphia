package evgoja

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

type EventLoopGoja struct {
	evl *eventloop.EventLoop
	rt  *goja.Runtime
}

func New() *EventLoopGoja {
	evl := eventloop.NewEventLoop()

	var rt *goja.Runtime

	evl.Run(func(r *goja.Runtime) {
		rt = r
	})

	return &EventLoopGoja{
		evl: evl,
		rt:  rt,
	}
}

func (e *EventLoopGoja) init(b bindx.Bindings) {

	e.rt.Set("NativeModule", NewNativeModule(b))

}

func NewNativeModule(b bindx.Bindings) func(call goja.ConstructorCall) *goja.Object {

	return func(call goja.ConstructorCall) *goja.Object {

		return nil
	}

}
