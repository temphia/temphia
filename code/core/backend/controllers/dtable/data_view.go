package dtable

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (c *Controller) NewView(uclaim *claim.Session, tslug string, model *entities.DataView) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)

	model.GroupID = uclaim.Path[2]
	model.TableID = tslug
	model.TenantID = uclaim.TenentId

	return dynDB.NewView(model)
}

func (c *Controller) ModifyView(uclaim *claim.Session, tslug string, id int64, data map[string]interface{}) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)

	return dynDB.ModifyView(uclaim.Path[2], tslug, id, data)
}

func (c *Controller) ListView(uclaim *claim.Session, tslug string) ([]*entities.DataView, error) {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)

	return dynDB.ListView(uclaim.Path[2], tslug)
}

func (c *Controller) DelView(uclaim *claim.Session, tslug string, id int64) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)

	return dynDB.DelView(uclaim.Path[2], tslug, id)
}

func (c *Controller) GetView(uclaim *claim.Session, tslug string, id int64) (*entities.DataView, error) {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.GetView(uclaim.Path[2], tslug, id)
}
