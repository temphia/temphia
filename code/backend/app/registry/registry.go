package registry

import (
	"errors"
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/extension"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type (
	DynamicScript func(ns string, ctx any) error
)

type Registry struct {
	repoBuilders  map[string]repox.Builder
	storeBuilders map[string]store.Builder
	extensions    map[string]extension.Builder

	freezed bool
	mlock   *sync.Mutex
}

var (
	errTooLate = errors.New("err too late")
	errTooSoon = errors.New("err too soon")
)

func New(fromGlobal bool) *Registry {
	reg := &Registry{
		freezed:       false,
		repoBuilders:  make(map[string]repox.Builder),
		storeBuilders: make(map[string]store.Builder),
		extensions:    make(map[string]extension.Builder),
		mlock:         &sync.Mutex{},
	}

	if !fromGlobal || G == nil {
		return reg
	}

	G.mlock.Lock()
	defer G.mlock.Unlock()

	for k, v := range G.storeBuilders {
		reg.storeBuilders[k] = v
	}

	for k, v := range G.repoBuilders {
		reg.repoBuilders[k] = v
	}

	for k, v := range G.extensions {
		reg.extensions[k] = v
	}

	return reg
}

func (r *Registry) Freeze() {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	r.freezed = true
}

func (r *Registry) SetRepoBuilder(name string, builder repox.Builder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}
	r.repoBuilders[name] = builder
}

func (r *Registry) SetExecutor(name string, builder etypes.BuilderFactory) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}

	r.extensions[name] = func(app xtypes.App, handle extension.Handle) error {
		eb, err := builder(app)
		if err != nil {
			return err
		}

		handle.SetExecutorBuilder(name, eb)

		return nil
	}

}

func (r *Registry) SetExecModule(name string, builder etypes.ModuleBuilderFunc) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}

	r.extensions[name] = func(app xtypes.App, handle extension.Handle) error {
		mod, err := builder(app)
		if err != nil {
			return err
		}

		handle.SetModuleBuilder(name, mod)
		return nil
	}
}

func (r *Registry) SetAapterBuilder(name string, rb httpx.Builder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}

	r.extensions[name] = func(app xtypes.App, handle extension.Handle) error {
		handle.SetAdapterBuilder(name, rb)
		return nil
	}

}

func (r *Registry) SetExtensionBuilder(name string, builder extension.Builder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}

	r.extensions[name] = builder
}

func (r *Registry) SetStoreBuilder(name string, b store.Builder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}

	r.storeBuilders[name] = b
}

func (r *Registry) GetRepoBuilders() map[string]repox.Builder {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}
	return r.repoBuilders
}

func (r *Registry) GetStoreBuilders() map[string]store.Builder {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}

	return r.storeBuilders
}

func (r *Registry) GetExecutorBuilder() map[string]extension.Builder {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}

	return r.extensions
}
