package instancers

import (
	"github.com/temphia/temphia/code/backend/services/repohub/instancerhub/instancers/dtable"
	"github.com/temphia/temphia/code/backend/services/repohub/instancerhub/instancers/plug"
	"github.com/temphia/temphia/code/backend/services/repohub/instancerhub/instancers/resource"
	"github.com/temphia/temphia/code/backend/services/repohub/instancerhub/instancers/sheet"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
)

func GetInstancers(app xtypes.App) map[string]xinstance.Instancer {

	return map[string]xinstance.Instancer{
		xbprint.TypeDataGroup: dtable.New(app),
		xbprint.TypePlug:      plug.New(app),
		xbprint.TypeResource:  resource.New(app),
		xbprint.TypeDataSheet: sheet.New(app),
	}
}
