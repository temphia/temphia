package cab

import (
	"context"

	"github.com/temphia/temphia/code/backend/engine/binder/handle"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx/ticket"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Binding struct {
	chub     store.CabinetHub
	handle   *handle.Handle
	tenantId string
}

func New(handle *handle.Handle) Binding {

	return Binding{
		chub:     handle.Deps.CabinetHub,
		tenantId: handle.Namespace,
	}
}

func (b *Binding) resourceFolder(bucket string) (string, string) {
	b.handle.LoadResources()

	res := b.handle.Resources[bucket]
	if res == nil {
		panic("Could not laod resource folder")
	}

	targets, err := res.SplitTarget(2)
	if err != nil {
		panic("parse resource target err")
	}

	return targets[0], targets[1]
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

func (b *Binding) Ticket(bucket string, opts *ticket.CabinetFolder) (string, error) {
	source, folder := b.resourceFolder(bucket)
	uctx := b.handle.Job.Invoker.UserContext()
	if uctx == nil {
		return "", easyerr.Error(etypes.EmptyUserContext)
	}

	return b.handle.Deps.Signer.SignFolder(b.tenantId, &claim.Folder{
		TenantId:  b.tenantId,
		UserId:    uctx.Id,
		SessionID: uctx.SessionId,
		DeviceId:  uctx.SessionId,
		Type:      "",
		Expiry:    0,
		Source:    source,
		Folder:    folder,
	})

}

// private

func (b *Binding) source(src string) store.CabinetSourced {
	return b.chub.GetSource(src, b.tenantId)
}
