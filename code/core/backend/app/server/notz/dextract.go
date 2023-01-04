package notz

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/app/server/notz/dnstoken"
	"github.com/temphia/temphia/code/core/backend/app/server/static"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
)

func (m *Notz) extract(c *gin.Context) (string, string, error) {
	hostname := c.Request.URL.Hostname()

	tenantId := ""
	if m.resolveHostTenantFn != nil {
		tenantId = m.resolveHostTenantFn(hostname)
	}

	if m.app.SingleTenant() {
		return m.app.StaticTenants()[0], hostname, nil
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
