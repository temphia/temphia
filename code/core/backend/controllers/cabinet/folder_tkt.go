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
func (c *Controller) TicketFile(tenantId, file string, ticket *claim.Folder) ([]byte, error) {
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

func (c *Controller) TicketPreview(tenantId, file string, ticket *claim.Folder) ([]byte, error) {
	// if !ticket.AllowGet {
	// 	return nil, easyerr.NotAuthorized()
	// }
	// fixme => implement preview
	sourced := c.hub.GetSource(ticket.Source, tenantId)
	return sourced.GetBlob(context.TODO(), ticket.Folder, file)
}

func (c *Controller) TicketList(tenantId string, ticket *claim.Folder) ([]*store.BlobInfo, error) {
	// pp.Println(ticket)
	// if !ticket.AllowList {
	// 	return nil, easyerr.NotAuthorized()
	// }

	sourced := c.hub.GetSource(ticket.Source, tenantId)
	return sourced.ListFolder(context.TODO(), ticket.Folder)
}

func (c *Controller) TicketUpload(tenantId, file string, data []byte, ticket *claim.Folder) error {
	// fixme =>  send back upload proof token
	sourced := c.hub.GetSource(ticket.Source, tenantId)
	return sourced.AddBlob(context.TODO(), ticket.Folder, file, data)
}
