package folder

import (
	"context"

	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Binding struct {
	chub      store.CabinetHub
	tenantId  string
	cabsource string
	cabfolder string
}

func New(chub store.CabinetHub, tenantId string) Binding {

	return Binding{
		chub:     chub,
		tenantId: tenantId,
	}
}

func (b *Binding) AddFile(bucket string, file string, contents []byte) error {
	return b.source(b.cabsource).AddBlob(context.TODO(), b.cabfolder, file, contents)
}

func (b *Binding) ListFolder(bucket string) ([]string, error) {

	files, err := b.source(b.cabsource).ListFolder(context.TODO(), b.cabfolder)
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

	return b.source(b.cabsource).GetBlob(context.TODO(), b.cabfolder, file)
}
func (b *Binding) DeleteFile(bucket string, file string) error {
	return b.source(b.cabsource).DeleteBlob(context.TODO(), b.cabfolder, file)
}

// private

func (b *Binding) source(src string) store.CabinetSourced {
	return b.chub.GetSource(src, b.tenantId)
}
