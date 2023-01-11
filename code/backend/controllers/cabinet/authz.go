package cabinet

import (
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

func (c *Controller) canAction(uclaim *claim.Session, target, action string) error {
	if uclaim.IsSuperAdmin() {
		return nil
	}

	err := c.scopeCheck(uclaim, target, action)
	if err != nil {
		return err
	}

	// if uclaim.SkipRBAC {
	// 	return nil
	// }

	return c.permCheck(uclaim, target, action)
}

func (c *Controller) scopeCheck(uclaim *claim.Session, target, action string) error {

	return easyerr.NotAuthorized()
}

func (c *Controller) permCheck(uclaim *claim.Session, target, action string) error {
	return easyerr.NotImpl()
}
