package binder

import (
	"fmt"

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

	// uctx := pkv.getUserCtx()
	// if uctx == nil {
	// 	return "", easyerr.Error(etypes.EmptyUserContext)
	// }

	// return pkv.signer.SignPlugState(pkv.namespace, &claim.PlugState{
	// 	TenantId:  pkv.namespace,
	// 	Type:      "",
	// 	UserId:    uctx.UserID,
	// 	DeviceId:  uctx.DeviceId,
	// 	SessionId: uctx.SessionID,
	// 	ExecId:    0,
	// 	PlugId:    pkv.plugId,
	// 	AgentId:   pkv.agentid,
	// 	KeyPrefix: opts.KeyPrefix,
	// })

	// Ticket(room string, opts *ticket.SockdRoom) (string, error)

	/*





		uctx := s.handle.Job.Invoker.UserContext()
		if uctx == nil {
			return "", easyerr.Error(etypes.EmptyUserContext)
		}

		s.handle.LoadResources()

		res := s.handle.Resources[room]
		if res == nil {
			return "", easyerr.NotFound("Resource room")
		}

		return s.handle.Deps.Signer.SignSockdTkt(s.tenantId, &claim.SockdTkt{
			UserId:    uctx.UserID,
			Room:      res.Id,
			DeviceId:  uctx.DeviceId,
			SessionId: uctx.SessionID,
		})

	*/

	return "", nil
}

func (b *SelfBindings) selfModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error) {

	fmt.Println("@", b.activeModules, mid)

	mod, ok := b.activeModules[mid]
	if !ok {
		return nil, easyerr.NotFound("resource module 1")
	}

	return mod.Handle(method, data)
}
