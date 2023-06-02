package bundled

import (
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type Invoker struct {
	name             string
	app              xtypes.App
	modules          map[string]Module
	arrts            map[string]any
	get_user_context func() *claim.UserContext
}

func (i *Invoker) Type() string { return i.name }

func (i *Invoker) ExecuteMethod(method, path string, data xtypes.LazyData) (xtypes.LazyData, error) {

	mpath := strings.Split(method, ".")

	mod := i.modules[mpath[0]]
	if mod == nil {
		panic("invoker module not found")
	}

	return mod(mpath[1], path, data)
}

func (i *Invoker) UserContext() *claim.UserContext {
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
