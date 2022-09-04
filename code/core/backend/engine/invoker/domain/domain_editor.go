package domain

import (
	"github.com/temphia/temphia/code/core/backend/engine/invoker"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

type EditorOptions struct {
	Domain *entities.TenantDomain
}

type Editor struct {
	domain *entities.TenantDomain
}

func NewEditor(opts EditorOptions) *Editor {
	return &Editor{
		domain: opts.Domain,
	}
}

func (e *Editor) Handle(method string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return nil, invoker.ErrInvokerActionNotImplemented
}

func (e *Editor) Name() string {
	return invoker.TypeDomainEditor
}

func (e *Editor) CurrentUser() *job.InvokeUser {
	return nil
}
