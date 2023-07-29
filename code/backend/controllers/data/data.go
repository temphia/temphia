package data

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type Controller struct {
	dynHub  dyndb.DataHub
	cabHub  store.CabinetHub
	signer  service.Signer
	repoman repox.Pacman
}

func New(dhub dyndb.DataHub, repoman repox.Pacman, cabHub store.CabinetHub, signer service.Signer) *Controller {
	return &Controller{
		dynHub:  dhub,
		cabHub:  cabHub,
		signer:  signer,
		repoman: repoman,
	}
}

func (d *Controller) IssueDataClaim(uclaim *claim.Session, source string, group string) (string, error) {

	dc := claim.Data{
		TenantId:   uclaim.TenantId,
		Type:       claim.CTypeData,
		UserID:     uclaim.UserID,
		UserGroup:  uclaim.UserGroup,
		SessionID:  uclaim.SessionID,
		DeviceId:   uclaim.DeviceId,
		DataSource: source,
		DataGroup:  group,
		DataTables: []string{"*"},
		IsExec:     false,
	}

	pp.Println("DATA_CLAIM =>", dc)

	return d.signer.SignData(uclaim.TenantId, &dc)
}
