package data

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Controller struct {
	dynHub store.DataHub
	cabHub store.CabinetHub
	signer service.Signer
}

func New(dhub store.DataHub, cabHub store.CabinetHub, signer service.Signer) *Controller {
	return &Controller{

		dynHub: dhub,
		cabHub: cabHub,
		signer: signer,
	}
}

func getTarget(uclaim *claim.Data) (string, string) {
	return uclaim.DataSource, uclaim.DataGroup
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
