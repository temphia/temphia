package datahub

import (
	"sync"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/services/datahub/handle"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

var _ dyndb.DataHub = (*DataHub)(nil)

type DataHub struct {
	source DataSource
	handle *handle.Handle
}

func New(dyn dyndb.DynDB) *DataHub {

	handle := &handle.Handle{
		SockdHub: nil,
		Engine:   nil,
		CoreHub:  nil,
	}

	dhub := &DataHub{
		source: DataSource{
			inner:  dyn,
			handle: handle,
			name:   "default",
			groups: make(map[string]dyndb.DataTableHub),
			gLock:  sync.RWMutex{},
			sheets: make(map[string]dyndb.DataSheetHub),
			sLock:  sync.RWMutex{},
		},
		handle: handle,
	}

	dhub.handle.GetDataSheetHub = dhub.GetDataSheetHub

	return dhub
}

func (d *DataHub) Inject(_app xtypes.App) {

	deps := _app.GetDeps()
	cplane := deps.ControlPlane().(xplane.ControlPlane)
	d.handle.MsgBus = cplane.GetMsgBus()

	sockdhub := deps.SockdHub().(sockdx.Hub)
	d.handle.SockdHub = sockdhub.GetDataSyncer()
	d.handle.CoreHub = deps.CoreHub().(store.CoreHub)
}

func (d *DataHub) GetDataTableHub(tenantId, group string) dyndb.DataTableHub {
	return d.source.GetDataTableHub(tenantId, group)
}

func (d *DataHub) GetDataSheetHub(tenantId, group string) dyndb.DataSheetHub {
	return d.source.GetDataSheetHub(tenantId, group)
}

func (d *DataHub) GetDynDB() dyndb.DynDB {
	return d.source.inner
}

func (d *DataHub) ApplyTargetHook(tenantId string, id int64, data *entities.TargetHook) {

	pp.Println("@fixme", data)
}
