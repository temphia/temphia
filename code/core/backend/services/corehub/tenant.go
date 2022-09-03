package corehub

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

func (c *CoreHub) AddTenant(tenant *entities.Tenant) error {
	err := c.coredb.AddTenant(tenant)
	if err != nil {
		return err
	}

	eb := c.cplane.GetEventBus()
	eb.EmitTenantEvent(tenant.Slug, xplane.EventCreateTenant, tenant)
	return nil
}

func (c *CoreHub) UpdateTenant(slug string, data map[string]interface{}) error {
	return c.coredb.UpdateTenant(slug, data)
}

func (c *CoreHub) GetTenant(tenant string) (*entities.Tenant, error) {
	return c.coredb.GetTenant(tenant)
}

func (c *CoreHub) RemoveTenant(slug string) error {
	return c.coredb.RemoveTenant(slug)
}

func (c *CoreHub) ListTenant() ([]*entities.Tenant, error) {
	return c.coredb.ListTenant()
}

// domain

func (c *CoreHub) AddDomain(domain *entities.TenantDomain) error {
	return c.coredb.AddDomain(domain)
}

func (c *CoreHub) UpdateDomain(tenantId string, id int64, data map[string]interface{}) error {
	return c.coredb.UpdateDomain(tenantId, id, data)
}

func (c *CoreHub) GetDomain(tenantId string, id int64) (*entities.TenantDomain, error) {
	return c.coredb.GetDomain(tenantId, id)
}

func (c *CoreHub) GetDomainByName(tenantId string, name string) (*entities.TenantDomain, error) {
	return c.coredb.GetDomainByName(tenantId, name)
}

func (c *CoreHub) RemoveDomain(tenantId string, id int64) error {
	return c.coredb.RemoveDomain(tenantId, id)
}

func (c *CoreHub) ListDomain(tenantId string) ([]*entities.TenantDomain, error) {
	return c.coredb.ListDomain(tenantId)
}

// widget

func (c *CoreHub) AddDomainWidget(widget *entities.DomainWidget) error {
	return c.coredb.AddDomainWidget(widget)
}

func (c *CoreHub) UpdateDomainWidget(tenantId string, id int64, data map[string]interface{}) error {
	return c.coredb.UpdateDomainWidget(tenantId, id, data)
}

func (c *CoreHub) GetDomainWidget(tenantId string, id int64) (*entities.DomainWidget, error) {
	return c.coredb.GetDomainWidget(tenantId, id)
}

func (c *CoreHub) RemoveDomainWidget(tenantId string, id int64) error {
	return c.coredb.RemoveDomainWidget(tenantId, id)
}

func (c *CoreHub) ListDomainWidget(tenantId string, did int64) ([]*entities.DomainWidget, error) {
	return c.coredb.ListDomainWidget(tenantId, did)
}
