package entities

import (
	"strings"
	"time"

	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type UserDevice struct {
	Id          int64      `json:"id,omitempty" db:"id,omitempty"`
	Name        string     `json:"name,omitempty" db:"name,omitempty"`
	UserId      string     `json:"user_id,omitempty" db:"user_id,omitempty"`
	DeviceType  string     `json:"device_type,omitempty" db:"device_type,omitempty"`
	APNToken    string     `json:"apn_token,omitempty" db:"apn_token,omitempty"`
	Scopes      string     `json:"scopes,omitempty" db:"scopes,omitempty"`
	LastData    JsonStrMap `json:"last_data,omitempty" db:"last_data,omitempty"`
	PairOptions JsonStrMap `json:"pair_options,omitempty" db:"pair_options,omitempty"`
	ExtraMeta   JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	ExpiresOn   time.Time  `json:"expires_on,omitempty" db:"expires_on"`
	TenantID    string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

func (ud *UserDevice) Derive(ugroup string) *claim.User {
	return &claim.User{
		TenentId:   ud.TenantID,
		UserID:     ud.UserId,
		UserGroup:  ugroup,
		Type:       ud.DeviceType,
		DeviceId:   ud.Id,
		Scopes:     strings.Split(ud.Scopes, ","),
		Attributes: make(map[string]string),
	}

}
