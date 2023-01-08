package repox

import (
	"encoding/json"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

type Hub interface {
	Start() error

	RepoCore
	RepoBprintOps
}

type RepoImportOpts struct {
	Slug      string   `json:"slug,omitempty"`
	Group     string   `json:"group,omitempty"`
	Source    int64    `json:"source,omitempty"`
	SkipFiles []string `json:"skip_files,omitempty"`
	NewId     string   `json:"new_id,omitempty"`
}

type RepoCore interface {
	RepoSources(tenantid string) (map[int64]string, error)
	RepoSourceList(tenantid, group string, source int64, tags ...string) ([]entities.BPrint, error)
	RepoSourceGet(tenantid, group, slug string, source int64) (*entities.BPrint, error)
	RepoSourceGetBlob(tenantid, group, slug string, source int64, file string) ([]byte, error)
	RepoSourceImport(tenantid string, data *RepoImportOpts) (string, error)
}

type RepoBprintOps interface {
	BprintList(tenantid, group string) ([]*entities.BPrint, error)
	BprintCreate(tenantid string, bp *entities.BPrint) (string, error)
	BprintUpdate(tenantid string, bp *entities.BPrint) error
	BprintGet(tenantid, bid string) (*entities.BPrint, error)
	BprintRemove(tenantid, bid string) error
	BprintListBlobs(tenantid, bid string) (map[string]string, error)

	BprintNewBlob(tenantid, bid, file string, payload []byte) error
	BprintUpdateBlob(tenantid, bid, file string, payload []byte) error

	BprintGetBlob(tenantid, bid, file string) ([]byte, error)
	BprintDeleteBlob(tenantid, bid, file string) error
	Instance(tenantId string, opts *InstanceOptions) (any, error)

	ParseInstanceFile(tenantId, bid, file string, target any) error
}

type InstanceOptions struct {
	BprintId       string          `json:"bprint_id,omitempty"`
	RepoId         string          `json:"repo_id,omitempty"`
	InstancerType  string          `json:"instancer_type,omitempty"`
	File           string          `json:"file,omitempty"`
	UserConfigData json.RawMessage `json:"data,omitempty"`
	UserSession    *claim.Session  `json:"-"`
}

type InstanceOps interface {
	Instance(opt InstanceOptions) (any, error)
	InstanceBundle(opt InstanceOptions) (any, error)
	InstanceBundleItem(opt InstanceOptions, item string) (any, error)
}
