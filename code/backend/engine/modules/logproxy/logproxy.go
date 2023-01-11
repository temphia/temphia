package logproxy

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type LogProxy struct {
	app xtypes.App
}

func (l *LogProxy) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	return nil, nil
}

func (l *LogProxy) Init(app any) error {

	l.app = app.(xtypes.App)

	return nil
}
