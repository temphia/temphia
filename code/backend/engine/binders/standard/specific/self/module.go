package self

import (
	"errors"

	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities/resource"
)

var (
	ErrModuleNotFound          = errors.New("MODULE NOT FOUND")
	ErrModuleResourceWrongType = errors.New("MODULE RESOURCE WRONG TYPE")
	ErrModuleResourceNotFound  = errors.New("MODULE RESOURCE NOT FOUND")
)

func (b *Binding) selfModuleExec(name, method, path string, data xtypes.LazyData) (xtypes.LazyData, error) {
	b.handle.LoadResources()

	res, ok := b.handle.Resources[name]
	if !ok {
		return nil, ErrModuleResourceNotFound
	}

	if resource.Module != res.Type {
		return nil, ErrModuleResourceWrongType
	}

	mbuilder, ok := b.handle.Deps.ModuleBuilders[res.Target]
	if !ok {
		return nil, ErrModuleNotFound
	}

	module, err := mbuilder.Instance(etypes.ModuleOptions{
		Binder:   nil,
		Resource: res,
	})
	if err != nil {
		return nil, err
	}

	return module.IPC(method, path, data)
}
