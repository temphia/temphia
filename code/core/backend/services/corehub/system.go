package corehub

import "github.com/temphia/temphia/code/core/backend/xtypes/models/entities"

func (c *CoreHub) AddSystemEvent(data *entities.SystemEvent) error {
	return c.coredb.AddSystemEvent(data)
}

func (c *CoreHub) RemoveSystemEvent(id int64) error {
	return c.coredb.RemoveSystemEvent(id)
}

func (c *CoreHub) ListSystemEvent(tenantId, etype string, last int64) ([]*entities.SystemEvent, error) {
	return c.coredb.ListSystemEvent(tenantId, etype, last)
}

func (c *CoreHub) AddSystemKV(tenantId, data *entities.SystemKV) error {
	return c.coredb.AddSystemKV(tenantId, data)
}

func (c *CoreHub) UpdateSystemKV(tenantId, key, ktype string, data map[string]any) error {
	return c.coredb.UpdateSystemKV(tenantId, key, ktype, data)
}

func (c *CoreHub) GetSystemKV(tenantId, key, ktype string) (*entities.SystemKV, error) {
	return c.coredb.GetSystemKV(tenantId, key, ktype)
}

func (c *CoreHub) RemoveSystemKV(tenantId, key, ktype string) error {
	return c.coredb.RemoveSystemKV(tenantId, key, ktype)
}

func (c *CoreHub) ListSystemKV(tenantId, ktype, prefix string) ([]*entities.SystemKV, error) {
	return c.coredb.ListSystemKV(tenantId, ktype, prefix)
}
