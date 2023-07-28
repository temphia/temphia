package websocket

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
)

type SockdMod struct {
	sockd    sockdx.SockdCore
	tenantId string
}

func (s *SockdMod) Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error) {

	// use modipc here

	return nil, nil
}

func (s *SockdMod) Close() error { return nil }

func (s *SockdMod) SendDirect(room string, connId int64, payload []byte) error {
	return s.sockd.SendDirect(s.tenantId, room, int64(connId), payload)
}

func (s *SockdMod) SendDirectBatch(room string, conns []int64, payload []byte) error {
	return s.sockd.SendDirectBatch(s.tenantId, room, conns, payload)
}

func (s *SockdMod) SendBroadcast(room string, ignores []int64, payload []byte) error {
	return s.sockd.SendBroadcast(s.tenantId, room, ignores, payload)
}

func (s *SockdMod) SendTagged(room string, tags []string, ignores []int64, payload []byte) error {
	return s.sockd.SendTagged(s.tenantId, room, tags, ignores, payload)
}

func (s *SockdMod) RoomUpdateTags(room string, opts sockdx.UpdateTagOptions) error {
	return s.sockd.RoomUpdateTags(s.tenantId, room, opts)
}

var _ etypes.ModuleBuilder = (*SockdModBuilder)(nil)

func NewBuilder(app any) (etypes.ModuleBuilder, error) {
	return &SockdModBuilder{}, nil
}

type SockdModBuilder struct{}

func (p *SockdModBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	return New(
		opts.Resource.TenantId,
		opts.Binder.GetApp().(xtypes.App),
	), nil
}

func New(tenantId string, app xtypes.App) *SockdMod {

	bm := &SockdMod{
		sockd: app.GetDeps().SockdHub().(sockdx.Sockd),
	}

	return bm

}
