package admin

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/models/scopes"
)

func (c *Controller) NewView(uclaim *claim.Session, source, group, tslug string, model *entities.DataView) error {
	if !c.HasScope(uclaim, "data") {
		return scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetDynDB()

	model.GroupID = group
	model.TableID = tslug
	model.TenantID = uclaim.TenantId

	pp.Println(model)

	return dynDB.NewView(uclaim.TenantId, model)
}

func (c *Controller) ModifyView(uclaim *claim.Session, source, group, tslug string, id int64, data map[string]interface{}) error {
	if !c.HasScope(uclaim, "data") {
		return scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetDynDB()

	return dynDB.ModifyView(uclaim.TenantId, group, tslug, id, data)
}

func (c *Controller) ListView(uclaim *claim.Session, source, group, tslug string) ([]*entities.DataView, error) {
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetDynDB()

	return dynDB.ListView(uclaim.TenantId, group, tslug)
}

func (c *Controller) DelView(uclaim *claim.Session, source, group, tslug string, id int64) error {
	if !c.HasScope(uclaim, "data") {
		return scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetDynDB()

	return dynDB.DelView(uclaim.TenantId, group, tslug, id)
}

func (c *Controller) GetView(uclaim *claim.Session, source, group, tslug string, id int64) (*entities.DataView, error) {
	if !c.HasScope(uclaim, "data") {
		return nil, scopes.ErrNoAdminDataScope
	}

	dynDB := c.dynHub.GetDynDB()

	return dynDB.GetView(uclaim.TenantId, group, tslug, id)
}
