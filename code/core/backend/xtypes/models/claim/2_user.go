package claim

import (
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/core/backend/xtypes"
)

type User struct {
	TenentId   string            `json:"tenent_id,omitempty"`
	UserID     string            `json:"user_id,omitempty"`
	UserGroup  string            `json:"user_group,omitempty"`
	Type       string            `json:"type,omitempty"`
	DeviceId   string            `json:"device_id,omitempty"`
	Scopes     []string          `json:"scopes,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

func NewUserDevice(tenantId, userId, groupId string, scopes []string) *User {
	return &User{
		TenentId:   tenantId,
		Type:       "user_device",
		UserID:     userId,
		UserGroup:  groupId,
		DeviceId:   xid.New().String(),
		Scopes:     scopes,
		Attributes: make(map[string]string),
	}
}

func NewUserLogged(tenantId, userId, groupId string, scopes []string) *User {
	return &User{
		TenentId:   tenantId,
		Type:       "user_logged",
		UserID:     userId,
		UserGroup:  groupId,
		DeviceId:   xid.New().String(),
		Scopes:     scopes,
		Attributes: make(map[string]string),
	}
}

func (u *User) IsSuperAdmin() bool {
	return u.UserGroup == xtypes.UserGroupSuperAdmin
}

func (u *User) IsGuest() bool {
	return u.UserGroup == xtypes.UserGroupGuest
}

// session

type Session struct {
	TenentId   string            `json:"-"`
	UserID     string            `json:"user,omitempty"`
	UserGroup  string            `json:"group,omitempty"`
	Type       string            `json:"type,omitempty"`
	SessionID  int64             `json:"session_id,omitempty"`
	DeviceId   string            `json:"device_id,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

func NewSession(user, group, device string, sid int64) *Session {
	return &Session{
		TenentId:   "",
		UserID:     user,
		UserGroup:  group,
		Type:       "session",
		SessionID:  sid,
		DeviceId:   device,
		Attributes: make(map[string]string),
	}
}

func (p *Session) SetAttr(key, value string) {
	p.Attributes[key] = value
}

func (u *Session) IsSuperAdmin() bool {
	return u.UserGroup == xtypes.UserGroupSuperAdmin
}

func (u *Session) IsGuest() bool {
	return u.UserGroup == xtypes.UserGroupGuest
}
