package logproxy

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type LogProxy struct {
	app      xtypes.App
	resource *entities.Resource
}

func (l *LogProxy) IPC(method string, path string, args xtypes.LazyData) (xtypes.LazyData, error) {

	ls := l.app.GetDeps().LogService().(logx.Service)
	lproxy := ls.GetLogProxy()

	lproxy.Query("", "", "", "", map[string]string{})

	return nil, easyerr.NotImpl()
}

func (l *LogProxy) Close() error {
	return nil
}
