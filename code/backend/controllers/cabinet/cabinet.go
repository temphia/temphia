package cabinet

import (
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
	// sourced := c.hub.GetSource(source, uclaim.TenantId)
	// return sourced.ListRoot(context.TODO())
	return nil, nil
}

func (c *Controller) AddFolder(uclaim *claim.Session, source, folder string) error {
	// sourced := c.hub.GetSource(source, uclaim.TenantId)
	// return sourced.AddFolder(context.TODO(), folder)
	return nil
}

func (c *Controller) AddBlob(uclaim *claim.Session, source, folder, file string, contents []byte) error {

	// sourced := c.hub.GetSource(source, uclaim.TenantId)
	// return sourced.AddBlob(context.TODO(), folder, file, contents)

	return nil
}

func (c *Controller) ListFolder(uclaim *claim.Session, source, folder string) ([]*store.BlobInfo, error) {
	// sourced := c.hub.GetSource(source, uclaim.TenantId)
	// return sourced.ListFolder(context.TODO(), folder)
	return nil, nil
}

func (c *Controller) GetBlob(uclaim *claim.Session, source, folder, file string) ([]byte, error) {
	// sourced := c.hub.GetSource(source, uclaim.TenantId)
	// return sourced.GetBlob(context.TODO(), folder, file)

	return nil, nil
}

func (c *Controller) DeleteBlob(uclaim *claim.Session, source, folder, file string) error {
	// sourced := c.hub.GetSource(source, uclaim.TenantId)
	// return sourced.DeleteBlob(context.TODO(), folder, file)
	return nil
}
