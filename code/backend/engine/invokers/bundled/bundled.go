package bundled

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
)

type Invoker struct {
	name             string
	app              xtypes.App
	modules          map[string]Module
	arrts            map[string]any
	get_user_context func() *invoker.User
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
	if i.get_user_context == nil {
		return nil
	}

	return i.get_user_context()
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

// handle

type Handle struct {
	invoker *Invoker
}

func (h *Handle) ExecuteModule(module, action string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return h.invoker.ExecuteModule(module, action, data)
}

func (h *Handle) ListModule() []string {
	return h.invoker.ListModules()
}

func (h *Handle) GetApp() interface{} {
	return h.invoker.app
}
