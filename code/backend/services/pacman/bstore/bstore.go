package bstore

import (
	"context"
	"path"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type bstore struct {
	filestore store.FileStore
}

func New(filestore store.FileStore) *bstore {
	return &bstore{
		filestore: filestore,
	}
}

func (b *bstore) NewRoot(tenantid, bid string) error {
	return b.filestore.AddFolder(context.TODO(), tenantid, bprintFolder(bid, ""))
}

func (b *bstore) DeleteRoot(tenantid, bid string) error {
	return easyerr.NotImpl()
}

func (b *bstore) NewBlob(tenantid, bid, folder, file string, payload []byte) error {
	return b.filestore.AddBlob(context.TODO(), tenantid, bprintFolder(bid, folder), file, payload)
}

func (b *bstore) NewFolder(tenantid, bid, folder string) error {
	return b.filestore.AddFolder(context.TODO(), tenantid, bprintFolder(bid, folder))
}

func (b *bstore) ListBlob(tenantid, bid, folder string) ([]*store.BlobInfo, error) {
	return b.filestore.ListFolderBlobs(context.TODO(), tenantid, bprintFolder(bid, folder))
}

func (b *bstore) UpdateBlob(tenantid, bid, folder, file string, payload []byte) error {
	return b.filestore.AddBlob(context.TODO(), tenantid, bprintFolder(bid, folder), file, payload)
}

func (b *bstore) GetBlob(tenantid, bid, folder, file string) ([]byte, error) {
	return b.filestore.GetBlob(context.TODO(), tenantid, bprintFolder(bid, folder), file)
}

func (b *bstore) DeleteBlob(tenantid, bid, folder, file string) error {
	return b.filestore.DeleteBlob(context.TODO(), tenantid, bprintFolder(bid, folder), file)
}

func bprintFolder(bid, folder string) string {
	return path.Join(xtypes.BprintBlobFolder, bid, folder)
}
