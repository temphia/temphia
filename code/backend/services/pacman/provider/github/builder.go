package github

import (
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
)

func init() {
	registry.SetRepoBuilder("official", NewOfficial)
	registry.SetRepoBuilder("github", New)
}

func NewOfficial(opts *xpacman.BuilderOptions) (xpacman.Repository, error) {

	return New(&xpacman.BuilderOptions{
		TenantId: opts.TenantId,
		BasePath: "",
		ExtraMeta: map[string]string{
			"branch": "main",
			"user":   "temphia",
			"repo":   "repo",
		},
	})
}

func New(opts *xpacman.BuilderOptions) (xpacman.Repository, error) {

	gh := &Github{
		user:   opts.ExtraMeta["user"],
		repo:   opts.ExtraMeta["repo"],
		branch: opts.ExtraMeta["branch"],
		cache:  nil,
	}

	return gh, nil
}
