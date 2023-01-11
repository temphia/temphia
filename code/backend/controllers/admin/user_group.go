package admin

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *Controller) AddUserGroup(uclaim *claim.Session, ugroup *entities.UserGroup) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.AddUserGroup(ugroup)
}

func (c *Controller) ListUserGroup(uclaim *claim.Session) ([]*entities.UserGroup, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.ListUserGroups(uclaim.TenentId)
}

func (c *Controller) GetUserGroup(uclaim *claim.Session, ugroup string) (*entities.UserGroup, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.GetUserGroup(uclaim.TenentId, ugroup)
}

func (c *Controller) UpdateUserGroup(uclaim *claim.Session, id string, data map[string]any) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.UpdateUserGroup(uclaim.TenentId, id, data)
}

func (c *Controller) RemoveUserGroup(uclaim *claim.Session, ugroup string) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.RemoveUserGroup(uclaim.TenentId, ugroup)
}
