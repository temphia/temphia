package binder

import (
	"github.com/temphia/temphia/code/backend/engine/binder/handle"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
)

type SockdBinding struct {
	sockd    sockdx.SockdCore
	tenantId string
	handle   *handle.Handle
}

func New(handle *handle.Handle) SockdBinding {
	return SockdBinding{
		sockd:    handle.Deps.Sockd,
		tenantId: handle.Namespace,
		handle:   handle,
	}
}

func (s *SockdBinding) SendDirect(room string, connId int64, payload []byte) error {
	return s.sockd.SendDirect(s.tenantId, room, int64(connId), payload)
}

func (s *SockdBinding) SendDirectBatch(room string, conns []int64, payload []byte) error {
	return s.sockd.SendDirectBatch(s.tenantId, room, conns, payload)
}

func (s *SockdBinding) SendBroadcast(room string, ignores []int64, payload []byte) error {
	return s.sockd.SendBroadcast(s.tenantId, room, ignores, payload)
}

func (s *SockdBinding) SendTagged(room string, tags []string, ignores []int64, payload []byte) error {
	return s.sockd.SendTagged(s.tenantId, room, tags, ignores, payload)
}

func (s *SockdBinding) RoomUpdateTags(room string, opts sockdx.UpdateTagOptions) error {
	return s.sockd.RoomUpdateTags(s.tenantId, room, opts)
}
