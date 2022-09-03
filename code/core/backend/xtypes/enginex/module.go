package enginex

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

type ModuleBuilderFunc func(app interface{}) (ModuleBuilder, error)

type ModuleBuilder interface {
	Instance(ModuleOptions) (Module, error)
}

type ModuleOptions struct {
	Binder   ExecutorBinder
	Resource *entities.Resource
}

type Module interface {
	IPC(method string, path string, args xtypes.LazyData) (xtypes.LazyData, error)
	Close() error
}
