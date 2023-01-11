package cabinet

import (
	"context"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

func (c *Controller) NewFolderTicket(uclaim *claim.Session, source, folder string) (string, error) {

	claim := &claim.Folder{
		Folder:    folder,
		Source:    source,
		Expiry:    0,
		TenentId:  uclaim.TenentId,
		UserId:    uclaim.UserID,
		SessionID: uclaim.SessionID,
	}

	return c.signer.SignFolder(uclaim.TenentId, claim)
}

// Ticket cabinet
func (c *Controller) TicketFile(ticket *claim.Folder, file string) ([]byte, error) {
	sourced := c.hub.GetSource(ticket.Source, ticket.TenentId)
	return sourced.GetBlob(context.TODO(), ticket.Folder, file)
}

func (c *Controller) TicketPreview(file string, ticket *claim.Folder) ([]byte, error) {
	sourced := c.hub.GetSource(ticket.Source, ticket.TenentId)
	return sourced.GetBlob(context.TODO(), ticket.Folder, file)
}

func (c *Controller) TicketList(ticket *claim.Folder) ([]*store.BlobInfo, error) {
	sourced := c.hub.GetSource(ticket.Source, ticket.TenentId)
	return sourced.ListFolder(context.TODO(), ticket.Folder)
}

func (c *Controller) TicketUpload(ticket *claim.Folder, file string, data []byte) error {
	// fixme =>  send back upload proof token
	sourced := c.hub.GetSource(ticket.Source, ticket.TenentId)
	return sourced.AddBlob(context.TODO(), ticket.Folder, file, data)
}
