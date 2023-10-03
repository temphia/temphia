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

func (b *Binder) ListResources() ([]*bindx.Resource, error) {
	b.loadResources()

	ress := make([]*bindx.Resource, 0, len(b.resources))
	for _, r := range b.resources {
		ress = append(ress, &bindx.Resource{
			Name:    r.Name,
			Type:    r.Type,
			Payload: "",
			Meta:    nil,
		})
	}

	return ress, nil
}

func (b *Binder) GetResource(name string) (*bindx.Resource, error) {
	b.loadResources()

	res, ok := b.resources[name]
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

func (b *Binder) selfNewModule(name string, args xtypes.LazyData) (int32, error) {
	b.loadResources()

	res, ok := b.resources[name]
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
		modbuilder, ok := b.Deps.ModuleBuilders[modname]
		if !ok {
			return 0, easyerr.NotFound("resource module")
		}

		mod, err := modbuilder.New(etypes.ModuleOptions{
			Binder:       nil,
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

func (b *Binder) moduleTicket(name string, opts xtypes.LazyData) (string, error) {

	signer := b.Deps.Signer

	switch name {
	case "self_plugstate":

		popts := &PlugState{}
		err := opts.AsObject(popts)
		if err != nil {
			return "", err
		}

		return signer.SignPlugState(b.Namespace, &claim.PlugState{
			TenantId: b.Namespace,
			Type:     "",
			// fixme =>
			// UserId:    uctx.UserID,
			// DeviceId:  uctx.DeviceId,
			// SessionId: uctx.SessionID,
			ExecId:    0,
			PlugId:    b.PlugId,
			AgentId:   b.AgentId,
			KeyPrefix: popts.KeyPrefix,
		})
	case "self_bprint":
		return "", easyerr.NotImpl()
	}

	b.loadResources()

	res, ok := b.resources[name]
	if !ok {
		return "", easyerr.Error(etypes.ResourceNotFound)
	}

	switch res.Type {
	case resource.DataGroup:
		target := strings.Split(res.Target, "/")

		return signer.SignData(b.Namespace, &claim.Data{
			TenantId: b.Namespace,
			// UserID:     uctx.UserID,
			// UserGroup:  uctx.UserGroup,
			// SessionID:  uctx.SessionID,
			// DeviceId:   uctx.DeviceId,
			DataSource: target[0],
			DataGroup:  target[1],
			DataTables: []string{"*"},
			IsExec:     true,
		})
	case resource.SockRoom:

		return signer.SignSockdTkt(b.Namespace, &claim.SockdTkt{
			// UserId:    uctx.UserID,
			// DeviceId:  uctx.DeviceId,
			// SessionId: uctx.SessionID,
			Room: res.Id,
		})

	default:
		return "", easyerr.NotImpl()

	}

}

func (b *Binder) selfModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error) {

	fmt.Println("@", b.activeModules, mid)

	mod, ok := b.activeModules[mid]
	if !ok {
		return nil, easyerr.NotFound("resource module 1")
	}

	return mod.Handle(method, data)
}
