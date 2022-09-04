package admin

import (
	"time"

	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (c *Controller) AddUser(uclaim *claim.Session, user *entities.User) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	user.TenantID = uclaim.TenentId
	user.CreatedAt = time.Now()

	return c.coredb.AddUser(user, &entities.UserData{
		UserId:             user.UserId,
		MFAEnabled:         false,
		MFAType:            "",
		MFAData:            "",
		PendingPassChange:  true,
		PendingEmailVerify: false,
		ExtraMeta:          nil,
		TenantID:           uclaim.TenentId,
	})
}

func (c *Controller) UpdateUser(uclaim *claim.Session, user map[string]any) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.UpdateUser(uclaim.TenentId, uclaim.UserID, user)
}

func (c *Controller) RemoveUser(uclaim *claim.Session, username string) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.RemoveUser(uclaim.TenentId, username)
}

func (c *Controller) GetUserByID(uclaim *claim.Session, username string) (*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.GetUserByID(uclaim.TenentId, username)
}

func (c *Controller) GetUserByEmail(uclaim *claim.Session, email string) (*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.GetUserByEmail(uclaim.TenentId, email)
}

func (c *Controller) ListUsers(uclaim *claim.Session) ([]*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.ListUsers(uclaim.TenentId)
}

func (c *Controller) ListUsersByGroup(uclaim *claim.Session, group string) ([]*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.ListUsersByGroup(uclaim.TenentId, group)
}
