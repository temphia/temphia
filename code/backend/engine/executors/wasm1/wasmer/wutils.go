package wasmer

import "github.com/wasmerio/wasmer-go/wasmer"

var (
	OkAtom    = []wasmer.Value{wasmer.NewF32(1)}
	ErrAtom   = []wasmer.Value{wasmer.NewF32(0)}
	EmptyAtom = []wasmer.Value{}
)

func importMemory(instance *wasmer.Instance) []byte {
	mem, err := instance.Exports.GetMemory("memory")
	if err != nil {
		panic(err)
	}
	return mem.Data()
}
