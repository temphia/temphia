package admin

import (
	"time"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *Controller) AddUser(uclaim *claim.Session, user *entities.User) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	user.TenantID = uclaim.TenantId
	user.CreatedAt = time.Now()

	return c.coredb.AddUser(user, &entities.UserData{
		UserId:             user.UserId,
		MFAEnabled:         false,
		MFAType:            "",
		MFAData:            "",
		PendingPassChange:  true,
		PendingEmailVerify: false,
		ExtraMeta:          nil,
		TenantID:           uclaim.TenantId,
	})
}

func (c *Controller) UpdateUser(uclaim *claim.Session, user map[string]any) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.UpdateUser(uclaim.TenantId, uclaim.UserID, user)
}

func (c *Controller) RemoveUser(uclaim *claim.Session, username string) error {
	if !uclaim.IsSuperAdmin() {
		return easyerr.NotImpl()
	}

	return c.coredb.RemoveUser(uclaim.TenantId, username)
}

func (c *Controller) GetUserByID(uclaim *claim.Session, username string) (*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.GetUserByID(uclaim.TenantId, username)
}

func (c *Controller) GetUserByEmail(uclaim *claim.Session, email string) (*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.GetUserByEmail(uclaim.TenantId, email)
}

func (c *Controller) ListUsers(uclaim *claim.Session) ([]*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.ListUsers(uclaim.TenantId)
}

func (c *Controller) ListUsersByGroup(uclaim *claim.Session, group string) ([]*entities.User, error) {
	if !uclaim.IsSuperAdmin() {
		return nil, easyerr.NotImpl()
	}

	return c.coredb.ListUsersByGroup(uclaim.TenantId, group)
}

func (c *Controller) UpdateUserDevice(uclaim *claim.Session, user string, id int64, data map[string]any) error {
	return c.coredb.UpdateUserDevice(uclaim.TenantId, user, id, data)
}

func (c *Controller) GetUserDevice(uclaim *claim.Session, user string, id int64) (*entities.UserDevice, error) {
	return c.coredb.GetUserDevice(uclaim.TenantId, user, id)
}

func (c *Controller) ListUserDevice(uclaim *claim.Session, user string) ([]*entities.UserDevice, error) {
	return c.coredb.ListUserDevice(uclaim.TenantId, user)
}

func (c *Controller) RemoveUserDevice(uclaim *claim.Session, user string, id int64) error {
	return c.coredb.RemoveUserDevice(uclaim.TenantId, user, id)
}

type NewUserDevice struct {
	Name       string            `json:"name,omitempty" db:"name,omitempty"`
	UserId     string            `json:"user_id,omitempty" db:"user_id,omitempty"`
	DeviceType string            `json:"device_type,omitempty" db:"device_type,omitempty"`
	Scopes     string            `json:"scopes,omitempty" db:"scopes,omitempty"`
	ExtraMeta  map[string]string `json:"extra_meta,omitempty"`
}

func (c *Controller) AddUserDevice(uclaim *claim.Session, user string, data *NewUserDevice) error {

	// fixme => more user device user paring options
	// fixme => return id

	return c.coredb.AddUserDevice(uclaim.TenantId, user, &entities.UserDevice{
		Id:          0,
		Name:        data.Name,
		UserId:      user,
		DeviceType:  data.DeviceType,
		ExpiresOn:   time.Now().Add(time.Hour * 24 * 60),
		APNToken:    "",
		Scopes:      data.Scopes,
		ExtraMeta:   data.ExtraMeta,
		TenantID:    uclaim.TenantId,
		LastData:    entities.JsonStrMap{},
		PairOptions: entities.JsonStrMap{},
	})
}
