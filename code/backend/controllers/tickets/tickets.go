package tickets

import "github.com/temphia/temphia/code/backend/xtypes/store"

type Controller struct {
	plugState  store.PlugStateKV
	corehub    store.CoreHub
	cabinethub store.CabinetHub
}

func New(corehub store.CoreHub, plugState store.PlugStateKV) *Controller {
	return &Controller{
		plugState: plugState,
		corehub:   corehub,
	}
}
