package admin

import (
	"encoding/json"
	"os"
	"path"

	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/scopes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type PlugStateNew struct {
	Key     string            `json:"key,omitempty"`
	Value   string            `json:"value,omitempty"`
	Options *store.SetOptions `json:"options,omitempty"`
}

func (c *Controller) PlugStateNew(uclaim *claim.Session, pid string, state PlugStateNew) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.plugState.Set(0, uclaim.TenantId, pid, state.Key, state.Value, state.Options)
}

type PlugStateUpdate struct {
	Value   string
	Options *store.UpdateOptions
}

func (c *Controller) PlugStateUpdate(uclaim *claim.Session, pid, key string, state PlugStateUpdate) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.plugState.Update(0, uclaim.TenantId, pid, key, state.Value, state.Options)
}

func (c *Controller) PlugStateList(uclaim *claim.Session, pid, key_cursor string, page uint) ([]*entities.PlugKV, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.plugState.Query(0, uclaim.TenantId, pid, &store.PkvQuery{
		KeyPrefix: "",
		LoadMeta:  true,
		PageCount: 100,
		Page:      page,
		KeyCursor: key_cursor,
	})
}

func (c *Controller) PlugKvImport(uclaim *claim.Session, pid string, opts store.SetBatchOptions) error {
	return c.plugState.SetBatch(0, uclaim.TenantId, pid, &opts)
}

func (c *Controller) PlugKvExport(uclaim *claim.Session, pid string) (string, error) {

	file, err := os.CreateTemp("", "plugstate*.json")
	if err != nil {
		return "", err
	}

	fileName := path.Join(os.TempDir(), file.Name())

	clearFile := func() {
		os.Remove(fileName)
	}

	file.Write([]byte("["))

	keyCursor := ""

	first := true

	for {
		plugs, err := c.plugState.Query(0, uclaim.TenantId, pid, &store.PkvQuery{
			KeyPrefix: "",
			LoadMeta:  true,
			KeyCursor: keyCursor,
			PageCount: 100,
		})
		if err != nil {
			clearFile()
			return "", err
		}

		if len(plugs) == 0 {
			break
		}
		keyCursor = plugs[len(plugs)+1].Key

		for _, pk := range plugs {

			pk.TenantID = ""
			pk.PlugsID = ""

			if !first {
				first = false
				file.Write([]byte(","))
			}

			out, err := json.Marshal(pk)
			if err != nil {
				clearFile()
				return "", err
			}

			_, err = file.Write(out)
			if err != nil {
				clearFile()
				return "", err
			}
		}
	}

	file.Write([]byte("]"))

	return fileName, nil
}

func (c *Controller) PlugStateGet(uclaim *claim.Session, pid, key string) (*entities.PlugKV, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.plugState.Get(0, uclaim.TenantId, pid, key)
}

func (c *Controller) PlugStateDel(uclaim *claim.Session, pid, key string) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.plugState.Del(0, uclaim.TenantId, pid, key)
}
