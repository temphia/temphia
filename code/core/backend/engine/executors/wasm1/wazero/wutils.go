package wazero

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/temphia/temphia/code/core/backend/libx/xutils/kosher"
)

type executorCtxType int8

const ExecutorCtx executorCtxType = 0

var (
	ErrOutofMemory = errors.New("OUT OF MEMORY")
	ErrOutofIndex  = errors.New("OUT OF INDEX")
)

func getCtx(ctx context.Context) *Executor {
	return ctx.Value(ExecutorCtx).(*Executor)
}

// utils

func (e *Executor) writeWithOffsetPtr(data []byte, roffset, rlen uint32) bool {
	offset := e.guestAllocateBytes(uint64(len(data)))

	ok := e.mem.Write(e.context, offset, data)
	if !ok {
		return false
	}

	e.mem.WriteUint32Le(e.context, roffset, offset)
	e.mem.WriteUint32Le(e.context, rlen, uint32(len(data)))

	return true
}

func (e *Executor) getString(offset, count uint32) string {
	out, _ := e.mem.Read(e.context, offset, count)
	return string(out)
}

func (e *Executor) writeJSON(respPtr, respLen uint32, resp any, err error) int32 {

	if err != nil {
		e.writeError(respPtr, respLen, err)
		return 0
	}

	out, err := json.Marshal(resp)
	if err != nil {
		e.writeError(respPtr, respLen, err)
		return 0
	}

	offset := e.guestAllocateBytes(uint64(len(out)))
	ok := e.mem.Write(e.context, offset, out)
	if !ok {
		panic(ErrOutofIndex)
	}

	e.mem.WriteUint32Le(e.context, respPtr, offset)
	e.mem.WriteUint32Le(e.context, respLen, uint32(len(out)))

	return 1
}

func (e *Executor) writeError(respPtr, respLen uint32, err error) {

	errstr := kosher.Byte(err.Error())
	offset := e.guestAllocateBytes(uint64(len(errstr)))
	ok := e.mem.Write(e.context, offset, errstr)
	if !ok {
		panic(ErrOutofIndex)
	}

	e.mem.WriteUint32Le(e.context, respPtr, offset)
	e.mem.WriteUint32Le(e.context, respLen, uint32(len(errstr)))

}

func (e *Executor) getJSONObject(optPtr, optLen uint32, target any) error {
	out, ok := e.mem.Read(e.context, optPtr, optLen)
	if !ok {
		panic(ErrOutofIndex)
	}

	return json.Unmarshal(out, target)
}

func (e *Executor) guestAllocateBytes(size uint64) uint32 {
	offset, err := e.allocator.Call(context.TODO(), size)
	if err != nil {
		panic("allocate_bytes not exported")
	}
	return uint32(offset[0])
}
