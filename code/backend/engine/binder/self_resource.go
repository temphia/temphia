package binder

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities/resource"
)

func (b *SelfBindings) SelfListResources() ([]*bindx.Resource, error) {
	b.handle.LoadResources()

	ress := make([]*bindx.Resource, 0, len(b.handle.Resources))
	for _, r := range b.handle.Resources {
		ress = append(ress, &bindx.Resource{
			Name:    r.Name,
			Type:    r.Type,
			Payload: "",
			Meta:    nil,
		})
	}

	return ress, nil
}

func (b *SelfBindings) SelfGetResource(name string) (*bindx.Resource, error) {
	b.handle.LoadResources()

	res, ok := b.handle.Resources[name]
	if !ok {
		return nil, easyerr.Error(etypes.ResourceNotFound)
	}

	return &bindx.Resource{
		Name:    res.Name,
		Type:    res.Type,
		Payload: "",
		Meta:    nil,
	}, nil
}

// module

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

// private

func (b *SelfBindings) execModule(mtype, method, path string, data xtypes.LazyData, res *entities.Resource) (xtypes.LazyData, error) {

	mod, err := b.getModule(mtype, res)
	if err != nil {
		return nil, err
	}

	return mod.IPC(method, path, data)
}

func (b *SelfBindings) getModule(mtype string, res *entities.Resource) (etypes.Module, error) {
	mod := b.activeModules[res.Id]
	if mod != nil {
		return mod, nil
	}

	mbuilder, ok := b.handle.Deps.ModuleBuilders[mtype]
	if !ok {
		return nil, easyerr.Error(etypes.ResourceModuleNotFound)
	}

	modInstance, err := mbuilder.Instance(etypes.ModuleOptions{
		Binder:   b.root,
		Resource: res,
	})
	if err != nil {
		return nil, err
	}

	b.activeModules[res.Id] = modInstance
	return modInstance, nil
}
