package bindx

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type Core interface {
	Log(msg string)
	LazyLog(msgs []string)
	Sleep(int32)
	GetFileWithMeta(file string) (data []byte, version int64, err error)

	ListResources() ([]*Resource, error)
	GetResource(name string) (*Resource, error)

	InLinks() ([]Link, error)
	OutLinks() ([]Link, error)

	LinkExec(name, method string, data xtypes.LazyData) (xtypes.LazyData, error)
	LinkExecEmit(name, method string, data xtypes.LazyData) error

	NewModule(name string, data xtypes.LazyData) (int32, error)
	ModuleTicket(name string, opts xtypes.LazyData) (string, error)
	ModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error)
	ForkExec(method string, data []byte) error

	HttpFetch(*HttpRequest) *HttpResponse

	// if executor doesnot have native async support then you could use these for asyncness

	AsyncLinkExec(name, method string, data xtypes.LazyData) (uint32, error)
	AsyncModuleExec(mid int32, method string, data xtypes.LazyData) (uint32, error)
	AsyncEventPoll(mid int32, eid uint32) (xtypes.LazyData, error)
	AsyncEventWait(mid int32, eid uint32) (xtypes.LazyData, error)

	UserContext() *claim.UserContext
}

type Bindings interface {
	Core

	// clone makes copy with some state (eg. Job) stripped so it can be stored for longer
	// duration(than usual req/resp event cycle)
	Clone() Core

	GetInvoker() Invoker
}

type Invoker interface {
	Name() string
	ExecMethod(method string, data xtypes.LazyData) (xtypes.LazyData, error)
	UserContext() *claim.UserContext
	UserInfo() (*entities.UserInfo, error)
	UserMessage(opts *UserMessage) error
}
