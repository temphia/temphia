package noop

import (
	"fmt"

	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

type NoOp struct {
	rendered []byte
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	source := fmt.Sprintf(`
		<h1>It Works!</h1>

		<p>
		you should probably change to more useful adapter or <a href="/z/auth">click</a> to goto portal auth page.
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
	}, nil
}

func (l *NoOp) ServeEditorFile(file string) ([]byte, error) {
	return []byte(``), nil
}

func (l *NoOp) PreformEditorAction(name string, data []byte) (any, error) {
	return nil, nil
}

func (l *NoOp) Handle(ctx httpx.Context) {
	ctx.Http.Writer.Write(l.rendered)
}
