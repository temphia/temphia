package websocket

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

var _ etypes.ModuleBuilder = (*SocketTargetBuilder)(nil)

func NewSocketTargetBuilder(app any) (etypes.ModuleBuilder, error) {
	return &SocketTargetBuilder{}, nil
}

type SocketTargetBuilder struct{}

func (p *SocketTargetBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	return NewSocketTarget(
		opts.Resource.TenantId,
		opts.Binder.GetApp().(xtypes.App),
	), nil
}

func NewSocketTarget(tenantId string, app xtypes.App) *SocketTargetBuilder {

	bm := &SocketTargetBuilder{}

	return bm

}

func (s *SocketTargetBuilder) Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error) {
	return nil, nil
}

func (s *SocketTargetBuilder) Close() error { return nil }
