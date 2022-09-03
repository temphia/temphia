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

func (c *Controller) UpdateTenant(uclaim *claim.Session, data map[string]interface{}) error {
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

func (c *Controller) UpdateDomain(uclaim *claim.Session, id int64, data map[string]interface{}) error {
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

// widget

func (c *Controller) AddDomainWidget(uclaim *claim.Session, domain *entities.DomainWidget) error {
	domain.TenantId = uclaim.TenentId
	return c.coredb.AddDomainWidget(domain)
}

func (c *Controller) UpdateDomainWidget(uclaim *claim.Session, id int64, data map[string]interface{}) error {
	return c.coredb.UpdateDomainWidget(uclaim.TenentId, id, data)
}

func (c *Controller) GetDomainWidget(uclaim *claim.Session, id int64) (*entities.DomainWidget, error) {
	return c.coredb.GetDomainWidget(uclaim.TenentId, id)
}

func (c *Controller) RemoveDomainWidget(uclaim *claim.Session, id int64) error {
	return c.coredb.RemoveDomainWidget(uclaim.TenentId, id)
}

func (c *Controller) ListDomainWidget(uclaim *claim.Session, did int64) ([]*entities.DomainWidget, error) {
	return c.coredb.ListDomainWidget(uclaim.TenentId, did)
}
