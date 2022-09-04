package forked

import (
	"github.com/temphia/temphia/code/core/backend/engine/invoker"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
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
	return nil, invoker.ErrInvokerActionNotImplemented
}

func (f *Forked) Name() string {
	return invoker.TypeForked
}

func (f *Forked) CurrentUser() *job.InvokeUser {
	return nil
}
