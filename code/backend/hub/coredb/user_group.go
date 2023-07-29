package corehub

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

func (c *CoreHub) AddUserGroup(ug *entities.UserGroup) error {
	err := c.coredb.AddUserGroup(ug)
	if err != nil {
		return err
	}

	c.stateHub.OnUserGroupChange(ug.TenantID, ug.Slug, ug)
	return nil
}

func (c *CoreHub) GetUserGroup(tenantId string, slug string) (*entities.UserGroup, error) {
	return c.coredb.GetUserGroup(tenantId, slug)
}

func (c *CoreHub) ListUserGroups(tenantId string) ([]*entities.UserGroup, error) {
	return c.coredb.ListUserGroups(tenantId)
}

func (c *CoreHub) UpdateUserGroup(tenantId, slug string, data map[string]any) error {
	err := c.coredb.UpdateUserGroup(tenantId, slug, data)
	if err != nil {
		return err
	}

	c.stateHub.OnUserGroupChange(tenantId, slug, data)
	return nil
}

func (c *CoreHub) RemoveUserGroup(tenantId string, ugslug string) error {
	err := c.coredb.RemoveUserGroup(tenantId, ugslug)
	if err != nil {
		return err
	}

	c.stateHub.OnUserGroupChange(tenantId, ugslug, nil)
	return nil
}

// extra

func (c *CoreHub) AddUserGroupAuth(data *entities.UserGroupAuth) error {
	return c.coredb.AddUserGroupAuth(data)
}

func (c *CoreHub) UpdateUserGroupAuth(tenantId string, gslug string, id int64, data map[string]any) error {
	return c.coredb.UpdateUserGroupAuth(tenantId, gslug, id, data)
}

func (c *CoreHub) ListUserGroupAuth(tenantId string, gslug string) ([]*entities.UserGroupAuth, error) {
	return c.coredb.ListUserGroupAuth(tenantId, gslug)
}

func (c *CoreHub) GetUserGroupAuth(tenantId string, gslug string, id int64) (*entities.UserGroupAuth, error) {
	return c.coredb.GetUserGroupAuth(tenantId, gslug, id)
}

func (c *CoreHub) RemoveUserGroupAuth(tenantId, gslug string, id int64) error {
	return c.coredb.RemoveUserGroupAuth(tenantId, gslug, id)
}

func (c *CoreHub) AddUserGroupData(data *entities.UserGroupData) error {
	return c.coredb.AddUserGroupData(data)
}

func (c *CoreHub) UpdateUserGroupData(tenantId string, gslug string, id int64, data map[string]any) error {
	return c.coredb.UpdateUserGroupData(tenantId, gslug, id, data)
}

func (c *CoreHub) ListUserGroupData(tenantId string, gslug string) ([]*entities.UserGroupData, error) {
	return c.coredb.ListUserGroupData(tenantId, gslug)
}

func (c *CoreHub) GetUserGroupData(tenantId string, gslug string, id int64) (*entities.UserGroupData, error) {
	return c.coredb.GetUserGroupData(tenantId, gslug, id)
}

func (c *CoreHub) RemoveUserGroupData(tenantId, gslug string, id int64) error {
	return c.coredb.RemoveUserGroupData(tenantId, gslug, id)
}
