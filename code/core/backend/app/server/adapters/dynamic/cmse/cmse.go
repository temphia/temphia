package cmse

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

type CMSE struct {
	app     xtypes.App
	runtime etypes.Runtime

	Assets            []Asset
	Actions           map[string][]Func
	Filters           map[string][]Func
	ShortCodeResolver map[string][]Func
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {
	engine := opts.App.GetDeps().Engine().(etypes.Engine)

	return &CMSE{
		app:     opts.App,
		runtime: engine.GetRuntime(),
	}, nil
}

func (d *CMSE) Handle(ctx httpx.Context) {

}
