package admin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
)

func (c *Controller) RepoNew(uclaim *claim.Session, data *entities.Repo) error {
	data.TenantId = uclaim.TenentId
	return c.coredb.RepoNew(uclaim.TenentId, data)
}

func (c *Controller) RepoUpdate(uclaim *claim.Session, id int64, data map[string]any) error {
	return c.coredb.RepoUpdate(uclaim.TenentId, id, data)
}

func (c *Controller) RepoGet(uclaim *claim.Session, id int64) (*entities.Repo, error) {
	return c.coredb.RepoGet(uclaim.TenentId, id)
}

func (c *Controller) RepoDel(uclaim *claim.Session, id int64) error {
	return c.coredb.RepoDel(uclaim.TenentId, id)
}

func (c *Controller) RepoList(uclaim *claim.Session) ([]*entities.Repo, error) {
	return c.coredb.RepoList(uclaim.TenentId)
}

func (c *Controller) RepoSources(uclaim *claim.Session) (map[int64]string, error) {
	return c.pacman.RepoSources(uclaim.TenentId)
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

func (c *Controller) RepoSourceImport(uclaim *claim.Session, data *service.RepoImportOpts) (string, error) {
	return c.pacman.RepoSourceImport(uclaim.TenentId, data)
}
