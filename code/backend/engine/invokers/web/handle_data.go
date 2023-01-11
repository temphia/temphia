package web

import (
	"github.com/temphia/temphia/code/core/backend/engine/invokers"
	"github.com/temphia/temphia/code/core/backend/xtypes"
)

func (r *WebRequest) dataHandle(method string, data xtypes.LazyData) (xtypes.LazyData, error) {

	if r.claim.ExecType != invokers.TypeDtableServerHook && r.claim.ExecType != invokers.TypeDtableClientHook {
		return nil, invokers.ErrInvokerActionNotAllowed
	}

	switch method {
	case "data.simple_query":
		return nil, nil
	default:
		return nil, invokers.ErrInvokerActionNotImplemented
	}

}
