package wazero

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"
)

func (e *Executor) writeFinal(ctxid, respOffset, respLen int32, err error) int32 {
	if err != nil {
		e.writeError(ctxid, respOffset, respLen, err)
		return 0
	}

	return 1
}

func (e *Executor) writeJSONFinal(ctxid, respOffset, respLen int32, resp any, err error) int32 {
	if err != nil {
		e.writeError(ctxid, respOffset, respLen, err)
		return 0
	}

	return e.writeJSON(ctxid, respOffset, respLen, resp)
}

func (e *Executor) writeJSON(ctxid, respOffset, respLen int32, resp any) int32 {

	out, err := json.Marshal(resp)
	if err != nil {
		e.writeError(ctxid, respOffset, respLen, err)
		return 0
	}

	offset := e.guestAllocateBytes(uint64(len(out)))
	ok := e.mem.Write(offset, out)
	if !ok {
		panic(ErrOutofIndex)
	}

	e.mem.WriteUint32Le(uint32(respOffset), offset)
	e.mem.WriteUint32Le(uint32(respLen), uint32(len(out)))

	return 1
}

func (e *Executor) writeError(ctxid, respOffset, respLen int32, err error) {

	errstr := kosher.Byte(err.Error())
	offset := e.guestAllocateBytes(uint64(len(errstr)))
	ok := e.mem.Write(offset, errstr)
	if !ok {
		panic(ErrOutofIndex)
	}

	e.mem.WriteUint32Le(uint32(respOffset), offset)
	e.mem.WriteUint32Le(uint32(respLen), uint32(len(errstr)))

}

func (e *Executor) writeBytesNPtr(data []byte, ctxid, roffset, rlen int32) {
	offset := e.guestAllocateBytes(uint64(len(data)))

	ok := e.mem.Write(offset, data)
	if !ok {
		panic(ErrOutofIndex)

	}

	e.mem.WriteUint32Le(uint32(roffset), offset)
	e.mem.WriteUint32Le(uint32(rlen), uint32(len(data)))
}