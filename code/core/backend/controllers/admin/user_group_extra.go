package admin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

// auth

func (c *Controller) AddUserGroupAuth(uclaim *claim.Session, gslug string, data *entities.UserGroupAuth) error {
	data.TenantId = uclaim.TenentId
	data.UserGroup = gslug
	return c.coredb.AddUserGroupAuth(data)
}

func (c *Controller) UpdateUserGroupAuth(uclaim *claim.Session, gslug string, id int64, data map[string]any) error {
	return c.coredb.UpdateUserGroupAuth(uclaim.TenentId, gslug, id, data)
}

func (c *Controller) ListUserGroupAuth(uclaim *claim.Session, gslug string) ([]*entities.UserGroupAuth, error) {
	return c.coredb.ListUserGroupAuth(uclaim.TenentId, gslug)
}

func (c *Controller) GetUserGroupAuth(uclaim *claim.Session, gslug string, id int64) (*entities.UserGroupAuth, error) {
	return c.coredb.GetUserGroupAuth(uclaim.TenentId, gslug, id)
}

func (c *Controller) RemoveUserGroupAuth(uclaim *claim.Session, gslug string, id int64) error {
	return c.coredb.RemoveUserGroupAuth(uclaim.TenentId, gslug, id)
}

// hook

func (c *Controller) AddUserGroupHook(uclaim *claim.Session, gslug string, data *entities.UserGroupHook) error {
	data.TenantId = uclaim.TenentId
	data.UserGroup = gslug
	return c.coredb.AddUserGroupHook(data)
}

func (c *Controller) UpdateUserGroupHook(uclaim *claim.Session, gslug string, id int64, data map[string]any) error {
	return c.coredb.UpdateUserGroupHook(uclaim.TenentId, gslug, id, data)
}

func (c *Controller) ListUserGroupHook(uclaim *claim.Session, gslug string) ([]*entities.UserGroupHook, error) {
	return c.coredb.ListUserGroupHook(uclaim.TenentId, gslug)
}

func (c *Controller) GetUserGroupHook(uclaim *claim.Session, gslug string, id int64) (*entities.UserGroupHook, error) {
	return c.coredb.GetUserGroupHook(uclaim.TenentId, gslug, id)
}

func (c *Controller) RemoveUserGroupHook(uclaim *claim.Session, gslug string, id int64) error {
	return c.coredb.RemoveUserGroupHook(uclaim.TenentId, gslug, id)
}

// plug

func (c *Controller) AddUserGroupPlug(uclaim *claim.Session, gslug string, data *entities.UserGroupPlug) error {
	data.TenantId = uclaim.TenentId
	data.UserGroup = gslug
	return c.coredb.AddUserGroupPlug(data)
}

func (c *Controller) UpdateUserGroupPlug(uclaim *claim.Session, gslug string, id int64, data map[string]any) error {
	return c.coredb.UpdateUserGroupPlug(uclaim.TenentId, gslug, id, data)
}

func (c *Controller) ListUserGroupPlug(uclaim *claim.Session, gslug string) ([]*entities.UserGroupPlug, error) {
	return c.coredb.ListUserGroupPlug(uclaim.TenentId, gslug)
}

func (c *Controller) GetUserGroupPlug(uclaim *claim.Session, gslug string, id int64) (*entities.UserGroupPlug, error) {
	return c.coredb.GetUserGroupPlug(uclaim.TenentId, gslug, id)
}

func (c *Controller) RemoveUserGroupPlug(uclaim *claim.Session, gslug string, id int64) error {
	return c.coredb.RemoveUserGroupPlug(uclaim.TenentId, gslug, id)
}

// data

func (c *Controller) AddUserGroupData(uclaim *claim.Session, gslug string, data *entities.UserGroupData) error {
	data.TenantId = uclaim.TenentId
	data.UserGroup = gslug
	return c.coredb.AddUserGroupData(data)
}

func (c *Controller) UpdateUserGroupData(uclaim *claim.Session, gslug string, id int64, data map[string]any) error {
	return c.coredb.UpdateUserGroupData(uclaim.TenentId, gslug, id, data)
}

func (c *Controller) ListUserGroupData(uclaim *claim.Session, gslug string) ([]*entities.UserGroupData, error) {
	return c.coredb.ListUserGroupData(uclaim.TenentId, gslug)
}

func (c *Controller) GetUserGroupData(uclaim *claim.Session, gslug string, id int64) (*entities.UserGroupData, error) {
	return c.coredb.GetUserGroupData(uclaim.TenentId, gslug, id)
}

func (c *Controller) RemoveUserGroupData(uclaim *claim.Session, gslug string, id int64) error {
	return c.coredb.RemoveUserGroupData(uclaim.TenentId, gslug, id)
}
