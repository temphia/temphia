package noop

import (
	_ "embed"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

//go:embed loader.js
var loader []byte

type NoopBuilder struct {
	app xtypes.App
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	return &NoopBuilder{
		app: app.(xtypes.App),
	}, nil
}

func (nb *NoopBuilder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {
	return &Noop{builder: nb, opts: opts}, nil
}

func (nb *NoopBuilder) ExecFile(file string) ([]byte, error) {
	switch file {
	case "loader.js":
		return loader, nil
	default:
		return []byte(``), nil
	}
}
