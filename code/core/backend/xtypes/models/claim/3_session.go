package claim

import "github.com/temphia/temphia/code/core/backend/xtypes"

type Session struct {
	TenentId   string            `json:"-"`
	UserID     string            `json:"user,omitempty"`
	UserGroup  string            `json:"group,omitempty"`
	Type       string            `json:"type,omitempty"`
	SessionID  int64             `json:"session_id,omitempty"`
	DeviceId   string            `json:"device_id,omitempty"`
	Path       []string          `json:"path,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

func NewSession(user, group, device string, sid int64, path ...string) *Session {
	return &Session{
		TenentId:   "",
		UserID:     user,
		UserGroup:  group,
		Type:       "session",
		SessionID:  sid,
		DeviceId:   device,
		Path:       path,
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
