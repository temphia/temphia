package corehub

import "github.com/temphia/temphia/code/core/backend/xtypes/models/entities"

func (c *CoreHub) AddTargetHook(data *entities.TargetHook) error {
	return c.coredb.AddTargetHook(data)
}

func (c *CoreHub) UpdateTargetHook(tenantId, targetType string, id int64, data map[string]any) error {
	return c.coredb.UpdateTargetHook(tenantId, targetType, id, data)
}

func (c *CoreHub) ListTargetHook(tenantId string) ([]*entities.TargetHook, error) {
	return c.coredb.ListTargetHook(tenantId)
}

func (c *CoreHub) ListTargetHookByType(tenantId, targetType string) ([]*entities.TargetHook, error) {
	return c.coredb.ListTargetHookByType(tenantId, targetType)
}

func (c *CoreHub) GetTargetHook(tenantId, targetType string, id int64) (*entities.TargetHook, error) {
	return c.coredb.GetTargetHook(tenantId, targetType, id)
}

func (c *CoreHub) RemoveTargetHook(tenantId, targetType string, id int64) error {
	return c.coredb.RemoveTargetHook(tenantId, targetType, id)
}

func (c *CoreHub) AddTargetApp(data *entities.TargetApp) error {
	return c.coredb.AddTargetApp(data)
}

func (c *CoreHub) UpdateTargetApp(tenantId, targetType string, id int64, data map[string]any) error {
	return c.coredb.UpdateTargetApp(tenantId, targetType, id, data)
}

func (c *CoreHub) ListTargetApp(tenantId string) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetApp(tenantId)
}

func (c *CoreHub) ListTargetAppByType(tenantId, targetType string) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetAppByType(tenantId, targetType)
}

func (c *CoreHub) GetTargetApp(tenantId, targetType string, id int64) (*entities.TargetApp, error) {
	return c.coredb.GetTargetApp(tenantId, targetType, id)
}

func (c *CoreHub) RemoveTargetApp(tenantId, targetType string, id int64) error {
	return c.coredb.RemoveTargetApp(tenantId, targetType, id)
}

func (c *CoreHub) ListTargetAppByUgroup(tenantId, ugroup string) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetAppByUgroup(tenantId, ugroup)
}
