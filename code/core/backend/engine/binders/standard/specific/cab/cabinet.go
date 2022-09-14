package cab

import (
	"context"

	"github.com/temphia/temphia/code/core/backend/engine/binders/standard/handle"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type Binding struct {
	chub     store.CabinetHub
	tenantId string
}

func New(handle *handle.Handle) Binding {

	return Binding{
		chub:     handle.Deps.CabinetHub,
		tenantId: handle.Namespace,
	}
}

func (b *Binding) resourceFolder(bucket string) (string, string) {

	// FIXME =>
	return "", ""
}

func (b *Binding) AddFile(bucket string, file string, contents []byte) error {
	folder, source := b.resourceFolder(bucket)
	return b.chub.GetSource(source, b.tenantId).AddBlob(context.TODO(), folder, file, contents)
}

func (b *Binding) ListFolder(bucket string) ([]string, error) {
	folder, source := b.resourceFolder(bucket)

	files, err := b.source(source).ListFolder(context.TODO(), folder)
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
	folder, src := b.resourceFolder(bucket)
	source := b.source(src)
	return source.GetBlob(context.TODO(), folder, file)
}
func (b *Binding) DeleteFile(bucket string, file string) error {
	folder, src := b.resourceFolder(bucket)
	source := b.source(src)
	return source.DeleteBlob(context.TODO(), folder, file)
}

func (b *Binding) GenerateTicket(bucket string, ticket *bindx.CabTicket) (string, error) {
	return "", nil
}

// private

func (b *Binding) source(src string) store.CabinetSourced {
	return b.chub.GetSource(src, b.tenantId)
}
