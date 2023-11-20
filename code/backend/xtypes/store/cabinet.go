package store

import (
	"context"
	"io"

	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

const (
	DefaultDataAssetsFolder   = "data_common"
	DefaultBprintFolder       = "bprints"
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
	AddFolder(ctx context.Context, tenant, folder string) error
	AddBlob(ctx context.Context, tenant, folder string, file string, contents []byte) error
	AddBlobStreaming(ctx context.Context, tenant, folder string, file string, contents io.ReadCloser) error
	ListRoot(ctx context.Context, tenant string) ([]string, error)
	ListFolderBlobs(ctx context.Context, tenant, folder string) ([]*BlobInfo, error)
	GetBlob(ctx context.Context, tenant, folder string, file string) ([]byte, error)
	GetFolderAsZip(ctx context.Context, tenant, folder string) (string, error)
	DeleteBlob(ctx context.Context, tenant, folder string, file string) error
}

// fixme => ref v2
type FileStore2 interface {
	ListFolder(fpath string) ([]string, error)
	NewFolder(fpath, name string) error
	DeleteFoler(fpath string) error
	RenameFolder(fpath, newname string) error
	CompressFolder(fpath string) (FData, error)
	TreeFolder(fpath string) ([]string, error)

	GetFile(fpath string) (FData, error)
	RenameFile(fpath, name, newname string) error
	DuplicateFile(fpath, name, newname string) error
	MoveFile(fpath, newfpath string) error
	NewFile(fpath, name string, data FData) error
	UpdateFile(fpath, name string) error
	DeleteFile(fpath, name string) error
	CompressFiles(fpath string, files []string) (FData, error)
}

type FData interface {
	AsBytes() ([]byte, error)
	AsReader() io.ReadCloser
}
