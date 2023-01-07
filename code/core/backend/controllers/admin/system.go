package admin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (c *Controller) ListSystemEvent(uclaim *claim.Session, etype string, last int64) ([]*entities.SystemEvent, error) {
	return c.coredb.ListSystemEvent(uclaim.TenentId, etype, last)
}

func (c *Controller) ListSystemKV(uclaim *claim.Session, ktype, prefix string, last int64) ([]*entities.SystemKV, error) {
	return c.coredb.ListSystemKV(uclaim.TenentId, ktype, prefix, last)
}
