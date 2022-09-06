package admin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

type Controller struct {
	pacman service.Pacman
	cplane xplane.ControlPlane
	coredb store.CoreHub
	signer service.Signer

	dynHub store.DynHub
	cabHub store.CabinetHub
}

func New(pacman service.Pacman, cplane xplane.ControlPlane, coredb store.CoreHub, signer service.Signer) *Controller {
	return &Controller{
		pacman: pacman,
		cplane: cplane,
		coredb: coredb,
		signer: signer,
		dynHub: nil,
		cabHub: nil,
	}
}
