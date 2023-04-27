package handle

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"

	ddb "github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

// shared handle between datasheethub and datatablehub etc

type Handle struct {
	MainHub  ddb.HubProvider
	MsgBus   xplane.MsgBus
	SockdHub sockdx.DataSyncer
	Engine   etypes.Engine
	CoreHub  store.CoreHub
}
