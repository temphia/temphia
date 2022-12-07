package invoker

import "github.com/temphia/temphia/code/core/backend/xtypes"

type User struct {
	Id        string
	Group     string
	SessionId int64
	DeviceId  string
}

type Invoker interface {
	User() *User
	Name() string
	Module
}

type Module interface {
	Handle(method string, data xtypes.LazyData) (xtypes.LazyData, error)
}

// new
type Invoker2 interface {
	Type() string
	ExecuteModule(module, action string, data xtypes.LazyData) (xtypes.LazyData, error)
	ListModules() []string
	UserContext() *User

	GetAttr(string) interface{}
	GetAttrs() map[string]interface{}
}
