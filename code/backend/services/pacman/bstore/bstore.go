package bstore

import (
	"context"
	"path"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/fdatautil"
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
	return b.filestore.NewFolder(context.TODO(), tenantid, xtypes.BprintBlobFolder, bid)

}

func (b *bstore) DeleteRoot(tenantid, bid string) error {
	return b.filestore.DeleteFolder(context.TODO(), tenantid, bprintFolder(bid, ""))
}

func (b *bstore) NewBlob(tenantid, bid, folder, file string, payload []byte) error {
	return b.filestore.NewFile(context.TODO(), tenantid, bprintFolder(bid, folder), file, fdatautil.NewFromBytes(payload))
}

func (b *bstore) NewFolder(tenantid, bid, folder string) error {
	return b.filestore.NewFolder(context.TODO(), tenantid, bprintFolder(bid, ""), folder)
}

func (b *bstore) ListBlob(tenantid, bid, folder string) ([]*store.BlobInfo, error) {
	return b.filestore.ListFolder(context.TODO(), tenantid, bprintFolder(bid, folder))
}

func (b *bstore) UpdateBlob(tenantid, bid, folder, file string, payload []byte) error {
	return b.filestore.UpdateFile(context.TODO(), tenantid, bprintFolder(bid, folder), file, fdatautil.NewFromBytes(payload))

}

func (b *bstore) GetBlob(tenantid, bid, folder, file string) ([]byte, error) {

	resp, err := b.filestore.GetFile(context.TODO(), tenantid, bprintFolder(bid, folder))
	if err != nil {
		return nil, err
	}

	return fdatautil.ReadAndClose(resp)
}

func (b *bstore) DeleteBlob(tenantid, bid, folder, file string) error {
	return b.filestore.DeleteFile(context.TODO(), tenantid, bprintFolder(bid, folder), file)
}

func bprintFolder(bid, folder string) string {
	return path.Join(xtypes.BprintBlobFolder, bid, folder)
}
