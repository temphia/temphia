package logproxy

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

var _ etypes.ModuleBuilder = (*LogProxyBuilder)(nil)

func NewBuilder(app any) (etypes.ModuleBuilder, error) {
	return &LogProxyBuilder{
		app: app.(xtypes.App),
	}, nil
}

type LogProxyBuilder struct {
	app xtypes.App
}

func (l LogProxyBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	return &LogProxy{
		app:      l.app,
		resource: opts.Resource,
	}, nil
}
