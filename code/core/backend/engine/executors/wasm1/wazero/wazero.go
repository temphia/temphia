package wazero

import (
	"context"

	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/event"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type Executor struct {
	builder  *Builder
	bindings bindx.Bindings

	compiled wazero.CompiledModule
	instance api.Module
	context  context.Context
}

func (e *Executor) Process(req *event.Request) (*event.Response, error) {
	e.context = context.WithValue(context.Background(), ExecutorCtx, e)

	err := e.execute(req.Name, req.Data)
	if err != nil {
		return nil, err
	}

	return &event.Response{
		Payload: nil,
	}, nil
}

func (e *Executor) execute(name string, data []byte) error {

	offset, ok := e.write(data)
	if !ok {
		return ErrOutofMemory
	}

	actionFunc := e.instance.ExportedFunction(name)
	_, err := actionFunc.Call(e.context, uint64(offset), uint64(len(data)))
	if err != nil {
		return err
	}

	return nil
}

func (e *Executor) write(data []byte) (uint32, bool) {
	offset := e.allocateBytes(uint64(len(data)))
	mem := e.getMem()

	return offset, mem.Write(context.TODO(), offset, data)
}

func (e *Executor) write2(data []byte, roffset, rlen uint32) bool {
	offset := e.allocateBytes(uint64(len(data)))
	mem := e.getMem()

	ok := mem.Write(context.TODO(), offset, data)
	if !ok {
		return false
	}

	mem.WriteUint32Le(context.TODO(), roffset, offset)
	mem.WriteUint32Le(context.TODO(), rlen, uint32(len(data)))

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
