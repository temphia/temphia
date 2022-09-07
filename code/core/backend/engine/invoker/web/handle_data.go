package web

import (
	"github.com/temphia/temphia/code/core/backend/engine/invoker"
	"github.com/temphia/temphia/code/core/backend/xtypes"
)

func (r *WebRequest) dataHandle(method string, data xtypes.LazyData) (xtypes.LazyData, error) {

	if r.claim.ExecType != invoker.TypeDtableServerHook && r.claim.ExecType != invoker.TypeDtableClientHook {
		return nil, invoker.ErrInvokerActionNotAllowed
	}

	switch method {
	case "data.simple_query":
		return nil, nil
	default:
		return nil, invoker.ErrInvokerActionNotImplemented
	}

}
