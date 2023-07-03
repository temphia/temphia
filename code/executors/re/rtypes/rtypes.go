package rtypes

import "encoding/json"

type Packet struct {
	Id   string
	Type string // proto.go
	Data json.RawMessage
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
