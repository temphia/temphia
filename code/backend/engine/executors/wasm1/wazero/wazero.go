package wazero

import (
	"context"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type Executor struct {
	builder   *Builder
	compiled  wazero.CompiledModule
	instance  api.Module
	allocator api.Function

	// cached/ cleared between each Process
	context context.Context
	mem     api.Memory

	// bindings cache
	bindings  bindx.Bindings
	invoker   bindx.Invoker
	bindSelf  bindx.Self
	bindPluKV bindx.PlugKV
}

func (e *Executor) Process(req *event.Request) (*event.Response, error) {
	e.context = context.WithValue(context.Background(), ExecutorCtx, e)
	e.mem = e.instance.Memory()

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
