package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/scopes"
)

func (c *Controller) GetTenant(uclaim *claim.Session) (*entities.Tenant, error) {
	if !c.HasScope(uclaim, "tenant") {
		return nil, scopes.ErrNoAdminTenantScope
	}

	return c.coredb.GetTenant(uclaim.TenantId)
}

func (c *Controller) UpdateTenant(uclaim *claim.Session, data map[string]any) error {
	if !c.HasScope(uclaim, "tenant") {
		return scopes.ErrNoAdminTenantScope
	}

	delete(data, "slug")

	return c.coredb.UpdateTenant(uclaim.TenantId, data)
}

// domain

func (c *Controller) AddDomain(uclaim *claim.Session, domain *entities.TenantDomain) error {
	if !c.HasScope(uclaim, "tenant") {
		return scopes.ErrNoAdminTenantScope
	}

	domain.TenantId = uclaim.TenantId
	_, err := c.coredb.AddDomain(domain)
	return err
}

func (c *Controller) UpdateDomain(uclaim *claim.Session, id int64, data map[string]any) error {
	if !c.HasScope(uclaim, "tenant") {
		return scopes.ErrNoAdminTenantScope
	}

	return c.coredb.UpdateDomain(uclaim.TenantId, id, data)
}

func (c *Controller) GetDomain(uclaim *claim.Session, id int64) (*entities.TenantDomain, error) {
	if !c.HasScope(uclaim, "tenant") {
		return nil, scopes.ErrNoAdminTenantScope
	}

	return c.coredb.GetDomain(uclaim.TenantId, id)
}

func (c *Controller) RemoveDomain(uclaim *claim.Session, id int64) error {
	if !c.HasScope(uclaim, "tenant") {
		return scopes.ErrNoAdminTenantScope
	}

	return c.coredb.RemoveDomain(uclaim.TenantId, id)
}

func (c *Controller) ListDomain(uclaim *claim.Session) ([]*entities.TenantDomain, error) {
	if !c.HasScope(uclaim, "tenant") {
		return nil, scopes.ErrNoAdminTenantScope
	}

	return c.coredb.ListDomain(uclaim.TenantId)
}
