package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

type Controller struct {
	pacman repox.Hub
	cplane xplane.ControlPlane
	coredb store.CoreHub
	signer service.Signer

	dynHub store.DataHub
	cabHub store.CabinetHub
	log    logx.Proxy

	plugState store.PlugStateKV
}

func New(
	pacman repox.Hub,
	cplane xplane.ControlPlane,
	coredb store.CoreHub,
	signer service.Signer,
	dynHub store.DataHub,
	cabHub store.CabinetHub,
	plugState store.PlugStateKV,
	log logx.Proxy) *Controller {

	return &Controller{

		cplane:    cplane,
		coredb:    coredb,
		signer:    signer,
		pacman:    pacman,
		dynHub:    dynHub,
		cabHub:    cabHub,
		log:       log,
		plugState: plugState,
	}
}
