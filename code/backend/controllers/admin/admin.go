package admin

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

type Controller struct {
	pacman xpacman.Pacman
	cplane xplane.ControlPlane
	coredb store.CoreHub
	signer service.Signer

	dynHub dyndb.DataHub
	cabHub store.CabinetHub
	log    logx.Proxy

	plugState store.PlugStateKV
}

func New(
	pacman xpacman.Pacman,
	cplane xplane.ControlPlane,
	coredb store.CoreHub,
	signer service.Signer,
	dynHub dyndb.DataHub,
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

func (c *Controller) HasScope(uclaim *claim.Session, sub string) bool {
	if uclaim.IsSuperAdmin() {
		return true
	}

	scoper := c.coredb.GetAuthZ(uclaim.TenantId, uclaim.UserGroup)
	if scoper != nil && (scoper.CheckScope("admin.*") || scoper.CheckScope(fmt.Sprintf("admin.%s", sub))) {
		return true
	}

	return false
}
