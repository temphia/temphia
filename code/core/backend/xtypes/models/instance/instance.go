package instance

import (
	"encoding/json"
)

type RepoOptions struct {
	BprintId      string          `json:"bprint_id,omitempty"`
	InstancerType string          `json:"instancer_type,omitempty"`
	Data          json.RawMessage `json:"data,omitempty"`
	File          string          `json:"file,omitempty"`
	UserId        string          `json:"-"`
}
