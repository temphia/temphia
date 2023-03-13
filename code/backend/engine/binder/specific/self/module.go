package self

import (
	"errors"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities/resource"
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

	switch res.Type {
	case resource.Module:
		return b.execModule(method, path, data, res)
	case resource.DataGroup:
		return nil, easyerr.NotImpl()
	case resource.Folder:
		return nil, easyerr.NotImpl()
	default:
		panic("Not impl")
	}
}

func (b *Binding) execModule(method, path string, data xtypes.LazyData, res *entities.Resource) (xtypes.LazyData, error) {

	mbuilder, ok := b.handle.Deps.ModuleBuilders[res.SubType]
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
