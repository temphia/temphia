package wasmer2

import (
	"encoding/binary"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func (w *wasmer2) wasmErr(err error) []wasmer.Value {
	if err == nil {
		return w.wasmOk()
	}

	errBytes := []byte(err.Error())
	lenErrbytes := -int32(len(errBytes))
	w.pendingResp = errBytes
	return []wasmer.Value{wasmer.NewI32(lenErrbytes)}
}

func (w *wasmer2) wasmResp(out []byte, err error) []wasmer.Value {
	if err != nil {
		return w.wasmErr(err)
	}

	if out == nil {
		return []wasmer.Value{wasmer.NewI32(0)}
	}

	respLen := len(out)
	w.pendingResp = out

	return []wasmer.Value{wasmer.NewI32(respLen)}
}

func (w *wasmer2) wasmOk() []wasmer.Value {
	return []wasmer.Value{wasmer.NewI32(0)}
}

func (w *wasmer2) getMemory() []byte {
	mem, err := w.instance.Exports.GetMemory("memory")
	if err != nil {
		panic(err)
	}
	return mem.Data()
}

func writeInt(mem []byte, ptr int32, ival int32) {
	binary.LittleEndian.PutUint32(mem[:ptr], uint32(ival))
}

func getStr(mem []byte, ptr, len int32) string {
	return string(mem[ptr : ptr+len])
}

func getByte(mem []byte, ptr, len int32) []byte {
	return mem[ptr : ptr+1]
}

type ExternFn func(args []wasmer.Value) ([]wasmer.Value, error)

type BindOptions struct {
	name    string
	kinds   []wasmer.ValueKind
	fn      func(args []wasmer.Value) ([]wasmer.Value, error)
	returns []wasmer.ValueKind
}

func (w *wasmer2) bind(opts BindOptions) {
	w.extenFns[opts.name] = wasmer.NewFunction(
		w.store,
		wasmer.NewFunctionType(
			wasmer.NewValueTypes(opts.kinds...),
			wasmer.NewValueTypes(opts.returns...),
		),
		opts.fn,
	)
}
