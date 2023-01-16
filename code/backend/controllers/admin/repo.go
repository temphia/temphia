package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *Controller) RepoNew(uclaim *claim.Session, data *entities.Repo) error {
	data.TenantId = uclaim.TenantId
	return c.coredb.RepoNew(uclaim.TenantId, data)
}

func (c *Controller) RepoUpdate(uclaim *claim.Session, id int64, data map[string]any) error {
	return c.coredb.RepoUpdate(uclaim.TenantId, id, data)
}

func (c *Controller) RepoGet(uclaim *claim.Session, id int64) (*entities.Repo, error) {
	return c.coredb.RepoGet(uclaim.TenantId, id)
}

func (c *Controller) RepoDel(uclaim *claim.Session, id int64) error {
	return c.coredb.RepoDel(uclaim.TenantId, id)
}

func (c *Controller) RepoList(uclaim *claim.Session) ([]*entities.Repo, error) {
	return c.coredb.RepoList(uclaim.TenantId)
}

func (c *Controller) RepoSources(uclaim *claim.Session) (map[int64]string, error) {
	return c.pacman.RepoSources(uclaim.TenantId)
}
