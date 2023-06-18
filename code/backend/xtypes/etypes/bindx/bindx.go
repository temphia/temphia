package bindx

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type Bindings interface {
	Core

	SelfBindingsGet() Self
	InvokerGet() Invoker
}

type Core interface {
	Log(msg string)
	LazyLog(msgs []string)
	Sleep(int32)
	GetFileWithMeta(file string) (data []byte, version int64, err error)
	GetApp() any
}

type Self interface {
	ListResources() ([]*Resource, error)
	GetResource(name string) (*Resource, error)

	InLinks() ([]Link, error)
	OutLinks() ([]Link, error)

	LinkExec(name, method string, data xtypes.LazyData) (xtypes.LazyData, error)
	LinkExecEmit(name, method string, data xtypes.LazyData) error

	NewModule(name string, data xtypes.LazyData) (int32, error)
	ModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error)
	ModuleTicket(name string, opts xtypes.LazyData) (string, error)

	ForkExec(method string, data []byte) error

	// if executor doesnot have native async support then you could use these for asyncness

	AsyncLinkExec(name, method string, data xtypes.LazyData) (uint32, error)
	AsyncModuleExec(mid int32, method string, data xtypes.LazyData) (uint32, error)
	AsyncEventPoll(mid int32, eid uint32) (xtypes.LazyData, error)
	AsyncEventWait(mid int32, eid uint32) (xtypes.LazyData, error)
}

type Invoker interface {
	Name() string
	ExecMethod(method string, data xtypes.LazyData) (xtypes.LazyData, error)
	UserContext() *claim.UserContext
	UserInfo() (*entities.UserInfo, error)
	UserMessage(opts *UserMessage) error
}
