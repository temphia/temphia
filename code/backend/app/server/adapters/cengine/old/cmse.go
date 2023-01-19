package cmse

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

type CMSE struct {
	app     xtypes.App
	runtime etypes.Runtime

	Assets  []Asset
	Actions map[string][]Func
	Filters map[string][]Func
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {
	engine := opts.App.GetDeps().Engine().(etypes.Engine)

	return &CMSE{
		app:     opts.App,
		runtime: engine.GetRuntime(),
	}, nil
}

func (d *CMSE) ServeEditorFile(file string) ([]byte, error)               { return nil, nil }
func (d *CMSE) PreformEditorAction(name string, data []byte) (any, error) { return nil, nil }
func (d *CMSE) Handle(ctx httpx.Context)                                  {}
