package wasmer

import "github.com/wasmerio/wasmer-go/wasmer"

func importMemory(instance *wasmer.Instance) []byte {
	mem, err := instance.Exports.GetMemory("memory")
	if err != nil {
		panic(err)
	}
	return mem.Data()
}
