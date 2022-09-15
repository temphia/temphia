package wazero

import (
	"context"

	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type Builder struct {
	runtime wazero.Runtime
	mbinded api.Module
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	rt := wazero.NewRuntime()

	tm, err := BuildTemphiaModule(rt)
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
	runtime := wazero.NewRuntime()

	cmodule, err := runtime.CompileModule(ctx, nil, wazero.NewCompileConfig())
	if err != nil {
		return nil, err
	}

	module, err := runtime.InstantiateModule(ctx, cmodule, wazero.NewModuleConfig())
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

		bindPluKV:  opts.Binder.PlugKVBindingsGet(),
		bindSockd:  opts.Binder.SockdBindingsGet(),
		bindUser:   opts.Binder.UserBindingsGet(),
		bindCab:    opts.Binder.CabinetBindingsGet(),
		bindSelf:   opts.Binder.SelfBindingsGet(),
		bindNcache: opts.Binder.NodeCacheGet(),
		bindNet:    opts.Binder.NetGet(),
	}, nil
}
