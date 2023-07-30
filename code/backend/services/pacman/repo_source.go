package pacman

import (
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
)

func (p *PacMan) RepoSources(tenantid string) (map[int64]string, error) {
	repos, err := p.corehub.RepoList(tenantid)
	if err != nil {
		return nil, err
	}

	resp := make(map[int64]string, len(repos))
	for _, repo := range repos {
		resp[repo.Id] = repo.Name
	}

	return resp, nil
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

	rrepo, err := p.corehub.RepoGet(tenantId, source)
	if err != nil {
		return nil
	}

	repoBuilder, ok := p.repoBuilders[rrepo.Provider]
	if !ok {
		return nil
	}

	newrepo, err := repoBuilder(&repox.BuilderOptions{
		TenantId:  tenantId,
		BasePath:  rrepo.URL,
		ExtraMeta: rrepo.ExtraMeta,
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
