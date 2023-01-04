package notz

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (am *AdapterManager) run() {
	pp.Println("@run")

	err := am.init()
	if err != nil {
		pp.Println(err)
	}

	pp.Println("@after_init", am.activeDomains)

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

func (am *AdapterManager) init() error {

	tenants := am.app.StaticTenants()
	for _, tenantId := range tenants {
		err := am.buildTenant(tenantId)
		if err != nil {
			return err
		}
	}

	return nil

}

func (am *AdapterManager) buildTenant(tenantId string) error {

	domains, err := am.corehub.ListDomain(tenantId)
	if err != nil {
		return err
	}

	for _, td := range domains {

		pp.Println("@building => ", tenantId, td)
		pp.Println(am.build(tenantId, td))
	}

	am.tenantInits[tenantId] = true

	return nil
}

func (am *AdapterManager) build(tenantId string, model *entities.TenantDomain) error {

	builder := am.adapterBuilders[model.AdapterType]
	if builder == nil {
		pp.Println("@builder nil, model =>", model)
		return easyerr.NotFound()
	}

	adpr, err := builder(httpx.BuilderOptions{
		App:      am.app,
		TenantId: tenantId,
		Domain:   model,
	})
	if err != nil {
		pp.Println("ERR WHEN BUILDING")
		return err
	}

	am.activeDomains[model.Id] = &DomainInstance{
		adapter: adpr,
		model:   model,
	}
	// "<host>|<tenant>"

	am.domainTenantIndex[model.Name+"|"+tenantId] = model.Id

	return nil

}
