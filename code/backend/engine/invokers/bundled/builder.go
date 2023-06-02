package bundled

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type Module func(method, path string, data xtypes.LazyData) (xtypes.LazyData, error)

type Builder struct {
	name             string
	modules          map[string]Module
	app              xtypes.App
	attrs            map[string]any
	get_user_context func() *claim.UserContext
}

func NewBuilder(name string) *Builder {
	return &Builder{
		name:    name,
		modules: make(map[string]Module),
	}
}

func (b *Builder) SetAttrs(attrs map[string]any) {
	b.attrs = attrs
}

func (b *Builder) SetApp(app xtypes.App) {
	b.app = app
}

func (b *Builder) SetModule(name string, mod Module) {
	b.modules[name] = mod
}

func (b *Builder) SetUserContextProvider(fn func() *claim.UserContext) {
	b.get_user_context = fn
}

func (b *Builder) SetModules(map[string]Module) {
	for k, v := range b.modules {
		b.modules[k] = v
	}
}

func (b *Builder) Build() *Invoker {
	return &Invoker{
		name:             b.name,
		modules:          b.modules,
		app:              b.app,
		arrts:            b.attrs,
		get_user_context: b.get_user_context,
	}
}
