package statehub

import (
	"github.com/k0kubun/pp"
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

var (
	_ store.StateHub = (*StateHub)(nil)
)

// statehub subscribes to msgbus event and apply/push that to other hub
type StateHub struct {
	datahub dyndb.DataHub
	corehub store.CoreHub
	msgbus  xplane.MsgBus
	logger  zerolog.Logger
}

func New() StateHub {
	return StateHub{
		corehub: nil,
		datahub: nil,
		msgbus:  nil,
	}
}

func (r *StateHub) Start(app xtypes.App) error {

	deps := app.GetDeps()

	logsvc := app.GetDeps().LogService().(logx.Service)

	r.datahub = deps.DataHub().(dyndb.DataHub)
	r.corehub = deps.CoreHub().(store.CoreHub)
	r.msgbus = deps.ControlPlane().(xplane.ControlPlane).GetMsgBus()
	r.logger = logsvc.GetServiceLogger("statehub")

	go r.run()

	return nil
}

func (r *StateHub) run() {

	mchan := make(chan xplane.Message)

	sid, err := r.msgbus.Subscribe("target", mchan)
	if err != nil {
		panic(err)
	}

	pp.Println("subcribtion", sid)

	for {
		schange := <-mchan

		pp.Println(schange)
	}

}
