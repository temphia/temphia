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

	// fixme => also pass tag1, tag2, tag3

	kv, err := c.plugState.Get(0, aclaim.TenantId, aclaim.PlugId, key)
	if err != nil {
		return nil, err
	}

	if aclaim.StateTag1 != "" {
		if kv.Tag1 != aclaim.StateTag1 {
			return nil, easyerr.NotFound()
		}
	}

	if aclaim.StateTag2 != "" {
		if kv.Tag1 != aclaim.StateTag2 {
			return nil, easyerr.NotFound()
		}
	}

	if aclaim.StateTag3 != "" {
		if kv.Tag1 != aclaim.StateTag3 {
			return nil, easyerr.NotFound()
		}
	}

	return kv, nil

}

func (c *Controller) AddPlugState(aclaim *claim.PlugState, data entities.PlugKV) error {

	return c.plugState.Set(0, aclaim.TenantId, aclaim.PlugId, data.Key, data.Value, &store.SetOptions{
		Tag1: aclaim.StateTag1,
		Tag2: aclaim.StateTag2,
		Tag3: aclaim.StateTag3,
		TTL:  0,
	})
}

func (c *Controller) UpdatePlugState(aclaim *claim.PlugState, key, value string) error {
	return c.plugState.Update(0, aclaim.TenantId, aclaim.PlugId, key, value, &store.UpdateOptions{})
}

func (c *Controller) DeletePlugState(aclaim *claim.PlugState, key string) error {
	if !strings.HasPrefix(key, aclaim.KeyPrefix) {
		return easyerr.NotAuthorized()
	}

	return c.plugState.Del(0, aclaim.TenantId, aclaim.PlugId, key)
}

func (c *Controller) ListPlugState(aclaim *claim.PlugState, page, pcount int, keycursor string) ([]*entities.PlugKV, error) {

	kv := &store.PkvQuery{
		KeyPrefix: aclaim.KeyPrefix,
		LoadMeta:  true,
		Tag1s:     nil,
		Tag2s:     nil,
		Tag3s:     nil,
		PageCount: uint(pcount),
		Page:      uint(page),
		KeyCursor: keycursor,
	}
	if kv.Tag1s != nil {
		kv.Tag1s = []string{aclaim.StateTag1}
	}

	if kv.Tag2s != nil {
		kv.Tag1s = []string{aclaim.StateTag2}
	}

	if kv.Tag3s != nil {
		kv.Tag3s = []string{aclaim.StateTag3}
	}

	return c.plugState.Query(0, aclaim.TenantId, aclaim.PlugId, kv)

}
