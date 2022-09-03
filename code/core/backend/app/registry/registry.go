package registry

import (
	"errors"
	"sync"

	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/config"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type (
	StoreBuilder  func(*config.StoreSource) (store.Store, error)
	DynamicScript func(ns string, ctx interface{}) error
)

type Registry struct {
	repoBuilders        map[string]repox.Builder
	executors           map[string]etypes.BuilderFactory
	execModules         map[string]etypes.ModuleBuilderFunc
	dynamicScripts      map[string]DynamicScript
	storeBuilders       map[string]StoreBuilder
	httpAdapterBuilders map[string]httpx.Builder

	freezed bool
	mlock   *sync.Mutex
}

var (
	errTooLate = errors.New("err too late")
	errTooSoon = errors.New("err too soon")
)

func New(fromGlobal bool) *Registry {
	reg := &Registry{
		freezed:             false,
		dynamicScripts:      make(map[string]DynamicScript),
		repoBuilders:        make(map[string]repox.Builder),
		executors:           make(map[string]etypes.BuilderFactory),
		execModules:         make(map[string]etypes.ModuleBuilderFunc),
		storeBuilders:       make(map[string]StoreBuilder),
		httpAdapterBuilders: make(map[string]httpx.Builder),
		mlock:               &sync.Mutex{},
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

	for k, v := range G.executors {
		reg.executors[k] = v
	}
	for k, v := range G.execModules {
		reg.execModules[k] = v
	}
	for k, v := range G.dynamicScripts {
		reg.dynamicScripts[k] = v
	}

	for k, v := range G.httpAdapterBuilders {
		reg.httpAdapterBuilders[k] = v
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
	r.executors[name] = builder
}

func (r *Registry) SetExecModule(name string, builder etypes.ModuleBuilderFunc) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}
	r.execModules[name] = builder
}

func (r *Registry) SetDynamicScript(name string, script func(ns string, ctx interface{}) error) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}
	r.dynamicScripts[name] = script
}

func (r *Registry) SetAapterBuilder(name string, rb httpx.Builder) {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if r.freezed {
		panic(errTooLate)
	}

	r.httpAdapterBuilders[name] = rb
}

func (r *Registry) SetStoreBuilder(name string, b StoreBuilder) {
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

func (r *Registry) GetExecutors() map[string]etypes.BuilderFactory {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}
	return r.executors
}

func (r *Registry) GetExecModules() map[string]etypes.ModuleBuilderFunc {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}
	return r.execModules
}

func (r *Registry) GetDynamicScripts() map[string]DynamicScript {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}

	return r.dynamicScripts
}

func (r *Registry) GetStoreBuilders() map[string]StoreBuilder {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}

	return r.storeBuilders
}

func (r *Registry) GetAdapterBuilders() map[string]httpx.Builder {
	r.mlock.Lock()
	defer r.mlock.Unlock()
	if !r.freezed {
		panic(errTooSoon)
	}

	return r.httpAdapterBuilders
}
