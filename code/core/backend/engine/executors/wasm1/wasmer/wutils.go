package wasmer

import "github.com/wasmerio/wasmer-go/wasmer"

func (e *Executor) getMemory() Memory {
	return Memory{
		inner:    importMemory(e.instance),
		instance: e.instance,
		executor: e,
	}
}

func importMemory(instance *wasmer.Instance) []byte {
	mem, err := instance.Exports.GetMemory("memory")
	if err != nil {
		panic(err)
	}
	return mem.Data()
}
