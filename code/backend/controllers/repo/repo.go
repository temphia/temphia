package repo

import (
	"io"

	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
)

type Controller struct {
	pacman xpacman.Pacman
}

func New(pacman xpacman.Pacman) *Controller {
	return &Controller{
		pacman: pacman,
	}
}

func (c *Controller) RepoSourceList(uclaim *claim.Session, group string, source int64, tags ...string) ([]xpacman.BPrint, error) {
	return c.pacman.RepoList(uclaim.TenantId, group, source, tags...)
}

func (c *Controller) RepoSourceGet(uclaim *claim.Session, group, slug string, source int64) (*xpacman.BPrint, error) {
	return c.pacman.RepoGet(uclaim.TenantId, slug, source)
}

func (c *Controller) RepoSourceGetZip(uclaim *claim.Session, group, slug string, version string, source int64) (io.ReadCloser, error) {
	return c.pacman.RepoGetZip(uclaim.TenantId, source, slug, version)
}

func (c *Controller) RepoSourceImport(uclaim *claim.Session, data *xpacman.RepoImportOpts) (string, error) {
	return c.pacman.RepoImport(uclaim.TenantId, data)
}
