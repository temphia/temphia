package web2agent

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type Web2Agent struct {
	app      xtypes.App
	tenantId string
	domain   *entities.TenantDomain
	handle   httpx.AdapterHandle

	handlePlug      string
	handleAgent     string
	handleTemplates map[string]string
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	return &Web2Agent{
		app:             opts.App,
		tenantId:        opts.TenantId,
		domain:          opts.Domain,
		handle:          opts.Handle,
		handlePlug:      "",
		handleAgent:     "",
		handleTemplates: make(map[string]string),
	}, nil
}

func (w *Web2Agent) ServeEditorFile(file string) ([]byte, error) {
	return nil, nil
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

/*

expected web2agent iface
	- action_handle_http
	- action_accept_ws
	- action_handle_ws
	- action_template_func

*/
