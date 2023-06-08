package noop

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

type NoOp struct {
	rendered []byte
	dataBox  xtypes.DataBox
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	source := fmt.Sprintf(`
		<h1>It Works!</h1>

		<p>
		you should probably change to more useful adapter. <a href="/z/auth">click</a> to goto portal auth page.
		<p>

		<details>
		<summary>Detail</summary>
		<div>Domain Id: %d </div>
		<div>Domain Name: %s </div>
		<div>Adapter Type: %s </div>
	  </details> 
	
	`,
		opts.Domain.Id,
		opts.Domain.Name,
		opts.Domain.AdapterType)

	return &NoOp{
		rendered: []byte(source),
		dataBox:  opts.App.Data(),
	}, nil
}

func (l *NoOp) ServeEditorFile(file string) ([]byte, error) {
	if file == "main.js" {
		return l.dataBox.GetAsset("build", "adapter_editor_noop.js")
	}

	return []byte(``), nil
}

func (l *NoOp) PreformEditorAction(ctx httpx.AdapterEditorContext) (any, error) {
	return nil, nil
}

func (l *NoOp) Handle(ctx httpx.Context) {
	ctx.Http.Writer.Write(l.rendered)
}

func (l *NoOp) Close() error { return nil }
