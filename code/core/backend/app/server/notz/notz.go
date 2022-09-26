package notz

// notz handles all routes except /z/

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/app/server/notz/dnstoken"
	"github.com/temphia/temphia/code/core/backend/app/server/static"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
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
	corehub             store.CoreHub
	cabinethub          store.CabinetHub

	rendererBuilders map[string]httpx.Builder
	renderers        map[string]httpx.Adapter
	rlock            sync.Mutex // DANGER => cannot copy manager cz its not pointer
}

func New(opts NotzOptions) Notz {
	deps := opts.App.GetDeps()

	return Notz{
		staticHosts:         opts.StaticHosts,
		app:                 opts.App,
		resolveHostTenantFn: opts.ResolveHostTenant,
		rootHost:            opts.RootHost,
		tenantHostBase:      opts.TenantHostBase,
		renderers:           make(map[string]httpx.Adapter),
		rlock:               sync.Mutex{},
		rendererBuilders:    deps.Registry().(*registry.Registry).GetAdapterBuilders(),
		corehub:             deps.CoreHub().(store.CoreHub),
		cabinethub:          deps.Cabinet().(store.CabinetHub),
	}
}

func (m *Notz) Serve(c *gin.Context) {
	tenantId, hostname, err := m.extract(c)
	if err != nil {
		return
	}

	m.runRenderer(tenantId, hostname, c)
}

func (m *Notz) extract(c *gin.Context) (string, string, error) {
	hostname := c.Request.URL.Hostname()

	tenantId := ""
	if m.resolveHostTenantFn != nil {
		tenantId = m.resolveHostTenantFn(hostname)
	}

	if m.app.SingleTenant() {

		return m.app.TenantId(), hostname, easyerr.NotFound()
	}

	if m.tenantHostBase != "" && strings.HasSuffix(hostname, m.tenantHostBase) {
		// tenant1.example.com
		tenantId := strings.TrimRight(strings.TrimRight(hostname, m.tenantHostBase), ".")
		return tenantId, hostname, nil
	}

	if hostname == "" || hostname == "localhost" || hostname == m.rootHost {
		c.Writer.Write(static.Root)
		return "", "", easyerr.NotFound()
	}

	if tenantId == "" {
		tenantId = m.staticHosts[hostname]
	}

	if tenantId == "" {
		tenantId = m.resolveTenant(hostname, c)
	}

	if tenantId == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return "", "", easyerr.NotFound()
	}

	return tenantId, hostname, nil

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

func (m *Notz) ListRenderers() []string {

	keys := make([]string, 0, len(m.rendererBuilders)+1)
	for k := range m.rendererBuilders {
		keys = append(keys, k)
	}

	keys = append(keys, "alias")

	return keys
}

// private

func (m *Notz) runRenderer(tenantId, host string, ctx *gin.Context) {
	r := m.get(tenantId, host)
	if r == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	r.Handle(httpx.Context{
		Http: ctx,
		Rid:  0, // fixme =>
	})
}

func (m *Notz) resolveTenant(host string, c *gin.Context) string {
	tenantId := ""
	if tenantId == "" {
		tenantId, _ = c.Cookie("tenant_id")
	}

	if tenantId == "" {
		tenantId, _ = dnstoken.DNSReverseResolve(m.app.ClusterId(), host)
	}

	return tenantId
}
