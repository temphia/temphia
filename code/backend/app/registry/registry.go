package registry

import (
	"errors"
	"sync"

	"github.com/temphia/temphia/code/backend/app/xtension"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/adapter"
)

type (
	DynamicScript func(ns string, ctx any) error
)

type Registry struct {
	repoBuilders  map[string]xpacman.Builder
	storeBuilders map[string]store.Builder
	extensions    map[string]xtension.Builder

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
		repoBuilders:  make(map[string]xpacman.Builder),
		storeBuilders: make(map[string]store.Builder),
		extensions:    make(map[string]xtension.Builder),
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

func (r *Registry) SetRepoBuilder(name string, builder xpacman.Builder) {
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

	r.extensions[name] = func(ctx *xtension.Context) (xtension.Xtension, error) {
		eb, err := builder(ctx.App)
		if err != nil {
			return nil, err
		}

		ctx.SetExecutorBuilder(name, eb)

		return nil, nil
	}

}

func (r *Registry) SetExecModule(name string, builder etypes.ModuleBuilderFunc) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}

	r.extensions[name] = func(ctx *xtension.Context) (xtension.Xtension, error) {
		mod, err := builder(ctx.App)
		if err != nil {
			return nil, err
		}

		ctx.SetModuleBuilder(name, mod)
		return nil, nil
	}
}

func (r *Registry) SetAapterBuilder(name string, rb adapter.Builder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}

	r.extensions[name] = func(ctx *xtension.Context) (xtension.Xtension, error) {
		ctx.SetAdapterBuilder(name, rb)
		return nil, nil
	}

}

func (r *Registry) SetExtensionBuilder(name string, builder xtension.Builder) {
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

func (r *Registry) GetRepoBuilders() map[string]xpacman.Builder {
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

func (r *Registry) GetXtensionBuilder() map[string]xtension.Builder {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}

	return r.extensions
}
