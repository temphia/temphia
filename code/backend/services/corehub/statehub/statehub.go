package statehub

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

// statehub subscribes to msgbus event and apply/push that to other hub
type StateHub struct {
	datahub dyndb.DataHub
	corehub store.CoreHub
	msgbus  xplane.MsgBus
}

func New(datahub dyndb.DataHub, corehub store.CoreHub, msgbus xplane.MsgBus) *StateHub {
	return &StateHub{
		corehub: corehub,
		datahub: datahub,
		msgbus:  msgbus,
	}
}

func (r *StateHub) Run() {

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
