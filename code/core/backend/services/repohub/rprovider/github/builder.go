package github

import (
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
)

func init() {
	registry.SetRepoBuilder("official", NewOfficial)
	registry.SetRepoBuilder("github", New)
}

func NewOfficial(opts *repox.BuilderOptions) (repox.Repository, error) {

	return New(&repox.BuilderOptions{
		TenantId: opts.TenantId,
		BasePath: "",
		ExtraMeta: map[string]string{
			"branch": "main",
			"user":   "temphia",
			"repo":   "repo",
		},
	})
}

func New(opts *repox.BuilderOptions) (repox.Repository, error) {

	gh := &Github{
		user:   opts.ExtraMeta["user"],
		repo:   opts.ExtraMeta["repo"],
		branch: opts.ExtraMeta["branch"],
		cache:  nil,
	}

	return gh, nil
}
