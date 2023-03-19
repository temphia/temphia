package wasmer

import (
	"crypto/sha1"
	"sync"

	"github.com/hashicorp/golang-lru/simplelru"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/wasmerio/wasmer-go/wasmer"
)

type Builder struct {
	engine *wasmer.Engine
	store  *wasmer.Store

	lru     simplelru.LRUCache
	lruLock sync.Mutex
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	lru, err := simplelru.NewLRU(20, nil)
	if err != nil {
		return nil, err
	}

	b := &Builder{
		engine:  engine,
		store:   store,
		lru:     lru,
		lruLock: sync.Mutex{},
	}

	return etypes.ExecBuilderFunc(b.Instance), nil

}

func (b *Builder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {

	data, _, err := opts.Binder.GetFileWithMeta(opts.File)
	if err != nil {
		return nil, err
	}

	module, err := b.compile(data)
	if err != nil {
		return nil, err
	}

	exec := &Executor{
		builder:  b,
		module:   module,
		bindings: opts.Binder,
		instance: nil,
		extenFns: make(map[string]wasmer.IntoExtern),

		bindPluKV: opts.Binder.PlugKVBindingsGet(),
		bindSockd: opts.Binder.SockdBindingsGet(),
		bindCab:   opts.Binder.CabinetBindingsGet(),
		bindSelf:  opts.Binder.SelfBindingsGet(),
		bindNet:   opts.Binder.NetGet(),
	}

	err = exec.init()
	if err != nil {
		return nil, err
	}

	return exec, nil
}

// private

func (e *Executor) init() error {

	importObject := wasmer.NewImportObject()

	memlimit, err := wasmer.NewLimits(10, 100)
	if err != nil {
		return err
	}

	memory := wasmer.NewMemory(
		e.builder.store,
		wasmer.NewMemoryType(memlimit),
	)

	importObject.Register("env", map[string]wasmer.IntoExtern{
		"abort": wasmer.NewFunction(
			e.builder.store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32),
				wasmer.NewValueTypes(),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"memory": memory,
	})

	e.buildBindings()

	importObject.Register("temphia1", e.extenFns)

	instance, err := wasmer.NewInstance(e.module, importObject)
	if err != nil {
		return err
	}

	e.instance = instance

	return nil
}

func (b *Builder) compile(data []byte) (*wasmer.Module, error) {

	hash := sha1.New()
	hash.Write(data)
	key := hash.Sum(nil)

	b.lruLock.Lock()

	m, ok := b.lru.Get(key)
	if !ok {
		b.lruLock.Unlock()

		module, err := wasmer.NewModule(b.store, data)
		if err != nil {
			return nil, err
		}

		b.lruLock.Lock()
		b.lru.Add(key, module)
		b.lruLock.Unlock()
	}

	return m.(*wasmer.Module), nil

}
