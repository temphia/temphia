package repo

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
)

type Controller struct {
	pacman repox.Hub
}

func New(pacman repox.Hub) *Controller {
	return &Controller{
		pacman: pacman,
	}
}

func (c *Controller) RepoSourceList(uclaim *claim.Session, group string, source int64, tags ...string) ([]entities.BPrint, error) {
	return c.pacman.RepoSourceList(uclaim.TenentId, group, source, tags...)
}

func (c *Controller) RepoSourceGet(uclaim *claim.Session, group, slug string, source int64) (*entities.BPrint, error) {
	return c.pacman.RepoSourceGet(uclaim.TenentId, group, slug, source)
}

func (c *Controller) RepoSourceGetBlob(uclaim *claim.Session, group, slug string, source int64, file string) ([]byte, error) {
	return c.pacman.RepoSourceGetBlob(uclaim.TenentId, group, slug, source, file)
}

func (c *Controller) RepoSourceImport(uclaim *claim.Session, data *repox.RepoImportOpts) (string, error) {
	return c.pacman.RepoSourceImport(uclaim.TenentId, data)
}
