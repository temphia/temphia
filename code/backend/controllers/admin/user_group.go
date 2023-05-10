package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/scopes"
)

func (c *Controller) AddUserGroup(uclaim *claim.Session, ugroup *entities.UserGroup) error {
	if !c.HasScope(uclaim, "user") {
		return scopes.ErrNoAdminUserScope
	}

	return c.coredb.AddUserGroup(ugroup)
}

func (c *Controller) ListUserGroup(uclaim *claim.Session) ([]*entities.UserGroup, error) {
	if !c.HasScope(uclaim, "user") {
		return nil, scopes.ErrNoAdminUserScope
	}

	return c.coredb.ListUserGroups(uclaim.TenantId)
}

func (c *Controller) GetUserGroup(uclaim *claim.Session, ugroup string) (*entities.UserGroup, error) {
	if !c.HasScope(uclaim, "user") {
		return nil, scopes.ErrNoAdminUserScope
	}

	return c.coredb.GetUserGroup(uclaim.TenantId, ugroup)
}

func (c *Controller) UpdateUserGroup(uclaim *claim.Session, id string, data map[string]any) error {
	if !c.HasScope(uclaim, "user") {
		return scopes.ErrNoAdminUserScope
	}

	return c.coredb.UpdateUserGroup(uclaim.TenantId, id, data)
}

func (c *Controller) RemoveUserGroup(uclaim *claim.Session, ugroup string) error {
	if !c.HasScope(uclaim, "user") {
		return scopes.ErrNoAdminUserScope
	}

	return c.coredb.RemoveUserGroup(uclaim.TenantId, ugroup)
}
