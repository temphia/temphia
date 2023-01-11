package entities

import "time"

type UserMessageReq struct {
	Cursor int64  `json:"cursor,omitempty"`
	Count  int64  `json:"count,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

type UserMessage struct {
	Id           int        `json:"id,omitempty" db:"id,omitempty"`
	Title        string     `json:"title,omitempty" db:"title,omitempty"`
	Read         bool       `json:"read,omitempty" db:"read,omitempty"`
	Type         string     `json:"type,omitempty" db:"type"`
	Contents     string     `json:"contents,omitempty" db:"contents"`
	UserId       string     `json:"user_id,omitempty" db:"user_id"`
	FromUser     string     `json:"from_user,omitempty" db:"from_user,omitempty,"`
	FromPlug     string     `json:"from_plug,omitempty" db:"from_plug,omitempty,"`
	FromAgent    string     `json:"from_agent,omitempty" db:"from_agent,omitempty,"`
	PlugCallback string     `json:"plug_callback,omitempty" db:"plug_callback,omitempty"`
	WarnLevel    int        `json:"warn_level,omitempty" db:"warn_level,omitempty"`
	Encrypted    bool       `json:"encrypted,omitempty" db:"encrypted,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	TenantId     string     `json:"tenant_id,omitempty" db:"tenant_id"`
}

type ModifyMessages struct {
	Operation string  `json:"ops,omitempty"`
	Ids       []int64 `json:"ids,omitempty"`
}
