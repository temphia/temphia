package instancers

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox/xinstance"
)

func GetInstancers(app xtypes.App) map[string]xinstance.Instancer {

	// dtable := dtable.New(app)
	// plug := plug.New(app)

	return map[string]xinstance.Instancer{
		// xbprint.TypeDataGroup: dtable,
		// xbprint.TypeDataTable: dtable,
		// xbprint.TypePlug:      plug,
	}
}
