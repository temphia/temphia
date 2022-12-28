package repox

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

type BuilderOptions struct {
	TenantId  string
	BasePath  string
	ExtraMeta map[string]string
}

type Builder func(opts *BuilderOptions) (Repository, error)

type RepoQuery struct {
	Group string
	Tags  []string
	Page  int64
}

type Repository interface {
	Name() string
	Query(tenantId string, opts *RepoQuery) ([]entities.BPrint, error)
	GetItem(tenantid, group, slug string) (*entities.BPrint, error)
	GetFile(tenantid, group, slug, file string) ([]byte, error)
	GetFileURL(tenantid, group, slug, file string) (string, error)
}
