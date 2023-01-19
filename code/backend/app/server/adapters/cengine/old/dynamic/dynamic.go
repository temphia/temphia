package dynamic

import (
	"sync"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

type dynamic struct {
	app     xtypes.App
	runtime etypes.Runtime

	evalProgram *vm.Program
	evalVMPool  sync.Pool
}

func New(opts httpx.BuilderOptions) (httpx.Adapter, error) {
	engine := opts.App.GetDeps().Engine().(etypes.Engine)

	program, err := expr.Compile(opts.Domain.AdapterPolicy)
	if err != nil {
		return nil, err
	}

	return &dynamic{
		app:         opts.App,
		runtime:     engine.GetRuntime(),
		evalProgram: program,
		evalVMPool: sync.Pool{
			New: func() any {
				return &vm.VM{}
			},
		},
	}, nil
}

func (d *dynamic) ServeEditorFile(file string) ([]byte, error) { return nil, nil }

func (d *dynamic) PreformEditorAction(name string, data []byte) (any, error) { return nil, nil }

func (d *dynamic) Handle(ctx httpx.Context) {

	vm := d.evalVMPool.Get().(*vm.VM)

	_, err := vm.Run(d.evalProgram, map[string]any{})
	if err != nil {
		return
	}

}
