package renderer

import (
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/app/server/adapters/cmse"
	"github.com/temphia/temphia/code/core/backend/app/server/adapters/easypage"
	"github.com/temphia/temphia/code/core/backend/app/server/adapters/launcher"
	"github.com/temphia/temphia/code/core/backend/app/server/adapters/noop"
	"github.com/temphia/temphia/code/core/backend/app/server/adapters/static"
)

func init() {
	registry.G.SetAapterBuilder("cmse", cmse.New)
	registry.G.SetAapterBuilder("noop", noop.New)
	registry.G.SetAapterBuilder("", noop.New)
	registry.G.SetAapterBuilder("launcher", launcher.New)
	registry.G.SetAapterBuilder("static", static.New)
	registry.G.SetAapterBuilder("easypage", easypage.New)

	/*

		templater
		simple_blog
		static_site


	*/

}
