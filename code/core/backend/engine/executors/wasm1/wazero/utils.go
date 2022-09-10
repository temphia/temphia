package wazero

import (
	"context"
	"encoding/json"

	"github.com/temphia/temphia/code/core/backend/libx/xutils/kosher"
	"github.com/tetratelabs/wazero/api"
)

func (e *Executor) write(data []byte) (uint32, bool) {
	offset := e.allocateBytes(uint64(len(data)))
	mem := e.getMem()

	return offset, mem.Write(e.context, offset, data)
}

func (e *Executor) write2(data []byte, roffset, rlen uint32) bool {
	offset := e.allocateBytes(uint64(len(data)))
	mem := e.getMem()

	ok := mem.Write(e.context, offset, data)
	if !ok {
		return false
	}

	mem.WriteUint32Le(e.context, roffset, offset)
	mem.WriteUint32Le(e.context, rlen, uint32(len(data)))

	return true
}

func (e *Executor) allocateBytes(size uint64) uint32 {
	fun := e.instance.ExportedFunction("allocate_bytes")
	offset, err := fun.Call(context.TODO(), size)
	if err != nil {
		panic("allocate_bytes not exported")
	}
	return uint32(offset[0])
}

func (e *Executor) getMem() api.Memory {
	return e.instance.Memory()
}

func (e *Executor) getString(offset, count uint32) string {
	out, _ := e.instance.Memory().Read(e.context, offset, count)
	return string(out)
}

func (e *Executor) writeJSON(respPtr, respLen uint32, resp any, err error) int32 {
	mem := e.getMem()
	if err != nil {
		e.writeMemError(respPtr, respLen, mem, err)
		return 0
	}

	out, err := json.Marshal(resp)
	if err != nil {
		e.writeMemError(respPtr, respLen, mem, err)
		return 0
	}

	offset := e.allocateBytes(uint64(len(out)))
	ok := mem.Write(e.context, offset, out)
	if !ok {
		panic(ErrOutofIndex)
	}

	mem.WriteUint32Le(e.context, respPtr, offset)
	mem.WriteUint32Le(e.context, respLen, uint32(len(out)))

	return 1
}

func (e *Executor) writeError(respPtr, respLen uint32, err error) {
	mem := e.getMem()

	errstr := kosher.Byte(err.Error())
	offset := e.allocateBytes(uint64(len(errstr)))
	ok := mem.Write(e.context, offset, errstr)
	if !ok {
		panic(ErrOutofIndex)
	}

	mem.WriteUint32Le(e.context, respPtr, offset)
	mem.WriteUint32Le(e.context, respLen, uint32(len(errstr)))

}

func (e *Executor) writeMemError(respPtr, respLen uint32, mem api.Memory, err error) {
	errstr := kosher.Byte(err.Error())
	offset := e.allocateBytes(uint64(len(errstr)))
	ok := mem.Write(e.context, offset, errstr)
	if !ok {
		panic(ErrOutofIndex)
	}

	mem.WriteUint32Le(e.context, respPtr, offset)
	mem.WriteUint32Le(e.context, respLen, uint32(len(errstr)))

}

func (e *Executor) getJSONObject(optPtr, optLen uint32, target any) error {
	mem := e.getMem()
	out, ok := mem.Read(e.context, optPtr, optLen)
	if !ok {
		panic(ErrOutofIndex)
	}
	return json.Unmarshal(out, target)
}
