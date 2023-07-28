package etypes

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type ModuleOptions struct {
	Binder       ExecutorBinder
	Resource     *entities.Resource
	InvokerToken string
	Args         xtypes.LazyData
}

type ModuleBuilderFunc func(app any) (ModuleBuilder, error)

type ModuleBuilder interface {
	Instance(ModuleOptions) (Module, error)
}

type Module interface {
	Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error)
	Close() error
}
