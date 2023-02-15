package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *Controller) NewView(uclaim *claim.Session, source, group, tslug string, model *entities.DataView) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	model.GroupID = group
	model.TableID = tslug
	model.TenantID = uclaim.TenantId

	return dynDB.NewView(uclaim.TenantId, model)
}

func (c *Controller) ModifyView(uclaim *claim.Session, source, group, tslug string, id int64, data map[string]interface{}) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.ModifyView(uclaim.TenantId, group, tslug, id, data)
}

func (c *Controller) ListView(uclaim *claim.Session, source, group, tslug string) ([]*entities.DataView, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.ListView(uclaim.TenantId, group, tslug)
}

func (c *Controller) DelView(uclaim *claim.Session, source, group, tslug string, id int64) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDB.DelView(uclaim.TenantId, group, tslug, id)
}

func (c *Controller) GetView(uclaim *claim.Session, source, group, tslug string, id int64) (*entities.DataView, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenantId)
	return dynDB.GetView(uclaim.TenantId, group, tslug, id)
}
