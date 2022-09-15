package wasmer

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/wasmerio/wasmer-go/wasmer"
)

type Builder struct {
	engine *wasmer.Engine
	store  *wasmer.Store
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	b := &Builder{
		engine: engine,
		store:  store,
	}

	return etypes.ExecBuilderFunc(b.Instance), nil

}

func (b *Builder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {

	data, _, err := opts.Binder.GetFileWithMeta(opts.File)
	if err != nil {
		return nil, err
	}

	module, err := wasmer.NewModule(b.store, data)
	if err != nil {
		return nil, err
	}

	exec := &Executor{
		builder:    b,
		module:     module,
		bindings:   opts.Binder,
		bindPluKV:  opts.Binder.PlugKVBindingsGet(),
		bindSockd:  opts.Binder.SockdBindingsGet(),
		bindUser:   opts.Binder.UserBindingsGet(),
		bindCab:    opts.Binder.CabinetBindingsGet(),
		bindSelf:   opts.Binder.SelfBindingsGet(),
		bindNcache: opts.Binder.NodeCacheGet(),
		bindNet:    opts.Binder.NetGet(),
	}

	err = exec.init()
	if err != nil {
		return nil, err
	}

	return exec, nil
}

// private

func (e *Executor) init() error {

	return nil
}
