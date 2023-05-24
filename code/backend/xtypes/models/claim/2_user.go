package claim

import (
	"github.com/temphia/temphia/code/backend/xtypes"
)

type User struct {
	TenantId   string            `json:"tenent_id,omitempty"`
	UserID     string            `json:"user_id,omitempty"`
	UserGroup  string            `json:"user_group,omitempty"`
	Type       string            `json:"type,omitempty"`
	DeviceId   int64             `json:"device_id,omitempty"`
	Scopes     []string          `json:"scopes,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

func NewUserDevice(tenantId, userId, groupId string, scopes []string) *User {
	return &User{
		TenantId:   tenantId,
		Type:       "device",
		UserID:     userId,
		UserGroup:  groupId,
		DeviceId:   0,
		Scopes:     scopes,
		Attributes: make(map[string]string),
	}
}

func NewUserLogged(tenantId, userId, groupId string, device int64, scopes []string) *User {
	return &User{
		TenantId:   tenantId,
		Type:       "logged",
		UserID:     userId,
		UserGroup:  groupId,
		DeviceId:   device,
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

func (u *User) DeriveSession(sid int64) *Session {

	return &Session{
		TenantId:   u.TenantId,
		UserID:     u.UserID,
		UserGroup:  u.UserGroup,
		Type:       "session",
		SessionID:  sid,
		DeviceId:   u.DeviceId,
		Attributes: make(map[string]string),
	}

}

// session

type Session struct {
	TenantId   string            `json:"-"`
	UserID     string            `json:"user,omitempty"`
	UserGroup  string            `json:"group,omitempty"`
	Type       string            `json:"type,omitempty"`
	SessionID  int64             `json:"session_id,omitempty"`
	DeviceId   int64             `json:"device_id,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
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

func (u *Session) AsUserCtx() *UserContext {
	return &UserContext{
		TenantId:  u.TenantId,
		UserID:    u.UserID,
		UserGroup: u.UserGroup,
		SessionID: u.SessionID,
		DeviceId:  u.DeviceId,
	}
}
