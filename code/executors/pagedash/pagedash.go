package pagedash

import (
	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
)

type PageDash struct {
	builder   *PdBuilder
	jsruntime *goja.Runtime
	binder    bindx.Bindings
}

func (pd *PageDash) Process(ev *event.Request) (*event.Response, error) {

	return nil, nil
}
