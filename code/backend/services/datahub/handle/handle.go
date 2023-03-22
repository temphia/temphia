package handle

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

// shared handle between datasheethub and datatablehub etc

type Handle struct {
	MsgBus   xplane.MsgBus
	SockdHub sockdx.DataSyncer
	Engine   etypes.Engine
	CoreHub  store.CoreHub
}
