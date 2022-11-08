package datahub

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

type DataHub struct {
	dyndbs   map[string]store.DynDB
	eventHub xplane.EventBus
	sockdhub sockdx.DataSyncer
	engine   etypes.Engine
	corehub  store.CoreHub
}

func New(dyns map[string]store.DynDB) *DataHub {

	return &DataHub{
		dyndbs:   dyns,
		eventHub: nil,
		sockdhub: nil,
		engine:   nil,
		corehub:  nil,
	}
}

func (s *DataHub) Inject(_app xtypes.App) {
	deps := _app.GetDeps()

	deps.ControlPlane().(xplane.ControlPlane).GetEventBus()
	sockdhub := deps.SockdHub().(sockdx.Hub)

	s.sockdhub = sockdhub.GetDataSyncer()
	s.corehub = deps.CoreHub().(store.CoreHub)

}

func (s *DataHub) GetSource(source, tenant string) store.DynSource {
	return &dynSource{
		hub:      s,
		source:   source,
		tenantId: tenant,
	}
}

func (s *DataHub) DefaultSource(tenant string) store.DynSource {

	// fixme => cache this

	tdata, err := s.corehub.GetTenant(tenant)
	if err != nil {
		panic(err)
	}

	source := "default"
	if tdata.DefaultDSource != "" {
		source = tdata.DefaultDSource
	}

	return &dynSource{
		hub:      s,
		source:   source,
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
