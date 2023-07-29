package folder

import (
	"context"

	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Binding struct {
	chub      store.FileStoreHub
	tenantId  string
	cabfolder string
}

func New(chub store.FileStore, tenantId string) Binding {

	return Binding{
		chub:     chub,
		tenantId: tenantId,
	}
}

func (b *Binding) AddFile(bucket string, file string, contents []byte) error {
	return b.chub.AddBlob(context.TODO(), b.tenantId, b.cabfolder, file, contents)
}

func (b *Binding) ListFolder(bucket string) ([]string, error) {

	files, err := b.chub.ListFolderBlobs(context.TODO(), b.tenantId, b.cabfolder)
	if err != nil {
		return nil, err
	}

	resp := make([]string, 0, len(files))
	for _, bi := range files {
		resp = append(resp, bi.Name)
	}

	return resp, nil
}

func (b *Binding) GetFile(bucket string, file string) ([]byte, error) {

	return b.chub.GetBlob(context.TODO(), b.tenantId, b.cabfolder, file)
}
func (b *Binding) DeleteFile(bucket string, file string) error {
	return b.chub.DeleteBlob(context.TODO(), b.tenantId, b.cabfolder, file)
}
