package web2agent

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
)

func (w *WATarget) executeMethod(method string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return nil, easyerr.NotImpl()
}
