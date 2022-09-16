package wasmer

import (
	"encoding/json"

	"github.com/wasmerio/wasmer-go/wasmer"
)

type Memory struct {
	inner      []byte
	instance   *wasmer.Instance
	executor   *Executor
	_allocFunc func(...interface{}) (interface{}, error)
}

func (e *Executor) getMemory() Memory {

	allocfunc, err := e.instance.Exports.GetFunction("allocate_bytes")
	if err != nil {
		panic(err)
	}

	return Memory{
		inner:      importMemory(e.instance),
		instance:   e.instance,
		executor:   e,
		_allocFunc: allocfunc,
	}
}

func (m Memory) allocate(size int32) int32 {
	out, err := m._allocFunc(size)
	if err != nil {
		panic(err)
	}

	return out.(int32)
}

// read

func (m Memory) getString(offset wasmer.Value, len wasmer.Value) string {

	_start := offset.I32()
	_end := _start + len.I32()

	return string(m.inner[_start:_end])
}

func (m Memory) getBytes(offset wasmer.Value, len wasmer.Value) []byte {
	_start := offset.I32()
	_end := _start + len.I32()

	return (m.inner[_start:_end])
}

func (m Memory) getJSON(offset wasmer.Value, len wasmer.Value, target any) error {
	out := m.getBytes(offset, len)
	return json.Unmarshal(out, target)
}

// write

// func (m Memory)
