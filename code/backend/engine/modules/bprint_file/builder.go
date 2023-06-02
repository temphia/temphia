package bprint

import (
	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
)

var _ etypes.ModuleBuilder = (*BprintModBuilder)(nil)

func NewBuilder(app any) (etypes.ModuleBuilder, error) {
	return &BprintModBuilder{}, nil
}

type BprintModBuilder struct{}

func (p *BprintModBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	return New(
		opts.Resource.TenantId,
		opts.Resource.Target,
		opts.Binder.GetApp().(xtypes.App),
	), nil
}

func New(tenantId, bid string, app xtypes.App) *BprintMod {

	bm := &BprintMod{
		tenantId: tenantId,
		bid:      bid,
		modipc:   nil,
		bhub:     app.GetDeps().RepoHub().(repox.Hub),
	}

	bm.modipc = modipc.NewModIPC(bm)

	return bm

}
