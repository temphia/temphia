package pagequery

import (
	"encoding/json"

	"github.com/dop251/goja"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type PageQuery struct {
	builder   *PgBuilder
	model     *PgModel
	jsruntime *goja.Runtime
	binder    bindx.Bindings
	tenantId  string
	datahub   dyndb.DataHub
}

func (pf *PageQuery) Process(ev *event.Request) (*event.Response, error) {

	pf.binder.Log("@test someting someting")

	var resp any
	var err error

	switch ev.Name {
	case "load":
		req := &LoadRequest{}
		err = json.Unmarshal(ev.Data, req)
		if err != nil {
			return nil, err
		}
		resp, err = pf.load(req)
	case "submit":
		req := &SubmitRequest{}
		err = json.Unmarshal(ev.Data, req)
		if err != nil {
			return nil, err
		}
		resp, err = pf.submit(req)
	default:
		return nil, easyerr.NotFound("event action")
	}

	if err != nil {
		return nil, err
	}

	out, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return &event.Response{
		Payload: out,
	}, nil
}
