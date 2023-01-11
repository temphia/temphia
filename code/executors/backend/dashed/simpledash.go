package dashed

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/executors/backend/dashed/dashmodels"
)

var _ etypes.Executor = (*SimpleDash)(nil)

type SimpleDash struct {
	bindings bindx.Bindings
	model    dashmodels.Dashboard
}

func (s *SimpleDash) Process(ev *event.Request) (*event.Response, error) {

	switch ev.Name {
	case "generate":
		return s.generate(ev)
	default:
		return nil, easyerr.NotImpl()
	}
}
