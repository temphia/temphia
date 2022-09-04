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
