package wasmer

import "github.com/wasmerio/wasmer-go/wasmer"

type Memory struct {
	inner    []byte
	instance *wasmer.Instance
	executor *Executor
}

func (m Memory) getString(offset wasmer.Value, len wasmer.Value) string {

	_start := offset.I32()
	_end := _start + len.I32()

	return string(m.inner[_start:_end])
}
