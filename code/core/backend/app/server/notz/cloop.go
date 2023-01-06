package notz

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/app/server/notz/ahandle"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (am *AdapterManager) run() {

	am.applogger.Info().Msg(logid.NotzAdaptersBeforeInit)
	am.init()
	am.applogger.Info().Msg(logid.NotzAdaptersAfterInit)

	for {

		select {
		case data := <-am.cReInstance:
			domain, err := am.corehub.GetDomain(data.tenantId, data.domainId)
			if err != nil {
				continue
			}

			am.build(data.tenantId, domain)
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
			Corehub:  am.corehub,
			Logger:   &logger,
			DomainId: model.Id,
			TenantId: tenantId,
		}),
	})
	if err != nil {
		am.applogger.Error().
			Str("tenant_id", tenantId).
			Str("domain", model.Name).
			Int64("domain_id", model.Id).
			Msg(logid.NotzAdapterBuildErr)
	} else {
		am.applogger.Info().
			Str("tenant_id", tenantId).
			Str("domain", model.Name).
			Int64("domain_id", model.Id).
			Msg(logid.NotzAdapterBuildOk)
	}

	am.activeDomains[model.Id] = &DomainInstance{
		adapter: adpr,
		model:   model,
	}
	// "<host>|<tenant>"

	am.domainTenantIndex[model.Name+"|"+tenantId] = model.Id

}
