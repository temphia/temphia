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
}

type Builder func(App xtypes.App) (Instancer, error)

type Instancer interface {
	Instance(opts Options) (any, error)
}
