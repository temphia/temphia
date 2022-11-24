package admin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type PlugStateNew struct {
	Key     string
	Value   string
	Options *store.SetOptions
}

func (c *Controller) PlugStateNew(uclaim *claim.Session, pid string, state PlugStateNew) error {
	return c.plugState.Set(0, uclaim.TenentId, pid, state.Key, state.Value, state.Options)
}

type PlugStateUpdate struct {
	Value   string
	Options *store.UpdateOptions
}

func (c *Controller) PlugStateUpdate(uclaim *claim.Session, pid, key string, state PlugStateUpdate) error {
	return c.plugState.Update(0, uclaim.TenentId, pid, key, state.Value, state.Options)
}

func (c *Controller) PlugStateList(uclaim *claim.Session, pid string) ([]*entities.PlugKV, error) {
	// fixme => implement pagination

	return c.plugState.Query(0, uclaim.TenentId, pid, &store.PkvQuery{
		KeyPrefix: "",
		LoadMeta:  true,
		PageCount: 100,
		Page:      0,
	})
}

func (c *Controller) PlugStateGet(uclaim *claim.Session, pid, key string) (*entities.PlugKV, error) {
	return c.plugState.Get(0, uclaim.TenentId, pid, key)
}

func (c *Controller) PlugStateDel(uclaim *claim.Session, pid, key string) error {
	return c.plugState.Del(0, uclaim.TenentId, pid, key)
}
