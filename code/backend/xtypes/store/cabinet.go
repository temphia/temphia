package store

import (
	"context"
	"io"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

const (
	DefaultDataAssetsFolder   = xtypes.DydnBlobFolder
	DefaultBprintFolder       = xtypes.BprintBlobFolder
	DefaultTenantPublicFolder = "public"
	DefaultSystemIconsFolder  = "system_icons"
)

var (
	DefaultFolders = []string{
		DefaultDataAssetsFolder,
		DefaultBprintFolder,
		DefaultTenantPublicFolder,
		DefaultSystemIconsFolder,
	}
)

type BlobInfo struct {
	Name         string `json:"name,omitempty"`
	IsDir        bool   `json:"is_dir,omitempty"`
	Size         int    `json:"size,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
}

type CabinetHub interface {
	Start(mb xplane.MsgBus) error

	FileStore
}

type FileStore interface {
	ListFolder(ctx context.Context, tenantId, fpath string) ([]*BlobInfo, error)
	NewFolder(ctx context.Context, tenantId, fpath, name string) error
	DeleteFolder(ctx context.Context, tenantId, fpath string) error
	RenameFolder(ctx context.Context, tenantId, fpath, newname string) error
	CompressFolder(ctx context.Context, tenantId, fpath string) (FData, error)
	TreeFolder(ctx context.Context, tenantId, fpath string) ([]*BlobInfo, error)
	GetFile(ctx context.Context, tenantId, fpath string) (FData, error)
	RenameFile(ctx context.Context, tenantId, fpath, name, newname string) error
	DuplicateFile(ctx context.Context, tenantId, fpath, name, newname string) error
	MoveFile(ctx context.Context, tenantId, fpath, newfpath string) error
	NewFile(ctx context.Context, tenantId, fpath, name string, data FData) error
	UpdateFile(ctx context.Context, tenantId, fpath, name string, data FData) error
	DeleteFile(ctx context.Context, tenantId, fpath, name string) error
	CompressFiles(ctx context.Context, tenantId, fpath string, files []string) (FData, error)
}

type FData interface {
	AsBytes() ([]byte, error)
	AsReader() (io.Reader, error)
	Close() error
}
