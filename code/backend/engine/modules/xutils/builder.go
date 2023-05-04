package xutils

import (
	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

var _ etypes.ModuleBuilder = (*XBuilder)(nil)

func NewBuilder(app any) (etypes.ModuleBuilder, error) {
	return &XBuilder{
		app: app.(xtypes.App),
	}, nil
}

type XBuilder struct {
	app xtypes.App
}

func (l XBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {

	xutils := &Xutils{}
	xutils.modipc = modipc.NewModIPC(xutils)

	return xutils, nil
}
