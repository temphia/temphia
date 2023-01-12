package easyserve

import (
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

type ExecServe struct {
	app      xtypes.App
	executor etypes.ExecutorBuilder
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {
	reg := opts.App.GetDeps().Registry().(registry.Registry)

	executor, err := reg.GetExecutors()["pageform"](opts.App)
	if err != nil {
		return nil, err
	}

	return &ExecServe{
		app:      opts.App,
		executor: executor,
	}, nil
}

func (e *ExecServe) ServeEditorFile(file string) ([]byte, error) {
	return e.executor.ExecFile(file)
}

func (e *ExecServe) PreformEditorAction(name string, data []byte) (any, error) {

	e.executor.Instance(etypes.ExecutorOption{})

	return nil, nil
}

func (e *ExecServe) Handle(ctx httpx.Context) {

}
