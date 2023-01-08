package admin

import (
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
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

func (c *Controller) BprintInstance(uclaim *claim.Session, opts *repox.InstanceOptions) (any, error) {
	opts.UserSession = uclaim

	// fixme => use new instance thing here
	// return c.pacman.Instance(uclaim.TenentId, opts)

	return nil, nil
}
