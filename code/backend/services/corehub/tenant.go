package corehub

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *CoreHub) AddTenant(tenant *entities.Tenant) error {
	err := c.coredb.AddTenant(tenant)
	if err != nil {
		return err
	}

	c.stateHub.OnTenantChange(tenant.Slug, tenant)

	return nil
}

func (c *CoreHub) UpdateTenant(slug string, data map[string]any) error {
	err := c.coredb.UpdateTenant(slug, data)
	if err != nil {
		return err
	}

	c.stateHub.OnTenantChange(slug, data)

	return nil

}

func (c *CoreHub) GetTenant(tenant string) (*entities.Tenant, error) {
	return c.coredb.GetTenant(tenant)
}

func (c *CoreHub) RemoveTenant(slug string) error {
	err := c.coredb.RemoveTenant(slug)
	if err != nil {
		return err
	}

	c.stateHub.OnTenantChange(slug, nil)
	return nil
}

func (c *CoreHub) ListTenant() ([]*entities.Tenant, error) {
	return c.coredb.ListTenant()
}

// domain

func (c *CoreHub) AddDomain(domain *entities.TenantDomain) (int64, error) {
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
