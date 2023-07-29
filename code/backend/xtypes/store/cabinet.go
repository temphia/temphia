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
}

type FileStoreHub interface {
	FileStore
}

type FileStore interface {
	AddFolder(ctx context.Context, tenant, folder string) error
	AddBlob(ctx context.Context, tenant, folder string, file string, contents []byte) error
	AddBlobStreaming(ctx context.Context, tenant, folder string, file string, contents io.ReadCloser) error
	ListRoot(ctx context.Context, tenant string) ([]string, error)
	ListFolderBlobs(ctx context.Context, tenant, folder string) ([]*BlobInfo, error)
	GetBlob(ctx context.Context, tenant, folder string, file string) ([]byte, error)
	DeleteBlob(ctx context.Context, tenant, folder string, file string) error
}
