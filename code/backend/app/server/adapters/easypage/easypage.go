package easypage

import (
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type EasyPage struct {
	options httpx.BuilderOptions
	dataBox xtypes.DataBox
	ahandle httpx.AdapterHandle
	cabHub  store.CabinetHub
	signer  service.Signer

	DomainId int64

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
		signer:    deps.Signer().(service.Signer),
		DomainId:  opts.Domain.Id,
	}, nil
}

func (e *EasyPage) ServeEditorFile(file string) ([]byte, error) {
	return e.serveEditorFile(file)
}

func (e *EasyPage) PreformEditorAction(uclaim *claim.UserContext, name string, data []byte) (any, error) {
	return e.preformEditorAction(uclaim, name, data)
}

func (e *EasyPage) Handle(ctx httpx.Context) {
	e.handle(ctx)
}

func (e *EasyPage) Close() error { return nil }
