package bprint

import (
	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
)

type BprintMod struct {
	tenantId string
	bid      string
	bhub     repox.Hub
	modipc   *modipc.ModIPC
}

func (bm *BprintMod) Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error) {
	return bm.modipc.Handle(method, args)
}

func (bm *BprintMod) Close() error {
	bm.modipc = nil

	return nil
}

type newBlobOptions struct {
	File    string `json:"file,omitempty"`
	Payload []byte `json:"payload,omitempty"`
}

func (bm *BprintMod) NewBlob(opts *newBlobOptions) error {
	return bm.bhub.BprintNewBlob(bm.tenantId, bm.bid, opts.File, opts.Payload, true)
}

type updateBlobOptions struct {
	File    string `json:"file,omitempty"`
	Payload []byte `json:"payload,omitempty"`
}

func (bm *BprintMod) UpdateBlob(opts *updateBlobOptions) error {
	return bm.bhub.BprintUpdateBlob(bm.tenantId, bm.bid, opts.File, opts.Payload)
}

type fileOpts struct {
	File string `json:"file,omitempty"`
}

func (bm *BprintMod) GetBlob(opts *fileOpts) ([]byte, error) {
	return bm.bhub.BprintGetBlob(bm.tenantId, bm.bid, opts.File)
}

func (bm *BprintMod) DeleteBlob(opts *fileOpts) error {
	return bm.bhub.BprintDeleteBlob(bm.tenantId, bm.bid, opts.File)
}

// builder

var _ etypes.ModuleBuilder = (*BprintModBuilder)(nil)

func NewBuilder(app any) (etypes.ModuleBuilder, error) {
	return &BprintModBuilder{}, nil
}

type BprintModBuilder struct{}

func (p *BprintModBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	return New(
		opts.Resource.TenantId,
		opts.Resource.Target,
		opts.Binder.GetApp().(xtypes.App),
	), nil
}

func New(tenantId, bid string, app xtypes.App) *BprintMod {

	bm := &BprintMod{
		tenantId: tenantId,
		bid:      bid,
		modipc:   nil,
		bhub:     app.GetDeps().RepoHub().(repox.Hub),
	}

	bm.modipc = modipc.NewModIPC(bm)

	return bm

}
