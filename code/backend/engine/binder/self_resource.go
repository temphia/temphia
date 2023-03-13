package binder

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
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
