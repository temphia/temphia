package notz

// notz handles all routes except /z/

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type NotzOptions struct {
	App               xtypes.App
	StaticHosts       map[string]string
	ResolveHostTenant func(host string) string
	RootHost          string
	TenantHostBase    string
}

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

	tenantId, hostname, err := m.extract(ctx)
	if err != nil {
		m.adapterManager.applogger.Error().
			Str("tenant_id", tenantId).
			Msg(logid.NotzHostExtractErr)
		return
	}

	m.adapterManager.Handle(tenantId, hostname, ctx)
}

func (m *Notz) ServePublic(c *gin.Context, file string) {
	tenantId, _, err := m.extract(c)
	if err != nil {
		return
	}

	source := m.cabinethub.Default(tenantId)
	out, err := source.GetBlob(c.Request.Context(), "public", file)
	if err != nil {
		return
	}

	c.Writer.Write(out)
}

func (m *Notz) Reset(tenantId string, domainId int64) {
	m.adapterManager.cReInstance <- DomainIdent{
		tenantId: tenantId,
		domainId: domainId,
	}
}

func (m *Notz) ListAdapters() []string {
	return m.adapterManager.ListAdapters()
}

func (m *Notz) ServeEditorFile(tenantId, file string, did int64, ctx *gin.Context) {
	out, err := m.adapterManager.serveEditorFile(tenantId, did, file)
	if err != nil {
		return
	}

	httpx.WriteFile(file, out, ctx)
}

func (m *Notz) PreformEditorAction(tenantId, name string, did int64, ctx *gin.Context) {
	out, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return
	}

	resp, err := m.adapterManager.preformEditorAction(tenantId, name, did, out)
	httpx.WriteJSON(ctx, resp, err)
}
