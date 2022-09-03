package instancer

import "github.com/temphia/temphia/code/core/backend/xtypes"

type Options struct {
	TenantId     string
	Bid          string
	InstanceType string
	File         string
	UserId       string
	Data         []byte
}

type InstancerBuilder func(App xtypes.App) (Instancer, error)

type Instancer interface {
	Instance(opts Options) (interface{}, error)
}
