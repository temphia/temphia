package invokers

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/invoker"
)

type Invoker struct {
	name    string
	app     xtypes.App
	modules map[string]Module
	arrts   map[string]any
}

func (i *Invoker) Type() string { return i.name }

func (i *Invoker) ExecuteModule(module, action string, data xtypes.LazyData) (xtypes.LazyData, error) {
	mod := i.modules[module]
	if mod == nil {
		panic("invoker module not found")
	}

	return mod(Handle{
		invoker: i,
	}, action, data)
}

func (i *Invoker) ListModules() []string {
	ms := make([]string, 0, len(i.modules))
	for k := range i.modules {
		ms = append(ms, k)
	}
	return ms
}

func (i *Invoker) UserContext() *invoker.User {
	return nil
}

func (i *Invoker) GetAttr(name string) any {
	if i.arrts == nil {
		return nil
	}

	return i.arrts[name]
}

func (i *Invoker) GetAttrs() map[string]any {
	return i.arrts
}
