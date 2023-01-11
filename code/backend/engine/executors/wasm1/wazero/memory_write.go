package wazero

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"
)

func (e *Executor) writeFinal(respOffset, respLen int32, err error) int32 {
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	return 1
}

func (e *Executor) writeJSONFinal(respOffset, respLen int32, resp any, err error) int32 {
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	return e.writeJSON(respOffset, respLen, resp)
}

func (e *Executor) writeBytesFinal(respOffset, respLen int32, resp []byte, err error) int32 {
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}
	e.writeBytesNPtr(resp, respOffset, respLen)
	return 1
}

func (e *Executor) writeJSON(respOffset, respLen int32, resp any) int32 {

	out, err := json.Marshal(resp)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	offset := e.guestAllocateBytes(uint64(len(out)))
	ok := e.mem.Write(e.context, offset, out)
	if !ok {
		panic(ErrOutofIndex)
	}

	e.mem.WriteUint32Le(e.context, uint32(respOffset), offset)
	e.mem.WriteUint32Le(e.context, uint32(respLen), uint32(len(out)))

	return 1
}

func (e *Executor) writeError(respOffset, respLen int32, err error) {

	errstr := kosher.Byte(err.Error())
	offset := e.guestAllocateBytes(uint64(len(errstr)))
	ok := e.mem.Write(e.context, offset, errstr)
	if !ok {
		panic(ErrOutofIndex)
	}

	e.mem.WriteUint32Le(e.context, uint32(respOffset), offset)
	e.mem.WriteUint32Le(e.context, uint32(respLen), uint32(len(errstr)))

}

func (e *Executor) writeBytesNPtr(data []byte, roffset, rlen int32) {
	offset := e.guestAllocateBytes(uint64(len(data)))

	ok := e.mem.Write(e.context, offset, data)
	if !ok {
		panic(ErrOutofIndex)

	}

	e.mem.WriteUint32Le(e.context, uint32(roffset), offset)
	e.mem.WriteUint32Le(e.context, uint32(rlen), uint32(len(data)))
}
