package corehub

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"

	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

var _ store.CoreHub = (*CoreHub)(nil)

type CoreHub struct {
	coredb   store.CoreDB
	notifier sockdx.UserSyncer
	cplane   xplane.ControlPlane
}

func New(coredb store.CoreDB) *CoreHub {
	return &CoreHub{
		coredb: coredb,
		//		sockdhub: nil,
		cplane: nil,
	}
}

func (c *CoreHub) Inject(app xtypes.App) {

	deps := app.GetDeps()

	c.notifier = deps.SockdHub().(sockdx.Hub).GetUserSyncer()
	c.cplane = deps.ControlPlane().(xplane.ControlPlane)
}
