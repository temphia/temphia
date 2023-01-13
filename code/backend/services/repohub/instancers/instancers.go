package instancers

import (
	"github.com/temphia/temphia/code/backend/services/repohub/instancers/dtable"
	"github.com/temphia/temphia/code/backend/services/repohub/instancers/plug"
	"github.com/temphia/temphia/code/backend/services/repohub/instancers/resource"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
)

func GetInstancers(app xtypes.App) map[string]xinstance.Instancer {

	return map[string]xinstance.Instancer{
		xbprint.TypeDataGroup: dtable.New(app),
		xbprint.TypePlug:      plug.New(app),
		xbprint.TypeResource:  resource.New(app),
	}
}
