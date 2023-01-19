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
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	return &Web2Agent{
		app:      opts.App,
		tenantId: opts.TenantId,
		domain:   opts.Domain,
		handle:   opts.Handle,
	}, nil
}

func (w *Web2Agent) ServeEditorFile(file string) ([]byte, error) {
	return nil, nil
}

func (w *Web2Agent) PreformEditorAction(name string, data []byte) (any, error) {
	return nil, nil
}

func (w *Web2Agent) Handle(ctx httpx.Context) {

}
