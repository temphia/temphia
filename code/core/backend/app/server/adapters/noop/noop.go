package noop

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
)

type NoOp struct {
	rendered []byte
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {

	return &NoOp{
		rendered: []byte(`<h1>It Works</h1>`),
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
