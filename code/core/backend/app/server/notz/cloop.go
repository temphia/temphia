package notz

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (am *AdapterManager) run() {

	err := am.init()
	if err != nil {
		pp.Println(err)
	}

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
	if am.app.SingleTenant() {
		return am.buildTenant(am.app.TenantId())
	}

	return nil

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
		return
	}

	adpr, err := builder(httpx.BuilderOptions{
		App:      am.app,
		TenantId: tenantId,
		Domain:   model,
	})

	if err != nil {
		return
	}

	am.activeDomains[model.Id] = &DomainInstance{
		adapter: adpr,
		model:   model,
	}

	if model.Name == "*" {
		am.domainTenantIndex["*/"+tenantId] = model.Id
	} else {
		am.domainTenantIndex[model.Name] = model.Id
	}

}
