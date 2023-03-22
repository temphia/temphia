package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func (c *Controller) ListSystemEvent(uclaim *claim.Session, etype string, last int64) ([]*entities.SystemEvent, error) {
	return c.coredb.QuerySystemEvent(store.EventQuery{
		TenantId: uclaim.TenantId,
		Etype:    etype,
		Last:     last,
	})
}

func (c *Controller) ListSystemKV(uclaim *claim.Session, ktype, prefix string, last int64) ([]*entities.SystemKV, error) {
	return c.coredb.ListSystemKV(uclaim.TenantId, ktype, prefix, last)
}
