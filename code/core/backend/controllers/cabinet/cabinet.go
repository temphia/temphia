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

func (c *Controller) ListRoot(uclaim *claim.Cabinet) ([]string, error) {
	sourced := c.hub.GetSource(uclaim.Source, uclaim.TenentId)
	return sourced.ListRoot(context.TODO())
}

func (c *Controller) AddFolder(uclaim *claim.Cabinet, folder string) error {
	sourced := c.hub.GetSource(uclaim.TenentId, uclaim.TenentId)
	return sourced.AddFolder(context.TODO(), folder)
}

func (c *Controller) AddBlob(uclaim *claim.Session, folder, file string, contents []byte) error {
	err := c.canAction(uclaim, "add_blob", (folder))
	if err != nil {
		return err
	}
	sourced := c.hub.GetSource("uclaim.Path[1]", uclaim.TenentId)

	return sourced.AddBlob(context.TODO(), folder, file, contents)
}

func (c *Controller) ListFolder(uclaim *claim.Session, folder string) ([]*store.BlobInfo, error) {
	err := c.canAction(uclaim, "list_folder", folder)
	if err != nil {
		return nil, err
	}

	sourced := c.hub.GetSource("uclaim.Path[1]", uclaim.TenentId)

	return sourced.ListFolder(context.TODO(), folder)
}

func (c *Controller) GetBlob(uclaim *claim.Session, folder, file string) ([]byte, error) {
	err := c.canAction(uclaim, "get_blob", folder)
	if err != nil {
		return nil, err
	}
	sourced := c.hub.GetSource("uclaim.Path[1]", uclaim.TenentId)

	return sourced.GetBlob(context.TODO(), folder, file)
}

func (c *Controller) DeleteBlob(uclaim *claim.Session, folder, file string) error {
	err := c.canAction(uclaim, "del_blob", folder)
	if err != nil {
		return err
	}
	sourced := c.hub.GetSource("uclaim.Path[1]", uclaim.TenentId)

	return sourced.DeleteBlob(context.TODO(), folder, file)
}

func (c *Controller) NewFolderTicket(uclaim *claim.Session, folder string) (string, error) {

	claim := &claim.Cabinet{
		Folder: folder,
		Source: "uclaim.Path[1]",
		Expiry: 0,

		//		DeviceId: uclaim.DeviceId,
	}

	return c.signer.SignCabinet(uclaim.TenentId, claim)
}

// Ticket cabinet
func (c *Controller) TicketFile(tenantId, file string, ticket *claim.Cabinet) ([]byte, error) {
	// if !ticket.AllowGet {
	// 	return nil, easyerr.NotAuthorized()
	// }

	// if ticket.PinnedFiles != nil {
	// 	if !funk.ContainsString(ticket.PinnedFiles, file) {
	// 		return nil, easyerr.NotAuthorized()
	// 	}
	// }

	// fixme => check prefix

	sourced := c.hub.GetSource(ticket.Source, tenantId)
	return sourced.GetBlob(context.TODO(), ticket.Folder, file)
}

func (c *Controller) TicketPreview(tenantId, file string, ticket *claim.Cabinet) ([]byte, error) {
	// if !ticket.AllowGet {
	// 	return nil, easyerr.NotAuthorized()
	// }
	// fixme => implement preview
	sourced := c.hub.GetSource(ticket.Source, tenantId)
	return sourced.GetBlob(context.TODO(), ticket.Folder, file)
}

func (c *Controller) TicketList(tenantId string, ticket *claim.Cabinet) ([]*store.BlobInfo, error) {
	// pp.Println(ticket)
	// if !ticket.AllowList {
	// 	return nil, easyerr.NotAuthorized()
	// }

	sourced := c.hub.GetSource(ticket.Source, tenantId)
	return sourced.ListFolder(context.TODO(), ticket.Folder)
}

func (c *Controller) TicketUpload(tenantId, file string, data []byte, ticket *claim.Cabinet) error {
	// fixme =>  send back upload proof token
	sourced := c.hub.GetSource(ticket.Source, tenantId)
	return sourced.AddBlob(context.TODO(), ticket.Folder, file, data)
}
