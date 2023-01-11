package wasmer

import (
	"encoding/binary"
	"encoding/json"

	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"
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

	// fixme => do we need to reimport memory []byte (inner)
	// cz mem could have expanded

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

func (m Memory) rwrite(respPtr wasmer.Value, respLen wasmer.Value, err error) ([]wasmer.Value, error) {
	if err != nil {
		m.writeError(respPtr, respLen, err)
		return ErrAtom, nil
	}

	return OkAtom, nil
}

func (m Memory) rwriteErr(respPtr wasmer.Value, respLen wasmer.Value, err error) ([]wasmer.Value, error) {
	m.writeError(respPtr, respLen, err)
	return ErrAtom, nil
}

func (m Memory) rwriteBytes1(respPtr wasmer.Value, respLen wasmer.Value, data []byte, err error) ([]wasmer.Value, error) {
	if err != nil {
		m.writeError(respPtr, respLen, err)
		return ErrAtom, nil
	}

	return m.rwriteBytes2(respPtr, respLen, data)
}

func (m Memory) rwriteBytes2(respPtr wasmer.Value, respLen wasmer.Value, data []byte) ([]wasmer.Value, error) {
	m.writeBytes(data, respPtr, respLen)
	return OkAtom, nil
}

func (m Memory) rwriteJson(respPtr wasmer.Value, respLen wasmer.Value, resp any, err error) ([]wasmer.Value, error) {
	if err != nil {
		m.writeError(respPtr, respLen, err)
		return ErrAtom, nil
	}

	m.writeJson(respPtr, respLen, resp)
	return OkAtom, nil

}

func (m Memory) writeJson(respPtr wasmer.Value, respLen wasmer.Value, resp any) error {
	out, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	m.writeBytes(out, respPtr, respLen)
	return nil
}

func (m Memory) writeError(respPtr wasmer.Value, respLen wasmer.Value, err error) {
	m.writeBytes(kosher.Byte(err.Error()), respPtr, respLen)
}

func (m Memory) writeBytes(data []byte, respPtr wasmer.Value, respLen wasmer.Value) {

	size := int32(len(data))
	ptr := m.allocate(size)

	chunk := m.inner[ptr : ptr+size]
	copy(chunk, data)

	m.writeUint32(respPtr.I32(), ptr)
	m.writeUint32(respLen.I32(), size)
}

func (m Memory) writeUint32(offset int32, value int32) {
	b := m.inner[offset : offset+4]
	binary.LittleEndian.PutUint32(b, uint32(value))
}

func (m Memory) writeUint64(offset int32, value int64) {
	b := m.inner[offset : offset+8]
	binary.LittleEndian.PutUint32(b, uint32(value))
}
