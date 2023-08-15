package cabinet

import (
	"context"

	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Controller struct {
	hub    store.CabinetHub
	signer service.Signer
}

func New(cabinet store.CabinetHub, signer service.Signer) *Controller {
	return &Controller{
		hub:    cabinet,
		signer: signer,
	}
}

func (c *Controller) ListRoot(uclaim *claim.Session, source string) ([]string, error) {
	return c.hub.ListRoot(context.TODO(), uclaim.TenantId)
}

func (c *Controller) AddFolder(uclaim *claim.Session, source, folder string) error {
	return c.hub.AddFolder(context.TODO(), uclaim.TenantId, folder)
}

func (c *Controller) AddBlob(uclaim *claim.Session, source, folder, file string, contents []byte) error {
	return c.hub.AddBlob(context.TODO(), uclaim.TenantId, folder, file, contents)
}

func (c *Controller) ListFolder(uclaim *claim.Session, source, folder string) ([]*store.BlobInfo, error) {
	return c.hub.ListFolderBlobs(context.TODO(), uclaim.TenantId, folder)
}

func (c *Controller) GetBlob(uclaim *claim.Session, source, folder, file string) ([]byte, error) {
	return c.hub.GetBlob(context.TODO(), uclaim.TenantId, folder, file)
}

func (c *Controller) DeleteBlob(uclaim *claim.Session, source, folder, file string) error {
	return c.hub.DeleteBlob(context.TODO(), uclaim.TenantId, folder, file)
}
