package corehub

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *CoreHub) AddTargetHook(data *entities.TargetHook) (int64, error) {

	id, err := c.coredb.AddTargetHook(data)
	if err != nil {
		return 0, err
	}

	data.Id = id

	c.stateHub.OnTargetHookChange(data.TenantId, 0, data)
	return id, nil
}

func (c *CoreHub) UpdateTargetHook(tenantId, ttype string, id int64, data map[string]any) error {
	err := c.coredb.UpdateTargetHook(tenantId, ttype, id, data)
	if err != nil {
		return err
	}

	c.stateHub.OnTargetHookChange(tenantId, id, data)

	return nil
}

func (c *CoreHub) ListTargetHook(tenantId string, cond map[string]any) ([]*entities.TargetHook, error) {
	return c.coredb.ListTargetHook(tenantId, cond)
}

func (c *CoreHub) ListTargetHookByType(tenantId, ttype, target string) ([]*entities.TargetHook, error) {
	return c.coredb.ListTargetHookByType(tenantId, ttype, target)
}

func (c *CoreHub) GetTargetHook(tenantId, ttype string, id int64) (*entities.TargetHook, error) {
	return c.coredb.GetTargetHook(tenantId, ttype, id)
}

func (c *CoreHub) RemoveTargetHook(tenantId, ttype string, id int64) error {

	err := c.coredb.RemoveTargetHook(tenantId, ttype, id)
	if err != nil {
		return err
	}

	c.stateHub.OnTargetHookChange(tenantId, id, nil)
	return nil

}

func (c *CoreHub) AddTargetApp(data *entities.TargetApp) (int64, error) {
	id, err := c.coredb.AddTargetApp(data)
	if err != nil {
		return 0, err
	}

	data.Id = id
	c.stateHub.OnTargetAppChange(data.TenantId, 0, data)

	return 0, nil
}

func (c *CoreHub) UpdateTargetApp(tenantId, ttype string, id int64, data map[string]any) error {
	err := c.coredb.UpdateTargetApp(tenantId, ttype, id, data)
	if err != nil {
		return err
	}

	c.stateHub.OnTargetAppChange(tenantId, id, data)
	return nil
}

func (c *CoreHub) ListTargetApp(tenantId string, cond map[string]any) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetApp(tenantId, cond)
}

func (c *CoreHub) ListTargetAppByType(tenantId, ttype, target string) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetAppByType(tenantId, ttype, target)
}

func (c *CoreHub) GetTargetApp(tenantId, ttype string, id int64) (*entities.TargetApp, error) {
	return c.coredb.GetTargetApp(tenantId, ttype, id)
}

func (c *CoreHub) RemoveTargetApp(tenantId, ttype string, id int64) error {

	err := c.coredb.RemoveTargetApp(tenantId, ttype, id)
	if err != nil {
		return err
	}

	c.stateHub.OnTargetAppChange(tenantId, id, nil)
	return nil

}

func (c *CoreHub) ListTargetAppByUgroup(tenantId, ugroup string) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetAppByUgroup(tenantId, ugroup)
}
