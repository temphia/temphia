package app

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
)

type extHandle struct {
	executors map[string]etypes.ExecutorBuilder
	modules   map[string]etypes.ModuleBuilder
	adapters  map[string]httpx.Builder
	scripts   map[string]func(tenantId string, ctx any) error
}

func newHandle() *extHandle {
	return &extHandle{
		executors: make(map[string]etypes.ExecutorBuilder),
		modules:   make(map[string]etypes.ModuleBuilder),
		adapters:  make(map[string]httpx.Builder),
		scripts:   make(map[string]func(tenantId string, ctx any) error),
	}
}

func (e *extHandle) SetExecutorBuilder(name string, builder etypes.ExecutorBuilder) {
	e.executors[name] = builder
}

func (e *extHandle) SetModuleBuilder(name string, builder etypes.ModuleBuilder) {
	e.modules[name] = builder
}

func (e *extHandle) SetAdapterBuilder(name string, builder httpx.Builder) {
	e.adapters[name] = builder
}

func (e *extHandle) SetScript(name string, s func(tenantId string, ctx any) error) {
	e.scripts[name] = s
}

func (e *extHandle) SetOnStart(hook func()) {}
func (e *extHandle) SetOnExit(hook func())  {}

func (e *extHandle) GetExecutorBuilder() map[string]etypes.ExecutorBuilder      { return e.executors }
func (e *extHandle) GetModuleBuilder() map[string]etypes.ModuleBuilder          { return e.modules }
func (e *extHandle) GetAdapterBuilder() map[string]httpx.Builder                { return e.adapters }
func (e *extHandle) GetScript() map[string]func(tenantId string, ctx any) error { return e.scripts }
