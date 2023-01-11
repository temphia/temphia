package corehub

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

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
