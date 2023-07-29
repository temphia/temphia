package corehub

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/hub/coredb/authz"
	"github.com/temphia/temphia/code/backend/hub/coredb/statehub"
	"github.com/temphia/temphia/code/backend/xtypes"

	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

var _ store.CoreHub = (*CoreHub)(nil)

type CoreHub struct {
	coredb   store.CoreDB
	notifier sockdx.UserSyncer
	cplane   xplane.ControlPlane
	smanager authz.Manager
	stateHub statehub.StateHub
	logger   zerolog.Logger
}

func New(coredb store.CoreDB) *CoreHub {
	return &CoreHub{
		coredb:   coredb,
		cplane:   nil,
		smanager: authz.New(),
		stateHub: statehub.New(),
	}
}

func (c *CoreHub) Inject(app xtypes.App) {

	deps := app.GetDeps()
	logsvc := app.GetDeps().LogService().(logx.Service)

	c.notifier = deps.SockdHub().(sockdx.Hub).GetUserSyncer()
	c.cplane = deps.ControlPlane().(xplane.ControlPlane)
	c.logger = logsvc.GetServiceLogger("corehub")

	err := c.stateHub.Start(app)
	if err != nil {
		panic(err)
	}

}

func (c *CoreHub) Ping() error {
	return c.coredb.Ping()
}

func (c *CoreHub) GetAuthZ(tenantId, group string) store.AuthZ {

	key := tenantId + group

	scoper := c.smanager.Get(key)
	if scoper == nil {
		ug, err := c.coredb.GetUserGroup(tenantId, group)
		if err != nil {
			return nil
		}
		c.smanager.Set(key, ug)
		return c.smanager.Get(key)
	}

	return scoper
}
