package admin

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

var (
	ErrNoTargetScope = easyerr.Error("scope not found [target]")
)

func (c *Controller) AddTargetHook(uclaim *claim.Session, data *entities.TargetHook) error {
	data.TenantId = uclaim.TenantId

	if !c.HasScope(uclaim, "target") {
		return ErrNoTargetScope
	}

	_, err := c.coredb.AddTargetHook(data)
	return err
}

func (c *Controller) UpdateTargetHook(uclaim *claim.Session, ttype string, id int64, data map[string]any) error {
	if !c.HasScope(uclaim, "target") {
		return ErrNoTargetScope
	}

	return c.coredb.UpdateTargetHook(uclaim.TenantId, ttype, id, data)
}

func (c *Controller) ListTargetHook(uclaim *claim.Session) ([]*entities.TargetHook, error) {
	if !c.HasScope(uclaim, "target") {
		return nil, ErrNoTargetScope
	}

	return c.coredb.ListTargetHook(uclaim.TenantId, nil)
}

func (c *Controller) ListTargetHookByType(uclaim *claim.Session, ttype, target string) ([]*entities.TargetHook, error) {
	if !c.HasScope(uclaim, "target") {
		return nil, ErrNoTargetScope
	}

	return c.coredb.ListTargetHookByType(uclaim.TenantId, ttype, target)
}

func (c *Controller) GetTargetHook(uclaim *claim.Session, ttype string, id int64) (*entities.TargetHook, error) {
	if !c.HasScope(uclaim, "target") {
		return nil, ErrNoTargetScope
	}

	return c.coredb.GetTargetHook(uclaim.TenantId, ttype, id)
}

func (c *Controller) RemoveTargetHook(uclaim *claim.Session, ttype string, id int64) error {
	if !c.HasScope(uclaim, "target") {
		return ErrNoTargetScope
	}

	return c.coredb.RemoveTargetHook(uclaim.TenantId, ttype, id)
}

func (c *Controller) AddTargetApp(uclaim *claim.Session, data *entities.TargetApp) error {
	if !c.HasScope(uclaim, "target") {
		return ErrNoTargetScope
	}

	data.TenantId = uclaim.TenantId
	_, err := c.coredb.AddTargetApp(data)
	return err
}

func (c *Controller) UpdateTargetApp(uclaim *claim.Session, ttype string, id int64, data map[string]any) error {
	if !c.HasScope(uclaim, "target") {
		return ErrNoTargetScope
	}

	return c.coredb.UpdateTargetApp(uclaim.TenantId, ttype, id, data)
}

func (c *Controller) ListTargetApp(uclaim *claim.Session) ([]*entities.TargetApp, error) {
	if !c.HasScope(uclaim, "target") {
		return nil, ErrNoTargetScope
	}

	return c.coredb.ListTargetApp(uclaim.TenantId, nil)
}

func (c *Controller) ListTargetAppByType(uclaim *claim.Session, ttype, target string) ([]*entities.TargetApp, error) {
	if !c.HasScope(uclaim, "target") {
		return nil, ErrNoTargetScope
	}

	return c.coredb.ListTargetAppByType(uclaim.TenantId, ttype, target)
}

func (c *Controller) GetTargetApp(uclaim *claim.Session, ttype string, id int64) (*entities.TargetApp, error) {
	if !c.HasScope(uclaim, "target") {
		return nil, ErrNoTargetScope
	}

	return c.coredb.GetTargetApp(uclaim.TenantId, ttype, id)
}

func (c *Controller) RemoveTargetApp(uclaim *claim.Session, ttype string, id int64) error {
	if !c.HasScope(uclaim, "target") {
		return ErrNoTargetScope
	}

	return c.coredb.RemoveTargetApp(uclaim.TenantId, ttype, id)
}
