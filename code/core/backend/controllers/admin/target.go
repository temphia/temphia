package admin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (c *Controller) AddTargetHook(uclaim *claim.Session, data *entities.TargetHook) error {
	data.TenantId = uclaim.TenentId
	return c.coredb.AddTargetHook(data)
}

func (c *Controller) UpdateTargetHook(uclaim *claim.Session, ttype string, id int64, data map[string]any) error {
	return c.coredb.UpdateTargetHook(uclaim.TenentId, ttype, id, data)
}

func (c *Controller) ListTargetHook(uclaim *claim.Session) ([]*entities.TargetHook, error) {
	return c.coredb.ListTargetHook(uclaim.TenentId)
}

func (c *Controller) ListTargetHookByType(uclaim *claim.Session, ttype, target string) ([]*entities.TargetHook, error) {
	return c.coredb.ListTargetHookByType(uclaim.TenentId, ttype, target)
}

func (c *Controller) GetTargetHook(uclaim *claim.Session, ttype string, id int64) (*entities.TargetHook, error) {
	return c.coredb.GetTargetHook(uclaim.TenentId, ttype, id)
}

func (c *Controller) RemoveTargetHook(uclaim *claim.Session, ttype string, id int64) error {
	return c.coredb.RemoveTargetHook(uclaim.TenentId, ttype, id)
}

func (c *Controller) AddTargetApp(uclaim *claim.Session, data *entities.TargetApp) error {
	data.TenantId = uclaim.TenentId
	return c.coredb.AddTargetApp(data)
}

func (c *Controller) UpdateTargetApp(uclaim *claim.Session, ttype string, id int64, data map[string]any) error {
	return c.coredb.UpdateTargetApp(uclaim.TenentId, ttype, id, data)
}

func (c *Controller) ListTargetApp(uclaim *claim.Session) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetApp(uclaim.TenentId)
}

func (c *Controller) ListTargetAppByType(uclaim *claim.Session, ttype, target string) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetAppByType(uclaim.TenentId, ttype, target)
}

func (c *Controller) GetTargetApp(uclaim *claim.Session, ttype string, id int64) (*entities.TargetApp, error) {
	return c.coredb.GetTargetApp(uclaim.TenentId, ttype, id)
}

func (c *Controller) RemoveTargetApp(uclaim *claim.Session, ttype string, id int64) error {
	return c.coredb.RemoveTargetApp(uclaim.TenentId, ttype, id)
}
