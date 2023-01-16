package admin

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *Controller) ListSystemEvent(uclaim *claim.Session, etype string, last int64) ([]*entities.SystemEvent, error) {
	return c.coredb.ListSystemEvent(uclaim.TenantId, etype, last)
}

func (c *Controller) ListSystemKV(uclaim *claim.Session, ktype, prefix string, last int64) ([]*entities.SystemKV, error) {
	return c.coredb.ListSystemKV(uclaim.TenantId, ktype, prefix, last)
}
