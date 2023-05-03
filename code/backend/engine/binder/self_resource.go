package binder

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
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

func (b *SelfBindings) selfNewModule(name string, args xtypes.LazyData) (int32, error) {
	b.handle.LoadResources()

	res, ok := b.handle.Resources[name]
	if !ok {
		return 0, easyerr.Error(etypes.ResourceNotFound)
	}

	modname := res.SubType

	switch res.Type {

	case resource.DataGroup:
		modname = resource.DataGroup
		fallthrough
	case resource.Module:
		modbuilder, ok := b.handle.Deps.ModuleBuilders[modname]
		if !ok {
			return 0, easyerr.NotFound("resource module")
		}

		mod, err := modbuilder.Instance(etypes.ModuleOptions{
			Binder:       b.root,
			Resource:     res,
			InvokerToken: "",
			Args:         args,
		})

		if err != nil {
			return 0, err
		}

		b.activeModCounter = b.activeModCounter + 1
		b.activeModules[b.activeModCounter] = mod

	case resource.Folder:
		fallthrough
	default:
		panic("Not impl")
	}

	return 0, nil
}

func (b *SelfBindings) selfModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error) {

	mod, ok := b.activeModules[mid]
	if !ok {
		return nil, easyerr.NotFound("resource module 1")
	}

	return mod.Handle(method, data)
}
