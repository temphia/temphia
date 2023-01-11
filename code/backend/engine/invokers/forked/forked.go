package forked

import (
	"github.com/temphia/temphia/code/backend/engine/invokers"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
)

type Forked struct {
	forkedFrom string
}

func New(from string) *Forked {
	return &Forked{
		forkedFrom: from,
	}
}

func (f *Forked) Handle(method string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return nil, invokers.ErrInvokerActionNotImplemented
}

func (f *Forked) Name() string {
	return invokers.TypeForked
}

func (f *Forked) User() *invoker.User {
	return nil
}
