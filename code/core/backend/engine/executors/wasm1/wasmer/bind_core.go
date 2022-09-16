package wasmer

import "github.com/wasmerio/wasmer-go/wasmer"

func (e *Executor) bindCore() {

	e.bind(BindOptions{
		name:    "_log",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32},
		fn:      e.log,
		returns: []wasmer.ValueKind{},
	})

	e.bind(BindOptions{
		name:    "_sleep",
		kinds:   []wasmer.ValueKind{wasmer.I32},
		fn:      e.sleep,
		returns: []wasmer.ValueKind{},
	})

}

func (w *Executor) log(args []wasmer.Value) ([]wasmer.Value, error) {
	mem := w.getMemory()
	w.bindings.Log(mem.getString(args[0], args[1]))
	return []wasmer.Value{}, nil
}

func (w *Executor) sleep(args []wasmer.Value) ([]wasmer.Value, error) {
	w.bindings.Sleep(args[0].I32())
	return []wasmer.Value{}, nil
}
