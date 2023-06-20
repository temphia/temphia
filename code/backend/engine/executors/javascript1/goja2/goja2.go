package evgoja

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
)

var _ etypes.Executor = (*Goja)(nil)

type Goja struct {
	instance *GojaInstance
	binding  bindx.Bindings
}

func NewExecutor(instance *GojaInstance) func(opts etypes.ExecutorOption) (etypes.Executor, error) {
	return func(opts etypes.ExecutorOption) (etypes.Executor, error) {
		return &Goja{
			instance: instance,
			binding:  opts.Binder,
		}, nil
	}
}

type Response struct {
	Payload any `json:"payload,omitempty"`
}

type Request struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Data any    `json:"data,omitempty"`
}

func (g *Goja) Process(ev *event.Request) (*event.Response, error) {

	return nil, nil
}
