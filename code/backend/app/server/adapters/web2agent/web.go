package web2agent

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
)

func (w *WATarget) executeModule(module, action string, data xtypes.LazyData) (xtypes.LazyData, error) {

	switch module {
	case "web":
		return w.webExecuteModule(action, data)
	default:
		return nil, easyerr.NotImpl()
	}
}

func (w *WATarget) webExecuteModule(action string, data xtypes.LazyData) (xtypes.LazyData, error) {

	return nil, nil
}
