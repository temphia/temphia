package entities

import (
	"strings"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

type UserDevice struct {
	Id         int64      `json:"id,omitempty" db:"id,omitempty"`
	Name       string     `json:"name,omitempty" db:"name,omitempty"`
	UserId     string     `json:"user_id,omitempty" db:"user_id,omitempty"`
	DeviceType string     `json:"device_type,omitempty" db:"device_type,omitempty"`
	LastAddr   string     `json:"last_addr,omitempty" db:"last_addr,omitempty"`
	APNToken   string     `json:"apn_token,omitempty" db:"apn_token,omitempty"`
	Scopes     string     `json:"scopes,omitempty" db:"scopes,omitempty"`
	ExtraMeta  JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantID   string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
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
