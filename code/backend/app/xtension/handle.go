package xtension

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/adapter"
)

type Context struct {
	data Data
	App  xtypes.App
}

func NewHandle() *Context {
	return &Context{
		data: Data{
			Executors: make(map[string]etypes.ExecutorBuilder),
			Modules:   make(map[string]etypes.ModuleBuilder),
			Adapters:  make(map[string]adapter.Builder),
			Scripts:   make(map[string]func(tenantId string, ctx any) error),
		},
	}
}

func (e *Context) SetExecutorBuilder(name string, builder etypes.ExecutorBuilder) {
	e.data.Executors[name] = builder
}

func (e *Context) SetModuleBuilder(name string, builder etypes.ModuleBuilder) {
	e.data.Modules[name] = builder
}

func (e *Context) SetAdapterBuilder(name string, builder adapter.Builder) {
	e.data.Adapters[name] = builder
}

func (e *Context) SetScript(name string, s func(tenantId string, ctx any) error) {
	e.data.Scripts[name] = s
}

type Data struct {
	Executors map[string]etypes.ExecutorBuilder
	Modules   map[string]etypes.ModuleBuilder
	Adapters  map[string]adapter.Builder
	Scripts   map[string]func(tenantId string, ctx any) error
}

func GetContextData(h *Context) Data {
	return h.data
}
