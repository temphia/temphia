package datahub

import (
	"sync"

	"github.com/temphia/temphia/code/backend/services/datahub2/handle"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

var _ dyndb.DataHub = (*DataHub)(nil)

type DataHub struct {
	sources map[string]*DataSource
	handle  *handle.Handle
}

func New(dyns map[string]dyndb.DynDB) *DataHub {

	handle := &handle.Handle{
		EventHub: nil,
		SockdHub: nil,
		Engine:   nil,
		CoreHub:  nil,
	}

	sources := make(map[string]*DataSource)

	for k, dyn := range dyns {
		sources[k] = &DataSource{
			inner:  dyn,
			handle: handle,
			name:   k,
			groups: make(map[string]dyndb.DataTableHub),
			gLock:  sync.RWMutex{},
			sheets: make(map[string]dyndb.DataSheetHub),
			sLock:  sync.RWMutex{},
		}
	}

	dhub := &DataHub{
		sources: sources,
		handle:  handle,
	}

	return dhub
}

func (d *DataHub) DefaultSource(tenant string) dyndb.DynSource {
	// fixme =>

	tdata, err := d.handle.CoreHub.GetTenant(tenant)
	if err != nil {
		panic(err)
	}

	source := "default"
	if tdata.DefaultDSource != "" {
		source = tdata.DefaultDSource
	}

	return d.sources[source]

}

func (d *DataHub) GetSource(source, tenant string) dyndb.DynSource {
	return d.sources[source]
}

func (d *DataHub) ListSources(tenant string) ([]string, error) {

	sources := make([]string, 0, len(d.sources))
	for srcName := range d.sources {
		sources = append(sources, srcName)
	}

	return sources, nil

}

func (d *DataHub) Inject(_app xtypes.App) {

	deps := _app.GetDeps()

	deps.ControlPlane().(xplane.ControlPlane).GetEventBus()
	sockdhub := deps.SockdHub().(sockdx.Hub)
	d.handle.SockdHub = sockdhub.GetDataSyncer()
	d.handle.CoreHub = deps.CoreHub().(store.CoreHub)
}

func (d *DataHub) GetDataTableHub(source, tenantId, group string) dyndb.DataTableHub {
	s := d.sources[source]
	if s == nil {
		return nil
	}

	return s.GetDataTableHub(tenantId, group)
}

func (d *DataHub) GetDataSheetHub(source, tenantId, group string) dyndb.DataSheetHub {
	s := d.sources[source]
	if s == nil {
		return nil
	}

	return s.GetDataSheetHub(tenantId, group)
}
