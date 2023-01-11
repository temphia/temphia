package wasmer

import "github.com/wasmerio/wasmer-go/wasmer"

type BindOptions struct {
	name    string
	kinds   []wasmer.ValueKind
	fn      func(args []wasmer.Value) ([]wasmer.Value, error)
	returns []wasmer.ValueKind
}

func (e *Executor) bind(opts BindOptions) {
	e.extenFns[opts.name] = wasmer.NewFunction(
		e.builder.store,
		wasmer.NewFunctionType(
			wasmer.NewValueTypes(opts.kinds...),
			wasmer.NewValueTypes(opts.returns...),
		),
		opts.fn,
	)
}

func (e *Executor) buildBindings() {
	e.bindCore()
}
