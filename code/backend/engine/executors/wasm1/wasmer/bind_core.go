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
	e.bind(BindOptions{
		name:    "_get_self_file",
		kinds:   []wasmer.ValueKind{wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32},
		fn:      e.getSelfFile,
		returns: []wasmer.ValueKind{wasmer.I32},
	})

}

func (w *Executor) log(args []wasmer.Value) ([]wasmer.Value, error) {
	mem := w.getMemory()
	w.bindings.Log(mem.getString(args[0], args[1]))
	return EmptyAtom, nil
}

func (w *Executor) sleep(args []wasmer.Value) ([]wasmer.Value, error) {
	w.bindings.Sleep(args[0].I32())
	return EmptyAtom, nil
}

func (w *Executor) getSelfFile(args []wasmer.Value) ([]wasmer.Value, error) {
	fPtr := args[0]
	fsize := args[1]
	rPtr := args[2]
	rLen := args[3]
	modPtr := args[4]

	mem := w.getMemory()

	data, mod, err := w.bindings.GetFileWithMeta(mem.getString(fPtr, fsize))
	if err != nil {
		return mem.rwriteErr(rPtr, rLen, err)
	}

	mem.writeBytes(data, rPtr, rLen)
	mem.writeUint64(modPtr.I32(), mod)

	return mem.rwriteBytes1(rPtr, rLen, data, err)
}
