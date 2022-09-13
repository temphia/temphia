package domain

import (
	"github.com/temphia/temphia/code/core/backend/engine/invokers"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/invoker"
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
	return nil, invokers.ErrInvokerActionNotImplemented
}

func (e *Editor) Name() string {
	return invokers.TypeDomainEditor
}

func (e *Editor) CurrentUser() *invoker.User {
	return nil
}
