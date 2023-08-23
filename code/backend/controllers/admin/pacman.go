package admin

import (
	"encoding/json"
	"mime/multipart"

	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/scopes"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xinstancer"
)

func (c *Controller) BprintList(uclaim *claim.Session, group string) ([]*entities.BPrint, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.pacman.BprintList(uclaim.TenantId, group)
}

func (c *Controller) BprintCreate(uclaim *claim.Session, bp *entities.BPrint) (string, error) {
	if !c.HasScope(uclaim, "engine") {
		return "", scopes.ErrNoAdminEngineScope
	}

	bp.TenantID = uclaim.TenantId
	return c.pacman.BprintCreate(uclaim.TenantId, bp)
}

func (c *Controller) BprintUpdate(uclaim *claim.Session, id string, data map[string]any) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.pacman.BprintUpdate(uclaim.TenantId, id, data)
}

func (c *Controller) BprintGet(uclaim *claim.Session, bid string) (*entities.BPrint, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.pacman.BprintGet(uclaim.TenantId, bid)
}

func (c *Controller) BprintRemove(uclaim *claim.Session, bid string) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	return c.pacman.BprintRemove(uclaim.TenantId, bid)
}

func (c *Controller) BprintListBlobs(uclaim *claim.Session, bid string) (any, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	return c.pacman.BprintListBlobs(uclaim.TenantId, bid)
}

func (c *Controller) BprintNewBlob(uclaim *claim.Session, bid, file string, payload []byte) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	bstore := c.pacman.GetBprintFileStore()

	return bstore.NewBlob(uclaim.TenantId, bid, file, "", payload)
}

func (c *Controller) BprintUpdateBlob(uclaim *claim.Session, bid, file string, payload []byte) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	bstore := c.pacman.GetBprintFileStore()

	return bstore.UpdateBlob(uclaim.TenantId, bid, "", file, payload)
}

func (c *Controller) BprintGetBlob(uclaim *claim.Session, bid, file string) ([]byte, error) {
	if !c.HasScope(uclaim, "engine") {
		return nil, scopes.ErrNoAdminEngineScope
	}

	bstore := c.pacman.GetBprintFileStore()

	return bstore.GetBlob(uclaim.TenantId, bid, "", file)
}

func (c *Controller) BprintDeleteBlob(uclaim *claim.Session, bid, file string) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	bstore := c.pacman.GetBprintFileStore()

	return bstore.DeleteBlob(uclaim.TenantId, bid, "", file)
}

func (c *Controller) BprintCreateFromZip(uclaim *claim.Session, form *multipart.Form) error {
	if !c.HasScope(uclaim, "engine") {
		return scopes.ErrNoAdminEngineScope
	}

	zfile := form.File["file"][0]
	zipfile, err := zfile.Open()
	if err != nil {
		return err
	}

	defer zipfile.Close()

	_, err = c.pacman.BprintCreateFromZip(uclaim.TenantId, zipfile)
	if err != nil {
		return err
	}

	return nil
}

// repo

func (c *Controller) BprintImport(uclaim *claim.Session, opts *xpacman.RepoImportOpts) (string, error) {
	if !c.HasScope(uclaim, "engine") {
		return "", scopes.ErrNoAdminEngineScope
	}

	return c.pacman.RepoImport(uclaim.TenantId, opts)
}

type InstanceOptions struct {
	InstancerType  string          `json:"instancer_type,omitempty"`
	File           string          `json:"file,omitempty"`
	UserConfigData json.RawMessage `json:"data,omitempty"`
	Auto           bool            `json:"auto,omitempty"`
}

func (c *Controller) BprintInstance(uclaim *claim.Session, bid string, opts *InstanceOptions) (any, error) {
	// if !c.HasScope(uclaim, "engine") {
	// 	return nil, scopes.ErrNoAdminEngineScope
	// }

	return c.pacman.GetInstancer().Instance(xinstancer.Options{
		BprintId:     bid,
		UserSession:  uclaim.AsUserCtx(),
		PlugId:       "",
		NextBprintId: "",
		TenantId:     uclaim.TenantId,
		InstancedIds: make(map[string]string),
	})

}
