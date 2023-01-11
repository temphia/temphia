package sockd

import (
	"github.com/temphia/temphia/code/backend/engine/binders/standard/handle"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
)

type Binding struct {
	sockd    sockdx.SockdCore
	tenantId string
}

func New(handle *handle.Handle) Binding {
	return Binding{
		sockd:    handle.Deps.Sockd,
		tenantId: handle.Namespace,
	}
}

func (s *Binding) SendDirect(room string, connId int64, payload []byte) error {
	return s.sockd.SendDirect(s.tenantId, room, int64(connId), payload)
}

func (s *Binding) SendDirectBatch(room string, conns []int64, payload []byte) error {
	return s.sockd.SendDirectBatch(s.tenantId, room, conns, payload)
}

func (s *Binding) SendBroadcast(room string, ignores []int64, payload []byte) error {
	return s.sockd.SendBroadcast(s.tenantId, room, ignores, payload)
}

func (s *Binding) SendTagged(room string, tags []string, ignores []int64, payload []byte) error {
	return s.sockd.SendTagged(s.tenantId, room, tags, ignores, payload)
}

func (s *Binding) RoomUpdateTags(room string, opts sockdx.UpdateTagOptions) error {
	return s.sockd.RoomUpdateTags(s.tenantId, room, opts)
}
