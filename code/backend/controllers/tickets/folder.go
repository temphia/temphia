package tickets

import (
	"context"

	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

// Ticket cabinet
func (c *Controller) TicketFile(ticket *claim.Folder, file string) ([]byte, error) {
	sourced := c.cabinethub.GetSource(ticket.Source, ticket.TenantId)
	return sourced.GetBlob(context.TODO(), ticket.Folder, file)
}

func (c *Controller) TicketPreview(file string, ticket *claim.Folder) ([]byte, error) {
	sourced := c.cabinethub.GetSource(ticket.Source, ticket.TenantId)
	return sourced.GetBlob(context.TODO(), ticket.Folder, file)
}

func (c *Controller) TicketList(ticket *claim.Folder) ([]*store.BlobInfo, error) {
	sourced := c.cabinethub.GetSource(ticket.Source, ticket.TenantId)
	return sourced.ListFolder(context.TODO(), ticket.Folder)
}

func (c *Controller) TicketUpload(ticket *claim.Folder, file string, data []byte) error {
	// fixme =>  send back upload proof token
	sourced := c.cabinethub.GetSource(ticket.Source, ticket.TenantId)
	return sourced.AddBlob(context.TODO(), ticket.Folder, file, data)
}
