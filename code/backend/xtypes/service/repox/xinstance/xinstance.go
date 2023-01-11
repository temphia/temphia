package xinstance

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
)

type Options struct {
	TenantId     string
	BprintId     string
	InstanceType string
	File         string
	UserId       string
	UserData     []byte
	Automatic    bool
	Handle       Handle
}

type Response struct {
	Ok      bool   `json:"ok,omitempty"`
	Message string `json:"message,omitempty"`
	Slug    string `json:"slug,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type Handle interface {
	GetFile(file string) ([]byte, error)
	LoadFile(file string, target any) error // loads json/yaml
	GetPrevObject(name string) *Response
}

type Builder func(App xtypes.App) (Instancer, error)

type Instancer interface {
	Instance(opts Options) (*Response, error)
}
