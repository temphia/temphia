package hubv2

import "encoding/json"

type Step struct {
	Name string          `json:"name,omitempty" yaml:"name,omitempty"`
	Type string          `json:"type,omitempty" yaml:"type,omitempty"`
	Data json.RawMessage `json:"data,omitempty" yaml:"data,omitempty"`
}
