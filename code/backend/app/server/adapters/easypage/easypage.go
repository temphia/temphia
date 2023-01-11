package easypage

import (
	"sync"

	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type EasyPage struct {
	options httpx.BuilderOptions
	dataBox xtypes.DataBox
	ahandle httpx.AdapterHandle
	cabHub  store.CabinetHub

	pageCache map[string][]byte
	pLock     sync.Mutex
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	deps := opts.App.GetDeps()

	return &EasyPage{
		options:   opts,
		dataBox:   opts.App.Data(),
		ahandle:   opts.Handle,
		pageCache: make(map[string][]byte),
		pLock:     sync.Mutex{},
		cabHub:    deps.Cabinet().(store.CabinetHub),
	}, nil
}

func (e *EasyPage) ServeEditorFile(file string) ([]byte, error) {
	return e.serveEditorFile(file)
}

func (e *EasyPage) PreformEditorAction(name string, data []byte) (any, error) {
	return e.preformEditorAction(name, data)
}

func (e *EasyPage) Handle(ctx httpx.Context) {
	e.handle(ctx)
}
