package repox

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type Hub interface {
	Start() error

	RepoCore
	RepoBprintOps

	GetInstanceHub() InstancHub
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
}
