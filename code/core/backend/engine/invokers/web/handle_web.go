package web

import (
	"github.com/temphia/temphia/code/core/backend/engine/invokers"
	"github.com/temphia/temphia/code/core/backend/libx/lazydata"
	"github.com/temphia/temphia/code/core/backend/xtypes"
)

func (r *WebRequest) webHandle(method string, data xtypes.LazyData) (xtypes.LazyData, error) {
	switch method {

	case "http.get_client_ip":
		return lazydata.NewAnyData(r.rctx.ClientIP()), nil
	default:
		return nil, invokers.ErrInvokerActionNotImplemented
	}

}
