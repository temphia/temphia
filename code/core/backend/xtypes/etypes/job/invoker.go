package job

import "github.com/temphia/temphia/code/core/backend/xtypes"

type InvokeUser struct {
	UserId    string
	UserGroup string
	SessionId int64
	DeviceId  string
}

type Invoker interface {
	Name() string
	CurrentUser() *InvokeUser
	Handle(method string, data xtypes.LazyData) (xtypes.LazyData, error)
}
