package wazero

import (
	"context"
	"errors"
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

func (e *Executor) guestAllocateBytes(size uint64) uint32 {
	offset, err := e.allocator.Call(context.TODO(), size)
	if err != nil {
		panic("allocate_bytes not exported")
	}
	return uint32(offset[0])
}
