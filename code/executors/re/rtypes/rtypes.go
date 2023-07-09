package rtypes

import "encoding/json"

type Packet struct {
	Id   string          `json:"id,omitempty"`
	Name string          `json:"name,omitempty"`
	Type string          `json:"type,omitempty"` // proto.go
	Data json.RawMessage `json:"data,omitempty"`
}

type BootstrapContext struct {
	Folder   string
	TenantId string
	PlugId   string
	AgentId  string
	File     string
	GetFile  func(name string) ([]byte, error)
}

type BootstrapFunc func(ctx BootstrapContext) error
