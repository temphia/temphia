package pageform

import "github.com/temphia/temphia/code/backend/xtypes/etypes/event"

type Pageform struct {
	builder *PfBuilder
}

func (pf *Pageform) Process(*event.Request) (*event.Response, error) {
	return nil, nil
}
