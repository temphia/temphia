package logproxy

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type LogProxy struct {
	app      xtypes.App
	resource *entities.Resource
	logproxy logx.Proxy
}

func (l *LogProxy) Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error) {

	switch method {
	case "query":
		qreq := logx.QueryRequest{}
		err := args.AsObject(&qreq)
		if err != nil {
			return nil, err
		}

		resp, err := l.logproxy.Query(l.resource.TenantId, qreq)
		if err != nil {
			return nil, err
		}
		return lazydata.NewAnyData(resp), nil
	}

	return nil, easyerr.NotImpl()
}

func (l *LogProxy) Close() error {
	return nil
}
