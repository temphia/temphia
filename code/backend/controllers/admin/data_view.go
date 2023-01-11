package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *Controller) NewView(uclaim *claim.Session, source, group, tslug string, model *entities.DataView) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)

	model.GroupID = group
	model.TableID = tslug
	model.TenantID = uclaim.TenentId

	return dynDB.NewView(model)
}

func (c *Controller) ModifyView(uclaim *claim.Session, source, group, tslug string, id int64, data map[string]interface{}) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)

	return dynDB.ModifyView(group, tslug, id, data)
}

func (c *Controller) ListView(uclaim *claim.Session, source, group, tslug string) ([]*entities.DataView, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)

	return dynDB.ListView(group, tslug)
}

func (c *Controller) DelView(uclaim *claim.Session, source, group, tslug string, id int64) error {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)

	return dynDB.DelView(group, tslug, id)
}

func (c *Controller) GetView(uclaim *claim.Session, source, group, tslug string, id int64) (*entities.DataView, error) {
	dynDB := c.dynHub.GetSource(source, uclaim.TenentId)
	return dynDB.GetView(group, tslug, id)
}
