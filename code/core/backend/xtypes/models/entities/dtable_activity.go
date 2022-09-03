package entities

import "time"

type DynActivity struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Type      string     `json:"type,omitempty" db:"type"`
	RowId     int64      `json:"row_id,omitempty" db:"row_id"`
	RowVerson int64      `json:"row_version,omitempty" db:"row_version"`
	InitSign  string     `json:"init_sign,omitempty" db:"init_sign,omitempty"`
	UserId    string     `json:"user_id,omitempty" db:"user_id,omitempty"`
	UserSign  string     `json:"user_sign,omitempty" db:"user_sign,omitempty"`
	Payload   string     `json:"payload,omitempty" db:"payload,omitempty"`
	Message   string     `json:"message,omitempty" db:"message,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
}

type ActivityQuery struct {
	Types       []string  `json:"types,omitempty"`
	UserId      string    `json:"user_id,omitempty"`
	BetweenTime [2]string `json:"between_time,omitempty"`
	Count       int64     `json:"count,omitempty"`
	Offset      string    `json:"offset,omitempty"`
}
