package pacman

import (
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
)

func (p *PacMan) RepoSources(tenantid string) (map[int64]string, error) {
	repos, err := p.syncer.RepoList(tenantid)
	if err != nil {
		return nil, err
	}

	resp := make(map[int64]string, len(repos))
	for _, repo := range repos {
		resp[repo.Id] = repo.Name
	}

	return resp, nil
}

func (p *PacMan) RepoSourceList(tenantid, group string, source int64, tags ...string) ([]entities.BPrint, error) {
	repo := p.getRepoSource(tenantid, source)
	if repo == nil {
		return nil, easyerr.NotFound()
	}

	return repo.Query(tenantid, &repox.RepoQuery{
		Group: group,
		Tags:  tags,
		Page:  0,
	})
}

func (p *PacMan) RepoSourceGet(tenantid, group, slug string, source int64) (*entities.BPrint, error) {
	repo := p.getRepoSource(tenantid, source)
	if repo == nil {
		return nil, easyerr.NotFound()
	}

	return repo.GetItem(tenantid, group, slug)
}

func (p *PacMan) RepoSourceGetBlob(tenantid, group, slug string, source int64, file string) ([]byte, error) {
	repo := p.getRepoSource(tenantid, source)
	if repo == nil {
		return nil, easyerr.NotFound()
	}
	return repo.GetFile(tenantid, group, slug, file)
}

// private

func (p *PacMan) getRepoSource(tenantId string, source int64) repox.Repository {

	p.activeRepoMutex.Lock()
	trepos, ok := p.activeRepo[tenantId]
	if ok {
		repo, ok := trepos[source]
		if ok {
			p.activeRepoMutex.Unlock()
			return repo
		}
	} else {
		p.activeRepo[tenantId] = make(map[int64]repox.Repository)
	}
	p.activeRepoMutex.Unlock()

	rrepo, err := p.syncer.RepoGet(tenantId, source)
	if err != nil {
		return nil
	}

	repoBuilder, ok := p.repoBuilders[rrepo.Provider]
	if !ok {
		return nil
	}

	newrepo, err := repoBuilder(&repox.BuilderOptions{
		TenantId: tenantId,
		BaseURL:  rrepo.URL,
	})
	if err != nil {
		return nil
	}

	p.activeRepoMutex.Lock()
	trepos = p.activeRepo[tenantId]
	oldrepo, ok := trepos[source]
	// another thread could have
	// inserted after we last checked
	if ok {
		p.activeRepoMutex.Unlock()
		return oldrepo
	}

	trepos[source] = newrepo

	p.activeRepoMutex.Unlock()

	return newrepo

}
