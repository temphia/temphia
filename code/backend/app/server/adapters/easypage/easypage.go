package easypage

import (
	"sync"

	"github.com/temphia/temphia/code/backend/app/server/adapters/common/autotarget"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type EasyPage struct {
	options httpx.BuilderOptions
	dataBox xtypes.DataBox
	ahandle httpx.AdapterHandle
	cabHub  store.CabinetHub
	signer  service.Signer
	corehub store.CoreHub

	domainId int64

	bpintId    string
	editorHook *entities.TargetHook

	pageCache map[string][]byte
	pLock     sync.Mutex
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	deps := opts.App.GetDeps()

	target := autotarget.New(opts.TenantId, opts.Domain.Id, deps.CoreHub().(store.CoreHub))

	ok, err := target.IsInit()
	if err != nil {
		return nil, err
	}

	if !ok {
		err = target.AutoInit()
		if err != nil {
			return nil, err
		}
	}

	return &EasyPage{
		options:    opts,
		dataBox:    opts.App.Data(),
		ahandle:    opts.Handle,
		pageCache:  make(map[string][]byte),
		pLock:      sync.Mutex{},
		cabHub:     deps.Cabinet().(store.CabinetHub),
		signer:     deps.Signer().(service.Signer),
		domainId:   opts.Domain.Id,
		corehub:    deps.CoreHub().(store.CoreHub),
		bpintId:    target.BprintId(),
		editorHook: target.EditorHooks(),
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

/*

	events
		on:page_add
		on:page_modify
		on:load_data
		on:before_build
		on:after_build


*/
