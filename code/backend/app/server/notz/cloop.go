package notz

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	lru "github.com/hashicorp/golang-lru"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/server/notz/adapter"
	"github.com/temphia/temphia/code/backend/app/server/notz/ahandle"
	"github.com/temphia/temphia/code/backend/app/server/notz/dnstoken"
	"github.com/temphia/temphia/code/backend/app/server/static"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (am *AdapterManager) run() {

	am.applogger.Info().Msg(logid.NotzAdaptersBeforeInit)
	am.init()
	am.applogger.Info().Msg(logid.NotzAdaptersAfterInit)

	lcahe, err := lru.New(10)
	if err != nil {
		panic(err)
	}

	for {

		select {
		case data := <-am.cReInstance:
			key := func() string {
				return fmt.Sprintf("%s__%d", data.tenantId, data.domainId)
			}

			lasttime, ok := lcahe.Get(key())
			if ok {
				lt := lasttime.(time.Time)
				sub := time.Until(lt)
				if sub < time.Second*60 {
					return
				}
			}

			domain, err := am.corehub.GetDomain(data.tenantId, data.domainId)
			if err != nil {
				continue
			}

			am.build(data.tenantId, domain)
			lcahe.Add(key(), time.Now())
		case tenantId := <-am.cInstanceTenant:

			pp.Println(am.buildTenant(tenantId))
		}

	}

}

func (am *AdapterManager) init() {

	tenants := am.app.StaticTenants()
	for _, tenantId := range tenants {
		err := am.buildTenant(tenantId)
		if err != nil {
			am.applogger.Error().
				Str("tenant_id", tenantId).
				Msg(logid.NotzAdapterInitErr)
		} else {
			am.applogger.Info().
				Str("tenant_id", tenantId).
				Msg(logid.NotzAdapterInitOk)
		}
	}

}

func (am *AdapterManager) buildTenant(tenantId string) error {

	domains, err := am.corehub.ListDomain(tenantId)
	if err != nil {
		return err
	}

	for _, td := range domains {
		am.build(tenantId, td)

	}

	am.tenantInits[tenantId] = true

	return nil
}

func (am *AdapterManager) resetDomain(tenantId string, domainId int64) {
	// fixme => this will work only on single node
	// send this to other node from msgbus

	am.cReInstance <- DomainIdent{
		tenantId: tenantId,
		domainId: domainId,
	}

}

func (am *AdapterManager) build(tenantId string, model *entities.TenantDomain) {

	builder := am.adapterBuilders[model.AdapterType]
	if builder == nil {
		am.applogger.Error().
			Str("tenant_id", tenantId).
			Str("domain", model.Name).
			Int64("domain_id", model.Id).
			Msg(logid.NotzAdapterBuilderNotFound)
		return
	}

	logger := am.logService.GetSiteLogger(tenantId, model.Name)

	adpr, err := builder(httpx.BuilderOptions{
		App:      am.app,
		TenantId: tenantId,
		Domain:   model,
		Handle: ahandle.New(ahandle.Options{
			App:      am.app,
			Logger:   &logger,
			DomainId: model.Id,
			TenantId: tenantId,
			ResetDomain: func() {
				am.resetDomain(tenantId, model.Id)
			},
		}),
		Cache: am.cache,
	})
	if err != nil {
		am.applogger.Error().
			Str("tenant_id", tenantId).
			Str("domain", model.Name).
			Err(err).
			Int64("domain_id", model.Id).
			Msg(logid.NotzAdapterBuildErr)
	} else {
		am.applogger.Info().
			Str("tenant_id", tenantId).
			Str("domain", model.Name).
			Int64("domain_id", model.Id).
			Msg(logid.NotzAdapterBuildOk)
	}

	am.activeDomains[model.Id] = adapter.New(adpr, model)

	// "<host>|<tenant>"

	am.domainTenantIndex[model.Name+"|"+tenantId] = model.Id

}

func (m *Notz) extract(c *gin.Context) (string, string, error) {
	hostname, _, _ := net.SplitHostPort(c.Request.Host)

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
		return "", "", easyerr.NotFound("host")
	}

	if tenantId == "" {
		tenantId = m.staticHosts[hostname]
	}

	if tenantId == "" {
		tenantId = m.resolveTenant(hostname, c)
	}

	if tenantId == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return "", "", easyerr.NotFound("tenant")
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
