package renderer

import (
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/app/server/adapters/easypage"
	"github.com/temphia/temphia/code/backend/app/server/adapters/noop"
	"github.com/temphia/temphia/code/backend/app/server/adapters/static"
	"github.com/temphia/temphia/code/backend/app/server/adapters/web2agent"
)

func init() {

	registry.G.SetAapterBuilder("noop", noop.New)
	registry.G.SetAapterBuilder("", noop.New)
	registry.G.SetAapterBuilder("static", static.New)
	registry.G.SetAapterBuilder("easypage", easypage.New)
	registry.G.SetAapterBuilder("web2agent", web2agent.New)

}
