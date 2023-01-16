package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type PlugStateNew struct {
	Key     string            `json:"key,omitempty"`
	Value   string            `json:"value,omitempty"`
	Options *store.SetOptions `json:"options,omitempty"`
}

func (c *Controller) PlugStateNew(uclaim *claim.Session, pid string, state PlugStateNew) error {
	return c.plugState.Set(0, uclaim.TenantId, pid, state.Key, state.Value, state.Options)
}

type PlugStateUpdate struct {
	Value   string
	Options *store.UpdateOptions
}

func (c *Controller) PlugStateUpdate(uclaim *claim.Session, pid, key string, state PlugStateUpdate) error {
	return c.plugState.Update(0, uclaim.TenantId, pid, key, state.Value, state.Options)
}

func (c *Controller) PlugStateList(uclaim *claim.Session, pid, key_cursor string, page uint) ([]*entities.PlugKV, error) {

	return c.plugState.Query(0, uclaim.TenantId, pid, &store.PkvQuery{
		KeyPrefix: "",
		LoadMeta:  true,
		PageCount: 100,
		Page:      page,
		KeyCursor: key_cursor,
	})
}

func (c *Controller) PlugStateGet(uclaim *claim.Session, pid, key string) (*entities.PlugKV, error) {
	return c.plugState.Get(0, uclaim.TenantId, pid, key)
}

func (c *Controller) PlugStateDel(uclaim *claim.Session, pid, key string) error {
	return c.plugState.Del(0, uclaim.TenantId, pid, key)
}
