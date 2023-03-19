package forked

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
)

type Forked struct {
	forkedFrom string
}

func New(from string) invoker.Invoker {
	return &Forked{
		forkedFrom: from,
	}
}

func (f *Forked) Type() string { return "forked" }
func (f *Forked) ExecuteMethod(module, action string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return nil, easyerr.NotImpl()
}

func (f *Forked) UserContext() *invoker.User {
	return nil
}

func (f *Forked) GetAttr(string) interface{} {
	return nil
}

func (f *Forked) GetAttrs() map[string]interface{} {
	return map[string]interface{}{}
}
