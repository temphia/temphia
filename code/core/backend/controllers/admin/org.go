package admin

import (
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (c *Controller) GetTenant(uclaim *claim.Session) (*entities.Tenant, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotAuthorized()
	}
	return c.coredb.GetTenant(uclaim.TenentId)
}

func (c *Controller) UpdateTenant(uclaim *claim.Session, data map[string]any) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotAuthorized()
	}

	delete(data, "slug")

	return c.coredb.UpdateTenant(uclaim.TenentId, data)
}

// domain

func (c *Controller) AddDomain(uclaim *claim.Session, domain *entities.TenantDomain) error {
	domain.TenantId = uclaim.TenentId
	return c.coredb.AddDomain(domain)
}

func (c *Controller) UpdateDomain(uclaim *claim.Session, id int64, data map[string]any) error {
	return c.coredb.UpdateDomain(uclaim.TenentId, id, data)
}

func (c *Controller) GetDomain(uclaim *claim.Session, id int64) (*entities.TenantDomain, error) {
	return c.coredb.GetDomain(uclaim.TenentId, id)
}

func (c *Controller) RemoveDomain(uclaim *claim.Session, id int64) error {
	return c.coredb.RemoveDomain(uclaim.TenentId, id)
}

func (c *Controller) ListDomain(uclaim *claim.Session) ([]*entities.TenantDomain, error) {
	return c.coredb.ListDomain(uclaim.TenentId)
}
