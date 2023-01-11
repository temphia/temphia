package admin

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

func (c *Controller) BprintList(uclaim *claim.Session, group string) ([]*entities.BPrint, error) {

	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintList(uclaim.TenentId, group)
}

func (c *Controller) BprintCreate(uclaim *claim.Session, bp *entities.BPrint) (string, error) {
	if !uclaim.IsSuperAdmin() {
		return "", easyerr.NotImpl()
	}

	return c.pacman.BprintCreate(uclaim.TenentId, bp)
}

func (c *Controller) BprintUpdate(uclaim *claim.Session, bp *entities.BPrint) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintUpdate(uclaim.TenentId, bp)
}

func (c *Controller) BprintGet(uclaim *claim.Session, bid string) (*entities.BPrint, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintGet(uclaim.TenentId, bid)
}

func (c *Controller) BprintRemove(uclaim *claim.Session, bid string) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintRemove(uclaim.TenentId, bid)
}

func (c *Controller) BprintListBlobs(uclaim *claim.Session, bid string) (any, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintListBlobs(uclaim.TenentId, bid)
}

func (c *Controller) BprintNewBlob(uclaim *claim.Session, bid, file string, payload []byte) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintNewBlob(uclaim.TenentId, bid, file, payload)
}

func (c *Controller) BprintUpdateBlob(uclaim *claim.Session, bid, file string, payload []byte) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}
	return c.pacman.BprintUpdateBlob(uclaim.TenentId, bid, file, payload)
}

func (c *Controller) BprintGetBlob(uclaim *claim.Session, bid, file string) ([]byte, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.pacman.BprintGetBlob(uclaim.TenentId, bid, file)
}

func (c *Controller) BprintDeleteBlob(uclaim *claim.Session, bid, file string) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.pacman.BprintDeleteBlob(uclaim.TenentId, bid, file)
}

// repo

func (c *Controller) BprintImport(uclaim *claim.Session, opts *repox.RepoImportOpts) (string, error) {
	return c.pacman.RepoSourceImport(uclaim.TenentId, opts)
}

type InstanceOptions struct {
	RepoId         int64           `json:"repo_id,omitempty"`
	InstancerType  string          `json:"instancer_type,omitempty"`
	File           string          `json:"file,omitempty"`
	UserConfigData json.RawMessage `json:"data,omitempty"`
	Auto           bool            `json:"auto,omitempty"`
}

func (c *Controller) BprintInstance(uclaim *claim.Session, bid string, opts *InstanceOptions) (any, error) {

	instancer := c.pacman.GetInstanceHub()

	fopt := repox.InstanceOptions{
		BprintId:       bid,
		RepoId:         opts.RepoId,
		InstancerType:  opts.InstancerType,
		File:           opts.File,
		UserConfigData: opts.UserConfigData,
		Auto:           opts.Auto,
		UserSession:    uclaim,
	}

	switch opts.InstancerType {
	case xbprint.TypeBundle:
		if opts.Auto {
			return instancer.AutomaticBundle(fopt)
		}
		return instancer.ManualBundleItem(fopt)

	default:
		if opts.Auto {
			return instancer.AutomaticSingle(fopt)
		}
		return instancer.ManualSingle(fopt)
	}
}
