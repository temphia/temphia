package cabinet

import (
	"context"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
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
	sourced := c.hub.GetSource(source, uclaim.TenentId)
	return sourced.ListRoot(context.TODO())
}

func (c *Controller) AddFolder(uclaim *claim.Session, source, folder string) error {
	sourced := c.hub.GetSource(source, uclaim.TenentId)
	return sourced.AddFolder(context.TODO(), folder)
}

func (c *Controller) AddBlob(uclaim *claim.Session, source, folder, file string, contents []byte) error {
	err := c.canAction(uclaim, "add_blob", (folder))
	if err != nil {
		return err
	}
	sourced := c.hub.GetSource(source, uclaim.TenentId)

	return sourced.AddBlob(context.TODO(), folder, file, contents)
}

func (c *Controller) ListFolder(uclaim *claim.Session, source, folder string) ([]*store.BlobInfo, error) {
	err := c.canAction(uclaim, "list_folder", folder)
	if err != nil {
		return nil, err
	}

	sourced := c.hub.GetSource(source, uclaim.TenentId)

	return sourced.ListFolder(context.TODO(), folder)
}

func (c *Controller) GetBlob(uclaim *claim.Session, source, folder, file string) ([]byte, error) {
	err := c.canAction(uclaim, "get_blob", folder)
	if err != nil {
		return nil, err
	}
	sourced := c.hub.GetSource(source, uclaim.TenentId)

	return sourced.GetBlob(context.TODO(), folder, file)
}

func (c *Controller) DeleteBlob(uclaim *claim.Session, source, folder, file string) error {
	err := c.canAction(uclaim, "del_blob", folder)
	if err != nil {
		return err
	}
	sourced := c.hub.GetSource(source, uclaim.TenentId)

	return sourced.DeleteBlob(context.TODO(), folder, file)
}
