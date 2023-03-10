package corehub

import "github.com/temphia/temphia/code/core/backend/xtypes/models/entities"

func (c *CoreHub) AddTargetHook(data *entities.TargetHook) error {
	return c.coredb.AddTargetHook(data)
}

func (c *CoreHub) UpdateTargetHook(tenantId, ttype string, id int64, data map[string]any) error {
	return c.coredb.UpdateTargetHook(tenantId, ttype, id, data)
}

func (c *CoreHub) ListTargetHook(tenantId string) ([]*entities.TargetHook, error) {
	return c.coredb.ListTargetHook(tenantId)
}

func (c *CoreHub) ListTargetHookByType(tenantId, ttype, target string) ([]*entities.TargetHook, error) {
	return c.coredb.ListTargetHookByType(tenantId, ttype, target)
}

func (c *CoreHub) GetTargetHook(tenantId, ttype string, id int64) (*entities.TargetHook, error) {
	return c.coredb.GetTargetHook(tenantId, ttype, id)
}

func (c *CoreHub) RemoveTargetHook(tenantId, ttype string, id int64) error {
	return c.coredb.RemoveTargetHook(tenantId, ttype, id)
}

func (c *CoreHub) AddTargetApp(data *entities.TargetApp) error {
	return c.coredb.AddTargetApp(data)
}

func (c *CoreHub) UpdateTargetApp(tenantId, ttype string, id int64, data map[string]any) error {
	return c.coredb.UpdateTargetApp(tenantId, ttype, id, data)
}

func (c *CoreHub) ListTargetApp(tenantId string) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetApp(tenantId)
}

func (c *CoreHub) ListTargetAppByType(tenantId, ttype, target string) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetAppByType(tenantId, ttype, target)
}

func (c *CoreHub) GetTargetApp(tenantId, ttype string, id int64) (*entities.TargetApp, error) {
	return c.coredb.GetTargetApp(tenantId, ttype, id)
}

func (c *CoreHub) RemoveTargetApp(tenantId, ttype string, id int64) error {
	return c.coredb.RemoveTargetApp(tenantId, ttype, id)
}

func (c *CoreHub) ListTargetAppByUgroup(tenantId, ugroup string) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetAppByUgroup(tenantId, ugroup)
}
