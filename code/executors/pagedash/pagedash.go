package pagedash

import (
	"encoding/json"

	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
)

type PageDash struct {
	builder   *PdBuilder
	jsruntime *goja.Runtime
	binder    bindx.Bindings
	model     *DashModel
}

func (pd *PageDash) Process(ev *event.Request) (*event.Response, error) {

	var resp any
	var err error

	switch ev.Name {
	case "load":
		req := LoadRequest{}
		err = json.Unmarshal(ev.Data, &req)
		if err != nil {
			return nil, err
		}

		pp.Println("@model", pd.model)
		pp.Println("@load", req)

		resp, err = pd.actionLoad(req)
		pp.Println("@resp", resp)

	case "build":

		req := BuildRequest{}
		err = json.Unmarshal(ev.Data, &req)
		if err != nil {
			return nil, err
		}
		resp, err = pd.actionBuild(req)
	default:
		return nil, easyerr.NotFound()
	}

	if err != nil {
		return nil, err
	}

	out, err := json.Marshal(resp)
	if err != nil {
		pp.Println("@whatttttt", err.Error())
		return nil, err
	}

	return &event.Response{
		Payload: out,
	}, nil

}
