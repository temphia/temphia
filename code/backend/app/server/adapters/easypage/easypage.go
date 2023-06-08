package easypage

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/app/server/adapters/common/autotarget"
	"github.com/temphia/temphia/code/backend/app/server/adapters/common/cache"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type EasyPage struct {
	dataBox xtypes.DataBox
	ahandle httpx.AdapterHandle
	cabHub  store.CabinetHub
	signer  service.Signer
	corehub store.CoreHub

	pkv store.PlugStateKV

	domainId int64
	tenantId string

	bpintId    string
	editorHook *entities.TargetHook

	filecache httpx.SubCache
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

	ep := &EasyPage{
		dataBox:    opts.App.Data(),
		ahandle:    opts.Handle,
		cabHub:     deps.Cabinet().(store.CabinetHub),
		signer:     deps.Signer().(service.Signer),
		domainId:   opts.Domain.Id,
		corehub:    deps.CoreHub().(store.CoreHub),
		bpintId:    target.BprintId(),
		editorHook: target.EditorHooks(),
		tenantId:   opts.TenantId,
		pkv:        deps.PlugKV().(store.PlugStateKV),
		filecache:  nil,
	}

	subcache, err := opts.Cache.GetSubCache(
		fmt.Sprintf("adapter-%d", opts.Domain.Id),
		cache.NewBprintLoader(ep.tenantId, ep.bpintId, deps.RepoHub().(repox.Hub)),
	)

	if err != nil {
		return nil, err
	}

	ep.filecache = subcache

	return ep, nil
}

func (e *EasyPage) ServeEditorFile(file string) ([]byte, error) {
	return e.serveEditorFile(file)
}

func (e *EasyPage) PreformEditorAction(ctx httpx.AdapterEditorContext) (any, error) {
	return e.preformEditorAction(ctx.User, ctx.Name, ctx.Data)
}

func (e *EasyPage) Handle(ctx httpx.Context) {
	e.handle(ctx)
}

func (e *EasyPage) Close() error {
	return nil
}

/*

	events
		on:page_add
		on:page_modify
		on:load_data
		on:before_build
		on:after_build


*/
