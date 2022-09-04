package dtable

import (
	"github.com/temphia/temphia/code/core/backend/engine/invoker"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type ServerSideHook struct {
	dhub store.DynSource
	hook *entities.DataHook
	// userclaim       *claim.Session
	// allowHqlQuery   bool
	// allowRawQuery   bool
	// allowedSiblings []string

}

func NewServerSideHook(dhub store.DynSource, hook *entities.DataHook) *ServerSideHook {
	return &ServerSideHook{
		dhub: dhub,
		hook: hook,
	}
}

func (r *ServerSideHook) Handle(method string, data xtypes.LazyData) (xtypes.LazyData, error) {

	return nil, nil
}

func (r *ServerSideHook) Name() string {
	return invoker.TypeDtableServerHook
}
