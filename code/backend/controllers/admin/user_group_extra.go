package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/models/scopes"
)

// auth

func (c *Controller) AddUserGroupAuth(uclaim *claim.Session, gslug string, data *entities.UserGroupAuth) error {
	if !c.HasScope(uclaim, "user") {
		return scopes.ErrNoAdminUserScope
	}

	data.TenantId = uclaim.TenantId
	data.UserGroup = gslug
	return c.coredb.AddUserGroupAuth(data)
}

func (c *Controller) UpdateUserGroupAuth(uclaim *claim.Session, gslug string, id int64, data map[string]any) error {
	if !c.HasScope(uclaim, "user") {
		return scopes.ErrNoAdminUserScope
	}

	return c.coredb.UpdateUserGroupAuth(uclaim.TenantId, gslug, id, data)
}

func (c *Controller) ListUserGroupAuth(uclaim *claim.Session, gslug string) ([]*entities.UserGroupAuth, error) {
	if !c.HasScope(uclaim, "user") {
		return nil, scopes.ErrNoAdminUserScope
	}

	return c.coredb.ListUserGroupAuth(uclaim.TenantId, gslug)
}

func (c *Controller) GetUserGroupAuth(uclaim *claim.Session, gslug string, id int64) (*entities.UserGroupAuth, error) {
	if !c.HasScope(uclaim, "user") {
		return nil, scopes.ErrNoAdminUserScope
	}

	return c.coredb.GetUserGroupAuth(uclaim.TenantId, gslug, id)
}

func (c *Controller) RemoveUserGroupAuth(uclaim *claim.Session, gslug string, id int64) error {
	if !c.HasScope(uclaim, "user") {
		return scopes.ErrNoAdminUserScope
	}

	return c.coredb.RemoveUserGroupAuth(uclaim.TenantId, gslug, id)
}

// data

func (c *Controller) AddUserGroupData(uclaim *claim.Session, gslug string, data *entities.UserGroupData) error {
	if !c.HasScope(uclaim, "user") {
		return scopes.ErrNoAdminUserScope
	}

	data.TenantId = uclaim.TenantId
	data.UserGroup = gslug
	return c.coredb.AddUserGroupData(data)
}

func (c *Controller) UpdateUserGroupData(uclaim *claim.Session, gslug string, id int64, data map[string]any) error {
	if !c.HasScope(uclaim, "user") {
		return scopes.ErrNoAdminUserScope
	}

	return c.coredb.UpdateUserGroupData(uclaim.TenantId, gslug, id, data)
}

func (c *Controller) ListUserGroupData(uclaim *claim.Session, gslug string) ([]*entities.UserGroupData, error) {
	if !c.HasScope(uclaim, "user") {
		return nil, scopes.ErrNoAdminUserScope
	}

	return c.coredb.ListUserGroupData(uclaim.TenantId, gslug)
}

func (c *Controller) GetUserGroupData(uclaim *claim.Session, gslug string, id int64) (*entities.UserGroupData, error) {
	if !c.HasScope(uclaim, "user") {
		return nil, scopes.ErrNoAdminUserScope
	}

	return c.coredb.GetUserGroupData(uclaim.TenantId, gslug, id)
}

func (c *Controller) RemoveUserGroupData(uclaim *claim.Session, gslug string, id int64) error {
	if !c.HasScope(uclaim, "user") {
		return scopes.ErrNoAdminUserScope
	}

	return c.coredb.RemoveUserGroupData(uclaim.TenantId, gslug, id)
}
