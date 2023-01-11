package corehub

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

func (c *CoreHub) AddTenant(tenant *entities.Tenant) error {
	err := c.coredb.AddTenant(tenant)
	if err != nil {
		return err
	}

	if c.cplane != nil {
		// fixme => remove this
		eb := c.cplane.GetEventBus()
		eb.EmitTenantEvent(tenant.Slug, xplane.EventCreateTenant, tenant)
	}

	return nil
}

func (c *CoreHub) UpdateTenant(slug string, data map[string]any) error {
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

func (c *CoreHub) UpdateDomain(tenantId string, id int64, data map[string]any) error {
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
