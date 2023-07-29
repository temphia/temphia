package filestore

import (
	"context"
	"io"

	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type FileStoreHub struct {
	// impl sth like this ?
	// needsCache  bool
	// cacheBudget int64

	inner store.FileStore
}

func New(impl store.FileStore) *FileStoreHub {
	return &FileStoreHub{
		inner: impl,
	}
}

func (f *FileStoreHub) AddFolder(ctx context.Context, tenant, folder string) error {
	return f.inner.AddFolder(ctx, tenant, folder)
}

func (f *FileStoreHub) AddBlob(ctx context.Context, tenant, folder string, file string, contents []byte) error {
	return f.inner.AddBlob(ctx, tenant, folder, file, contents)
}

func (f *FileStoreHub) AddBlobStreaming(ctx context.Context, tenant, folder string, file string, contents io.ReadCloser) error {
	return f.inner.AddBlobStreaming(ctx, tenant, folder, file, contents)
}

func (f *FileStoreHub) ListRoot(ctx context.Context, tenant string) ([]string, error) {
	return f.inner.ListRoot(ctx, tenant)
}

func (f *FileStoreHub) ListFolderBlobs(ctx context.Context, tenant, folder string) ([]*store.BlobInfo, error) {
	return f.inner.ListFolderBlobs(ctx, tenant, folder)
}

func (f *FileStoreHub) GetBlob(ctx context.Context, tenant, folder string, file string) ([]byte, error) {
	return f.inner.GetBlob(ctx, tenant, folder, file)
}

func (f *FileStoreHub) DeleteBlob(ctx context.Context, tenant, folder string, file string) error {
	return f.inner.DeleteBlob(ctx, tenant, folder, file)
}
