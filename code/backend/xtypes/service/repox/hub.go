package repox

import (
	"io"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type Pacman interface {
	Start() error

	GetBprintFileStore() BStore

	RepoCore
	RepoBprintOps

	GetInstancerHubV1() InstancerHubV1

	GetInstancerHubV2() InstancerHubV2
}

type RepoImportOpts struct {
	Slug    string `json:"slug,omitempty"`
	Group   string `json:"group,omitempty"`
	Source  int64  `json:"source,omitempty"`
	Version string `json:"version,omitempty"`
	NewId   string `json:"new_id,omitempty"`
}

type RepoCore interface {
	RepoSources(tenantid string) (map[int64]string, error)

	RepoSourceImport(tenantid string, opts *RepoImportOpts) (string, error)
	RepoSourceList(tenantid, group string, source int64, tags ...string) ([]BPrint, error)
	RepoSourceGet(tenantid, slug string, source int64) (*BPrint, error)
	RepoSourceGetZip(tenantid string, source int64, slug, version string) (io.ReadCloser, error)

	BprintCreateFromZip(tenantId string, rawreader io.ReadCloser) (string, error)
}

type RepoBprintOps interface {
	BprintList(tenantid, group string) ([]*entities.BPrint, error)
	BprintCreate(tenantid string, bp *entities.BPrint) (string, error)
	BprintUpdate(tenantid string, bp *entities.BPrint) error
	BprintGet(tenantid, bid string) (*entities.BPrint, error)
	BprintRemove(tenantid, bid string) error
	BprintListBlobs(tenantid, bid string) (map[string]string, error)
}

type BStore interface {
	NewRoot(tenantid, bid string) error
	DeleteRoot(tenantid, bid string) error
	NewBlob(tenantid, bid, folder, file string, payload []byte) error
	NewFolder(tenantid, bid, folder string) error
	UpdateBlob(tenantid, bid, folder, file string, payload []byte) error
	GetBlob(tenantid, bid, folder, file string) ([]byte, error)
	DeleteBlob(tenantid, bid, folder, file string) error
}
