package admin

import (
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (c *Controller) AddPerm(sess *claim.Session, data *entities.Permission) error {
	if !sess.IsSuperAdmin() {
		return easyerr.NotAuthorized()
	}

	return nil
}

func (c *Controller) UpdatePerm(sess *claim.Session, data map[string]any) error {
	if !sess.IsSuperAdmin() {
		return easyerr.NotAuthorized()
	}

	return nil
}
func (c *Controller) GetPerm(sess *claim.Session, id int64) (*entities.Permission, error) {
	if !sess.IsSuperAdmin() {
		return nil, easyerr.NotAuthorized()
	}

	return nil, nil
}

func (c *Controller) RemovePerm(sess *claim.Session, id int64) error {
	if !sess.IsSuperAdmin() {
		return easyerr.NotAuthorized()
	}

	return nil
}
func (c *Controller) AddRole(sess *claim.Session, data *entities.Role) error {
	if !sess.IsSuperAdmin() {
		return easyerr.NotAuthorized()
	}

	return nil
}

func (c *Controller) GetRole(sess *claim.Session, tenantId string, id int64) (*entities.Role, error) {
	if !sess.IsSuperAdmin() {
		return nil, easyerr.NotAuthorized()
	}
	return nil, nil
}

func (c *Controller) UpdateRole(sess *claim.Session, data map[string]any) error {
	if !sess.IsSuperAdmin() {
		return easyerr.NotAuthorized()
	}
	return nil
}

func (c *Controller) RemoveRole(sess *claim.Session, data *entities.Role) error {
	if !sess.IsSuperAdmin() {
		return easyerr.NotAuthorized()
	}
	return nil
}

func (c *Controller) AddUserRole(sess *claim.Session, data *entities.UserRole) error {
	if !sess.IsSuperAdmin() {
		return easyerr.NotAuthorized()
	}

	return nil
}

func (c *Controller) GetUserRole(sess *claim.Session, tenantId string, id int64) (*entities.UserRole, error) {
	if !sess.IsSuperAdmin() {
		return nil, easyerr.NotAuthorized()
	}

	return nil, nil
}
func (c *Controller) UpdateUserRole(sess *claim.Session, data map[string]any) error {
	if !sess.IsSuperAdmin() {
		return easyerr.NotAuthorized()
	}

	return nil
}

func (c *Controller) RemoveUserRole(sess *claim.Session, data *entities.UserRole) error {
	if !sess.IsSuperAdmin() {
		return easyerr.NotAuthorized()
	}

	return nil
}

func (c *Controller) ListAllPerm(sess *claim.Session, tenantId string) ([]*entities.Permission, error) {
	if !sess.IsSuperAdmin() {
		return nil, easyerr.NotAuthorized()
	}

	return nil, nil
}

func (c *Controller) ListAllRole(sess *claim.Session, tenantId string) ([]*entities.Permission, error) {
	if !sess.IsSuperAdmin() {
		return nil, easyerr.NotAuthorized()
	}

	return nil, nil
}

func (c *Controller) ListAllUserRole(sess *claim.Session, tenantId string) ([]*entities.Permission, error) {
	if !sess.IsSuperAdmin() {
		return nil, easyerr.NotAuthorized()
	}

	return nil, nil
}

func (c *Controller) ListAllUserPerm(sess *claim.Session, tenantId string) ([]*entities.Permission, error) {
	if !sess.IsSuperAdmin() {
		return nil, easyerr.NotAuthorized()
	}

	return nil, nil
}

func (c *Controller) ListUserPerm(sess *claim.Session, tenantId string, userId, objType, objsub string) ([]*entities.Permission, error) {
	if !sess.IsSuperAdmin() {
		return nil, easyerr.NotAuthorized()
	}

	return nil, nil

}
