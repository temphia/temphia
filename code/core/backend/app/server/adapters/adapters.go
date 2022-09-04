package renderer

import (
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/app/server/adapters/dynamic"
	"github.com/temphia/temphia/code/core/backend/app/server/adapters/landings"
	"github.com/temphia/temphia/code/core/backend/app/server/adapters/singleton"
	"github.com/temphia/temphia/code/core/backend/app/server/adapters/static"
)

func init() {
	registry.G.SetAapterBuilder("dynamic", dynamic.New)
	registry.G.SetAapterBuilder("landings", landings.New)
	registry.G.SetAapterBuilder("singleton", singleton.New)
	registry.G.SetAapterBuilder("static", static.New)
}
