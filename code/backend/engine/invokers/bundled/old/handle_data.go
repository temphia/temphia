package web

import (
	"github.com/temphia/temphia/code/backend/engine/invokers"
	"github.com/temphia/temphia/code/backend/xtypes"
)

func (r *WebRequest) dataHandle(method string, data xtypes.LazyData) (xtypes.LazyData, error) {

	switch method {
	case "data.simple_query":
		return nil, nil
	default:
		return nil, invokers.ErrInvokerActionNotImplemented
	}

}
