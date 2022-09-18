package instancers

import (
	"github.com/temphia/temphia/code/core/backend/services/repohub/instancers/dtable"
	"github.com/temphia/temphia/code/core/backend/services/repohub/instancers/plug"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints/instancer"
)

func GetInstancers(app xtypes.App) map[string]instancer.Instancer {

	dtable := dtable.New(app)
	plug := plug.New(app)

	return map[string]instancer.Instancer{
		bprints.TypeDataGroup: dtable,
		bprints.TypeTSchema:   dtable,
		bprints.TypePlug:      plug,
	}
}
