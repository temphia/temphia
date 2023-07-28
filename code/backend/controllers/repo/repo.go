package repo

import (
	"io"

	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
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

func (c *Controller) RepoSourceList(uclaim *claim.Session, group string, source int64, tags ...string) ([]repox.BPrint, error) {
	return c.pacman.RepoSourceList(uclaim.TenantId, group, source, tags...)
}

func (c *Controller) RepoSourceGet(uclaim *claim.Session, group, slug string, source int64) (*repox.BPrint, error) {
	return c.pacman.RepoSourceGet(uclaim.TenantId, slug, source)
}

func (c *Controller) RepoSourceGetZip(uclaim *claim.Session, group, slug string, version string, source int64) (io.ReadCloser, error) {
	return c.pacman.RepoSourceGetZip(uclaim.TenantId, source, slug, version)
}

func (c *Controller) RepoSourceImport(uclaim *claim.Session, data *repox.RepoImportOpts) (string, error) {
	return c.pacman.RepoSourceImport(uclaim.TenantId, data)
}
