package basic

import (
	"time"

	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type NewUserDevice struct {
	Name       string            `json:"name,omitempty"`
	DeviceType string            `json:"device_type,omitempty"`
	APNToken   string            `json:"apn_token,omitempty"`
	Scopes     string            `json:"scopes,omitempty"`
	ExtraMeta  map[string]string `json:"extra_meta,omitempty"`
}

func (c *Controller) AddUserDevice(uclaim *claim.Session, data *NewUserDevice) error {
	// fixme => more paring options
	// fixme => return new id

	return c.coredb.AddUserDevice(uclaim.TenantId, uclaim.UserID, &entities.UserDevice{
		Id:          0,
		Name:        data.Name,
		UserId:      uclaim.UserID,
		DeviceType:  data.DeviceType,
		APNToken:    data.APNToken,
		Scopes:      data.Scopes,
		ExtraMeta:   data.ExtraMeta,
		TenantID:    uclaim.TenantId,
		LastData:    entities.JsonStrMap{},
		PairOptions: entities.JsonStrMap{},
		ExpiresOn:   time.Now().Add(time.Hour * 24 * 60),
	})

}

func (c *Controller) UpdateUserDevice(uclaim *claim.Session, id int64, data map[string]any) error {
	return c.coredb.UpdateUserDevice(uclaim.TenantId, uclaim.UserID, id, data)
}

func (c *Controller) GetUserDevice(uclaim *claim.Session, id int64) (*entities.UserDevice, error) {
	return c.coredb.GetUserDevice(uclaim.TenantId, uclaim.UserID, id)
}

func (c *Controller) ListUserDevice(uclaim *claim.Session) ([]*entities.UserDevice, error) {
	return c.coredb.ListUserDevice(uclaim.TenantId, uclaim.UserID)
}

func (c *Controller) RemoveUserDevice(uclaim *claim.Session, id int64) error {
	return c.coredb.RemoveUserDevice(uclaim.TenantId, uclaim.UserID, id)
}
