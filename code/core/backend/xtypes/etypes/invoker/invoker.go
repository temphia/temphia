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
