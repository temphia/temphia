package web2agent

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type Web2Agent struct {
	app      xtypes.App
	tenantId string
	domain   *entities.TenantDomain
	handle   httpx.AdapterHandle
	engine   etypes.Engine

	mainHook  *entities.TargetHook
	databox   xtypes.DataBox
	initError string
	intOk     bool

	state WAState
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	deps := opts.App.GetDeps()

	wa := &Web2Agent{
		app:       opts.App,
		tenantId:  opts.TenantId,
		domain:    opts.Domain,
		handle:    opts.Handle,
		engine:    deps.Engine().(etypes.Engine),
		databox:   opts.App.Data(),
		intOk:     false,
		initError: "Not initilized",
		mainHook:  nil,
		state: WAState{
			templates:     make(map[string]string),
			routes:        make(map[string]string),
			templateFuncs: make(map[string]string),
		},
	}

	wa.init()

	return wa, nil
}

func (w *Web2Agent) ServeEditorFile(file string) ([]byte, error) {

	switch file {
	case "main.js":
		return w.databox.GetAsset("build", "adapter_editor_web2agent.js")
	case "main.css":
		return w.databox.GetAsset("build", "adapter_editor_web2agent.css")
	}

	return []byte(``), nil
}

func (w *Web2Agent) PreformEditorAction(name string, data []byte) (any, error) {
	return nil, nil
}

func (w *Web2Agent) Handle(ctx httpx.Context) {

	target := WATarget{
		adapter: w,
		rid:     ctx.Rid,
		http:    ctx.Http,
	}

	target.handle()
}
