package wizard

import (
	"encoding/json"

	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/executors/backend/wizard/wmodels"

	"github.com/ztrue/tracerr"

	_ "embed"
)

var (
	//go:embed embed/wizard.loader.js
	loaderJS []byte

	//go:embed embed/wizard.loader.css
	loaderCSS []byte

	DevPath = "./code/executors/frontend/public/build/"
)

type SimpleWizard struct {
	model         wmodels.Wizard
	binding       bindx.Bindings
	jsRuntime     *goja.Runtime
	nativeScripts map[string]registry.DynamicScript
}

func (s *SimpleWizard) Process(ev *event.Request) (*event.Response, error) {

	var resp interface{}
	var err error

	switch ev.Name {
	case "get_splash":
		resp, err = s.GetSplash(ev, "")
	case "run_start":
		resp, err = s.RunStart(ev)
	case "run_nested_start":
		resp, err = s.RunNestedStart(ev)
	case "run_back":
		resp, err = s.RunBack(ev)
	case "run_next":
		resp, err = s.RunNext(ev)
	default:
		return nil, easyerr.NotImpl()
	}

	if err != nil {
		return nil, err
	}

	out, err := json.Marshal(resp)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return &event.Response{
		Payload: (out),
	}, nil

}
