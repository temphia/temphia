package cabinethub

import (
	"context"
	"io"

	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

/*
	future
	- needsCache  bool
	- cacheBudget int64
*/

type CabinetHub struct {
	inner store.FileStore
}

func New(impl store.FileStore) *CabinetHub {
	return &CabinetHub{
		inner: impl,
	}
}

func (f *CabinetHub) Start(mb xplane.MsgBus) error {
	return nil
}

func (f *CabinetHub) AddFolder(ctx context.Context, tenant, folder string) error {
	return f.inner.AddFolder(ctx, tenant, folder)
}

func (f *CabinetHub) AddBlob(ctx context.Context, tenant, folder string, file string, contents []byte) error {
	return f.inner.AddBlob(ctx, tenant, folder, file, contents)
}

func (f *CabinetHub) AddBlobStreaming(ctx context.Context, tenant, folder string, file string, contents io.ReadCloser) error {
	return f.inner.AddBlobStreaming(ctx, tenant, folder, file, contents)
}

func (f *CabinetHub) ListRoot(ctx context.Context, tenant string) ([]string, error) {
	return f.inner.ListRoot(ctx, tenant)
}

func (f *CabinetHub) ListFolderBlobs(ctx context.Context, tenant, folder string) ([]*store.BlobInfo, error) {
	return f.inner.ListFolderBlobs(ctx, tenant, folder)
}

func (f *CabinetHub) GetBlob(ctx context.Context, tenant, folder string, file string) ([]byte, error) {
	return f.inner.GetBlob(ctx, tenant, folder, file)
}

func (f *CabinetHub) DeleteBlob(ctx context.Context, tenant, folder string, file string) error {
	return f.inner.DeleteBlob(ctx, tenant, folder, file)
}

func (f *CabinetHub) GetFolderAsZip(ctx context.Context, tenant, folder string) (string, error) {
	return f.inner.GetFolderAsZip(ctx, tenant, folder)
}
