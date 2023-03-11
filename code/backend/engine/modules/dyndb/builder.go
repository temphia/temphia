package dyndb

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

var _ etypes.ModuleBuilder = (*DyndbBuilder)(nil)

func NewBuilder(app any) (etypes.ModuleBuilder, error) {
	return &DyndbBuilder{}, nil
}

type DyndbBuilder struct{}

func (DyndbBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {

	deps := opts.Binder.GetApp().(xtypes.App).GetDeps()

	target, err := opts.Resource.SplitTarget(2)
	if err != nil {
		return nil, err
	}

	dynhub := deps.DataHub().(dyndb.DataHub).GetSource(target[0], opts.Resource.TenantId)

	return &DyndbModule{
		binder:   opts.Binder,
		res:      opts.Resource,
		dynsrc:   dynhub,
		group:    target[1],
		tenantId: opts.Resource.Target,
	}, nil

}
