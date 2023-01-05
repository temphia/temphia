package easypage

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

type EasyPage struct {
	options httpx.BuilderOptions
	dataBox xtypes.DataBox
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	return &EasyPage{
		options: opts,
		dataBox: opts.App.Data(),
	}, nil
}

func (e *EasyPage) ServeEditorFile(file string) ([]byte, error) {
	return e.serveEditorFile(file)
}

func (e *EasyPage) PreformEditorAction(name string, data []byte) (any, error) {
	return e.preformEditorAction(name, data)
}

func (e *EasyPage) Handle(ctx httpx.Context) {
	e.handle(ctx)
}
