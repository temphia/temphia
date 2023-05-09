package static

import (
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

type static struct {
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {
	return &static{}, nil
}

func (s *static) ServeEditorFile(file string) ([]byte, error) {
	return nil, nil
}

func (d *static) PreformEditorAction(name string, data []byte) (any, error) {
	return nil, nil
}

func (s *static) Handle(ctx httpx.Context) {}

func (s *static) Close() error {
	return nil
}
