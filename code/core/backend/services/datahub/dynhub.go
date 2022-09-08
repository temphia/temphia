package datahub

import (
	sockdhub "github.com/temphia/temphia/code/core/backend/services/sockd/hub"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

type DataHub struct {
	dyndbs   map[string]store.DynDB
	eventHub xplane.EventBus
	sockdhub sockdhub.SockdHub
	engine   etypes.Engine
	app      xtypes.App
}

func New(_app xtypes.App, dyns map[string]store.DynDB) *DataHub {

	deps := _app.GetDeps()

	return &DataHub{
		dyndbs:   dyns,
		eventHub: deps.ControlPlane().(xplane.ControlPlane).GetEventBus(),
		sockdhub: *sockdhub.New(deps.Sockd().(sockdx.Sockd)),
		app:      _app,
		engine:   deps.Engine().(etypes.Engine),
	}
}

func (s *DataHub) GetSource(source, tenant string) store.DynSource {
	return &dynSource{
		hub:      s,
		source:   source,
		tenantId: tenant,
	}
}

func (s *DataHub) DefaultSource(tenant string) store.DynSource {

	chub := s.app.GetDeps().CoreHub().(store.CoreHub)

	tdata, err := chub.GetTenant(tenant)
	if err != nil {
		panic(err)
	}

	return &dynSource{
		hub:      s,
		source:   tdata.DefaultDSource,
		tenantId: tenant,
	}
}

func (s *DataHub) ListSources(tenantId string) ([]string, error) {

	sources := make([]string, 0, len(s.dyndbs))
	for srcName := range s.dyndbs {
		sources = append(sources, srcName)
	}

	return sources, nil
}
