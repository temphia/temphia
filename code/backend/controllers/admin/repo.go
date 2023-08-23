package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/models/scopes"
)

func (c *Controller) RepoNew(uclaim *claim.Session, data *entities.Repo) error {
	if !c.HasScope(uclaim, "repo") {
		return scopes.ErrNoAdminRepoScope
	}

	data.TenantId = uclaim.TenantId
	return c.coredb.RepoNew(uclaim.TenantId, data)
}

func (c *Controller) RepoUpdate(uclaim *claim.Session, id int64, data map[string]any) error {
	if !c.HasScope(uclaim, "repo") {
		return scopes.ErrNoAdminRepoScope
	}

	return c.coredb.RepoUpdate(uclaim.TenantId, id, data)
}

func (c *Controller) RepoGet(uclaim *claim.Session, id int64) (*entities.Repo, error) {
	if !c.HasScope(uclaim, "repo") {
		return nil, scopes.ErrNoAdminRepoScope
	}

	return c.coredb.RepoGet(uclaim.TenantId, id)
}

func (c *Controller) RepoDel(uclaim *claim.Session, id int64) error {
	if !c.HasScope(uclaim, "repo") {
		return scopes.ErrNoAdminRepoScope
	}

	return c.coredb.RepoDel(uclaim.TenantId, id)
}

func (c *Controller) RepoList(uclaim *claim.Session) ([]*entities.Repo, error) {
	if !c.HasScope(uclaim, "repo") {
		return nil, scopes.ErrNoAdminRepoScope
	}

	return c.coredb.RepoList(uclaim.TenantId)
}

func (c *Controller) RepoSources(uclaim *claim.Session) (map[int64]string, error) {
	if !c.HasScope(uclaim, "repo") {
		return nil, scopes.ErrNoAdminRepoScope
	}

	return c.pacman.RepoSources(uclaim.TenantId)
}
