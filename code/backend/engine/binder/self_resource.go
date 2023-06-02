package binder

import (
	"fmt"
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
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

	counter := int32(0)

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

		counter = b.activeModCounter + 1
		b.activeModCounter = counter
		b.activeModules[counter] = mod

	case resource.Folder:
		fallthrough
	default:
		panic("Not impl")
	}

	return counter, nil
}

type PlugState struct {
	KeyPrefix string `json:"key_prefix,omitempty"`
}

type CabinetFolder struct {
	Prefix      string   `json:"prefix,omitempty"`
	PinnedFiles []string `json:"pinned_files,omitempty"`
	Operations  []string `json:"ops,omitempty"`
}

type SockdRoom struct {
	AllowBroadcast string `json:"allow_broadcast,omitempty"`
}

type DataGroup struct {
	ReadOnly bool `json:"read_only,omitempty"`
}

func (b *SelfBindings) SelfModuleTicket(name string, opts xtypes.LazyData) (string, error) {

	signer := b.handle.Deps.Signer
	uctx := b.root.invoker.ContextUser()

	switch name {
	case "self_plugstate":

		popts := PlugState{}
		err := opts.AsObject(popts)
		if err != nil {
			return "", err
		}

		return signer.SignPlugState(b.handle.Namespace, &claim.PlugState{
			TenantId:  b.handle.Namespace,
			Type:      "",
			UserId:    uctx.UserID,
			DeviceId:  uctx.DeviceId,
			SessionId: uctx.SessionID,
			ExecId:    0,
			PlugId:    b.handle.PlugId,
			AgentId:   b.handle.AgentId,
			KeyPrefix: popts.KeyPrefix,
		})
	case "self_bprint":
		return "", easyerr.NotImpl()
	}

	b.handle.LoadResources()

	res, ok := b.handle.Resources[name]
	if !ok {
		return "", easyerr.Error(etypes.ResourceNotFound)
	}

	switch res.Type {
	case resource.DataGroup:
		target := strings.Split(res.Target, "/")

		return signer.SignData(b.handle.Namespace, &claim.Data{
			TenantId:   b.handle.Namespace,
			Type:       "",
			UserID:     uctx.UserID,
			UserGroup:  uctx.UserGroup,
			SessionID:  uctx.SessionID,
			DeviceId:   uctx.DeviceId,
			DataSource: target[0],
			DataGroup:  target[1],
			DataTables: []string{"*"},
			IsExec:     true,
		})

	default:
		return "", easyerr.NotImpl()

	}

}

func (b *SelfBindings) selfModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error) {

	fmt.Println("@", b.activeModules, mid)

	mod, ok := b.activeModules[mid]
	if !ok {
		return nil, easyerr.NotFound("resource module 1")
	}

	return mod.Handle(method, data)
}
