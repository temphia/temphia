package logproxy

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type LogProxy struct {
	app      xtypes.App
	resource *entities.Resource
}

func (l *LogProxy) IPC(method string, path string, args xtypes.LazyData) (xtypes.LazyData, error) {
	return nil, easyerr.NotImpl()
}

func (l *LogProxy) Close() error {
	return nil
}
