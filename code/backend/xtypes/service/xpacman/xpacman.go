package xpacman

import (
	"io"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xinstancer"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Pacman interface {
	Start() error

	RepoCore
	RepoBprintOps

	GetBprintFileStore() BStore
	GetInstancer() xinstancer.Instancer
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
	RepoImport(tenantid string, opts *RepoImportOpts) (string, error)
	RepoList(tenantid, group string, source int64, tags ...string) ([]BPrint, error)
	RepoGet(tenantid, slug string, source int64) (*BPrint, error)
	RepoGetZip(tenantid string, source int64, slug, version string) (io.ReadCloser, error)

	BprintCreateFromZip(tenantId string, rawreader io.ReadCloser) (string, error)
}

type RepoBprintOps interface {
	BprintList(tenantid, group string) ([]*entities.BPrint, error)
	BprintCreate(tenantid string, bp *entities.BPrint) (string, error)
	BprintUpdate(tenantid, bid string, data map[string]any) error
	BprintGet(tenantid, bid string) (*entities.BPrint, error)
	BprintRemove(tenantid, bid string) error
	BprintListBlobs(tenantid, bid string) (map[string]string, error)
}

type BStore interface {
	NewRoot(tenantid, bid string) error
	DeleteRoot(tenantid, bid string) error
	ListBlob(tenantid, bid, folder string) ([]*store.BlobInfo, error)
	NewBlob(tenantid, bid, folder, file string, payload []byte) error
	NewFolder(tenantid, bid, folder string) error
	UpdateBlob(tenantid, bid, folder, file string, payload []byte) error
	GetBlob(tenantid, bid, folder, file string) ([]byte, error)
	DeleteBlob(tenantid, bid, folder, file string) error
}
