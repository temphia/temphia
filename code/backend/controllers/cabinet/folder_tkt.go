package cabinet

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (c *Controller) NewFolderTicket(uclaim *claim.Session, source, folder string) (string, error) {

	claim := &claim.Folder{
		Folder:    folder,
		Source:    source,
		Expiry:    0,
		TenantId:  uclaim.TenantId,
		UserId:    uclaim.UserID,
		SessionID: uclaim.SessionID,
	}

	return c.signer.SignFolder(uclaim.TenantId, claim)
}
