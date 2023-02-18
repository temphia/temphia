package pageform

import (
	"encoding/json"

	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
)

type Pageform struct {
	builder *PfBuilder
	model   *FormModel
	runtime *goja.Runtime
	binder  bindx.Bindings
}

func (pf *Pageform) Process(ev *event.Request) (*event.Response, error) {

	pf.binder.Log("@test someting someting")

	var resp any
	var err error

	switch ev.Name {
	case "load":
		req := LoadRequest{}
		err = json.Unmarshal(ev.Data, &req)
		if err != nil {
			return nil, err
		}
		resp, err = pf.actionLoad(req)

	case "submit":
		req := SubmitRequest{}
		err = json.Unmarshal(ev.Data, &req)
		if err != nil {
			return nil, err
		}
		resp, err = pf.actionSubmit(req)
	default:
		return nil, easyerr.NotFound()
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
