package data

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
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
