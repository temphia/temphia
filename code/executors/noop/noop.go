package noop

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
)

type Noop struct {
	builder *NoopBuilder
	opts    etypes.ExecutorOption
}

func (n *Noop) Process(ev *event.Request) (*event.Response, error) {

	return &event.Response{
		Payload: []byte(`{}`),
	}, nil

}
