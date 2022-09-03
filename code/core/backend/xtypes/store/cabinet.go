package store

import (
	"context"
	"io"
)

type BlobInfo struct {
	Name         string `json:"name,omitempty"`
	IsDir        bool   `json:"is_dir,omitempty"`
	Size         int    `json:"size,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
}

type CabinetHub interface {
	Start(eventbus interface{}) error
	Default(tenant string) CabinetSourced
	ListSources(tenant string) ([]string, error)
	GetSource(source, tenant string) CabinetSourced
	DefaultName(tenantId string) string
}

type CabinetSourced interface {
	AddFolder(Context context.Context, folder string) error
	AddBlob(Context context.Context, folder, file string, contents []byte) error // fixme putblob
	ListRoot(Context context.Context) ([]string, error)
	ListFolder(Context context.Context, folder string) ([]*BlobInfo, error)
	GetBlob(Context context.Context, folder, file string) ([]byte, error)
	DeleteBlob(Context context.Context, folder, file string) error
}

type CabinetSource interface {
	InitilizeTenent(tenent string, folders []string) error
	AddFolder(ctx context.Context, tenant, folder string) error
	AddBlob(ctx context.Context, tenant, folder string, file string, contents []byte) error
	AddBlobStreaming(ctx context.Context, tenant, folder string, file string, contents io.ReadCloser) error
	ListRoot(ctx context.Context, tenant string) ([]string, error)
	ListFolderBlobs(ctx context.Context, tenant, folder string) ([]*BlobInfo, error)
	GetBlob(ctx context.Context, tenant, folder string, file string) ([]byte, error)
	DeleteBlob(ctx context.Context, tenant, folder string, file string) error
}
