package repox

import (
	"io"
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

type BPrint struct {
	Name        string   `json:"name,omitempty"`
	Slug        string   `json:"slug,omitempty"`
	Type        string   `json:"type,omitempty"`
	SubType     string   `json:"sub_type,omitempty"`
	Description string   `json:"description,omitempty"`
	Icon        string   `json:"icon,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type Repository interface {
	Name() string
	Query(tenantId string, opts *RepoQuery) ([]BPrint, error)

	Get(tenantid, slug string) (*BPrint, error)
	GetZip(tenantid, slug, version string) (io.ReadCloser, error)
}
