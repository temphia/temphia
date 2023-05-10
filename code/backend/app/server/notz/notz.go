package notz

// notz handles all routes except /z/

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type NotzOptions struct {
	App               xtypes.App
	StaticHosts       map[string]string
	ResolveHostTenant func(host string) string
	RootHost          string
	TenantHostBase    string
}

var _ httpx.AdapterHub = (*Notz)(nil)

type Notz struct {
	app                 xtypes.App
	staticHosts         map[string]string
	resolveHostTenantFn func(host string) string
	rootHost            string
	tenantHostBase      string
	cabinethub          store.CabinetHub
	adapterManager      AdapterManager
}

func New(opts NotzOptions) *Notz {

	am := newAdapterManager(opts.App)

	n := &Notz{
		staticHosts:         opts.StaticHosts,
		app:                 opts.App,
		resolveHostTenantFn: opts.ResolveHostTenant,
		rootHost:            opts.RootHost,
		tenantHostBase:      opts.TenantHostBase,
		cabinethub:          nil,
		adapterManager:      am,
	}

	go n.adapterManager.run()

	return n
}

func (m *Notz) Serve(ctx *gin.Context) {
	m.serve(ctx)
}

func (m *Notz) ServePublic(c *gin.Context, file string) {
	m.servePublic(c, file)
}

func (m *Notz) IsAllowed(tenantId, host string) bool {
	return m.isAllowed(tenantId, host)
}

func (m *Notz) ApplyTargetHook(tenantId string, id int64, data *entities.TargetHook) {

}

func (m *Notz) ApplyAdapter(tenantId string, id int64, data *entities.TenantDomain) {

}

func (m *Notz) ListAdapters() []string {
	return m.adapterManager.ListAdapters()
}

func (m *Notz) Reset(tenantId string, domainId int64) {
	m.adapterManager.cReInstance <- DomainIdent{
		tenantId: tenantId,
		domainId: domainId,
	}
}

func (m *Notz) ServeEditorFile(tenantId, file string, did int64, ctx *gin.Context) {
	out, err := m.adapterManager.serveEditorFile(tenantId, did, file)
	if err != nil {
		return
	}

	httpx.WriteFile(file, out, ctx)
}

func (m *Notz) PreformEditorAction(aclaim *claim.AdapterEditor, name string, did int64, ctx *gin.Context) {
	out, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	resp, err := m.adapterManager.preformEditorAction(aclaim, name, did, out)
	httpx.WriteJSON(ctx, resp, err)
}
