package admin

import (
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func (c *Controller) GetPlugState(aclaim *claim.PlugState, key string) (*entities.PlugKV, error) {
	if !strings.HasPrefix(key, aclaim.KeyPrefix) {
		return nil, easyerr.NotAuthorized()
	}

	return c.plugState.Get(0, aclaim.TenantId, aclaim.PlugId, key)
}

type AddPlugStateOptions struct {
	Key     string           `json:"key,omitempty"`
	Value   string           `json:"value,omitempty"`
	Options store.SetOptions `json:"options,omitempty"`
}

func (c *Controller) AddPlugState(aclaim *claim.PlugState, opts AddPlugStateOptions) error {
	if !strings.HasPrefix(opts.Key, aclaim.KeyPrefix) {
		return easyerr.NotAuthorized()
	}

	return c.plugState.Set(0, aclaim.TenantId, aclaim.PlugId, opts.Key, opts.Value, &opts.Options)
}

type UpdatePlugStateOptions struct {
	Key     string              `json:"key,omitempty"`
	Value   string              `json:"value,omitempty"`
	Options store.UpdateOptions `json:"options,omitempty"`
}

func (c *Controller) UpdatePlugState(aclaim *claim.PlugState, opts UpdatePlugStateOptions) error {
	if !strings.HasPrefix(opts.Key, aclaim.KeyPrefix) {
		return easyerr.NotAuthorized()
	}

	return c.plugState.Update(0, aclaim.TenantId, aclaim.PlugId, opts.Key, opts.Value, &opts.Options)
}

func (c *Controller) DeletePlugState(aclaim *claim.PlugState, key string) error {
	if !strings.HasPrefix(key, aclaim.KeyPrefix) {
		return easyerr.NotAuthorized()
	}

	return c.plugState.Del(0, aclaim.TenantId, aclaim.PlugId, key)
}

func (c *Controller) ListPlugState(aclaim *claim.PlugState, opts *store.PkvQuery) ([]*entities.PlugKV, error) {

	if !strings.HasPrefix(opts.KeyCursor, aclaim.KeyPrefix) {
		return nil, easyerr.NotAuthorized()
	}

	return c.plugState.Query(0, aclaim.TenantId, aclaim.PlugId, opts)

}
