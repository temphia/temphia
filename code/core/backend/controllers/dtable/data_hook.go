package dtable

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (c *Controller) NewHook(uclaim *claim.Session, tslug string, model *entities.DataHook) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)

	model.GroupID = uclaim.Path[2]
	model.TableID = tslug
	model.TenantID = uclaim.TenentId
	return dynDB.NewHook(model)
}

func (c *Controller) ModifyHook(uclaim *claim.Session, tslug string, id int64, data map[string]interface{}) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)

	return dynDB.ModifyHook(uclaim.Path[2], tslug, id, data)
}

func (c *Controller) ListHook(uclaim *claim.Session, tslug string) ([]*entities.DataHook, error) {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)

	return dynDB.ListHook(uclaim.Path[2], tslug)
}

func (c *Controller) DelHook(uclaim *claim.Session, tslug string, id int64) error {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.DelHook(uclaim.Path[2], tslug, id)
}

func (c *Controller) GetHook(uclaim *claim.Session, tslug string, id int64) (*entities.DataHook, error) {
	dynDB := c.dynHub.GetSource(uclaim.Path[1], uclaim.TenentId)
	return dynDB.GetHook(uclaim.Path[2], tslug, id)
}
