package entities

import (
	"time"
)

type PlugKV struct {
	Key     string `json:"key,omitempty" db:"key"`
	Value   string `json:"value,omitempty" db:"value"`
	Version int64  `json:"version,omitempty" db:"Version"`

	PlugsID  string `json:"plug_id,omitempty" db:"plug_id"`
	TenantID string `json:"tenant_id,omitempty" db:"tenant_id"`

	// meta
	Tag1 string     `json:"tag1,omitempty" db:"tag1"`
	Tag2 string     `json:"tag2,omitempty" db:"tag2"`
	Tag3 string     `json:"tag3,omitempty" db:"tag3"`
	TTL  *time.Time `json:"ttl,omitempty" db:"ttl"`
}
