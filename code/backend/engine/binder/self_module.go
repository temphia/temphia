package binder

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities/resource"
)

func (b *SelfBindings) selfModuleExec(name, method, path string, data xtypes.LazyData) (xtypes.LazyData, error) {
	b.handle.LoadResources()

	res, ok := b.handle.Resources[name]
	if !ok {
		return nil, easyerr.Error(etypes.ResourceNotFound)
	}

	switch res.Type {
	case resource.Module:
		return b.execModule(res.SubType, method, path, data, res)
	case resource.DataGroup:
		return b.execModule(resource.DataGroup, method, path, data, res)
	case resource.Folder:
		return nil, easyerr.NotImpl()
	default:
		panic("Not impl")
	}
}

func (b *SelfBindings) execModule(name, method, path string, data xtypes.LazyData, res *entities.Resource) (xtypes.LazyData, error) {

	mbuilder, ok := b.handle.Deps.ModuleBuilders[name]
	if !ok {
		return nil, easyerr.Error(etypes.ResourceModuleNotFound)
	}

	module, err := mbuilder.Instance(etypes.ModuleOptions{
		Binder:   b.root,
		Resource: res,
	})
	if err != nil {
		return nil, err
	}

	return module.IPC(method, path, data)
}
