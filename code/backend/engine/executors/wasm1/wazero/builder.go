package wazero

import (
	"context"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type Builder struct {
	runtime wazero.Runtime
	mbinded api.Module
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {
	ctx := context.Background()

	rt := wazero.NewRuntime(ctx)
	tm, err := BuildTemphiaModule(ctx, rt)
	if err != nil {
		return nil, err
	}

	b := &Builder{
		runtime: rt,
		mbinded: tm,
	}

	return etypes.ExecBuilderFunc(b.Instance), nil

}

func (b *Builder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {
	ctx := context.Background()

	data, _, err := opts.Binder.GetFileWithMeta(opts.File)
	if err != nil {
		return nil, err
	}

	cmodule, err := b.runtime.CompileModule(ctx, data)
	if err != nil {
		return nil, err
	}

	module, err := b.runtime.InstantiateModule(ctx, cmodule, wazero.NewModuleConfig())
	if err != nil {
		return nil, err
	}

	allocator := module.ExportedFunction("allocate_bytes")
	if err != nil {
		return nil, easyerr.Error("allocate_bytes not exported")
	}

	return &Executor{
		builder:   b,
		compiled:  cmodule,
		instance:  module,
		bindings:  opts.Binder,
		allocator: allocator,

		context: nil,
		mem:     nil,

		bindSelf: opts.Binder.SelfBindingsGet(),
	}, nil
}
