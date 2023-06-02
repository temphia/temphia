package bindx

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx/ticket"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
)

type Bindings interface {
	Core

	PlugKVBindingsGet() any
	SockdBindingsGet() Sockd
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

type Sockd interface {
	SendDirect(room string, connId int64, payload []byte) error
	SendDirectBatch(room string, conns []int64, payload []byte) error
	SendBroadcast(room string, ignores []int64, payload []byte) error
	SendTagged(room string, tags []string, ignores []int64, payload []byte) error
	RoomUpdateTags(room string, opts sockdx.UpdateTagOptions) error

	Ticket(room string, opts *ticket.SockdRoom) (string, error)
}

type Self interface {
	SelfListResources() ([]*Resource, error)
	SelfGetResource(name string) (*Resource, error)

	SelfInLinks() ([]Link, error)
	SelfOutLinks() ([]Link, error)

	SelfLinkExec(name, method string, data xtypes.LazyData, async, detached bool) (xtypes.LazyData, error)
	SelfNewModule(name string, data xtypes.LazyData) (int32, error)
	SelfModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error)

	SelfForkExec(method string, data []byte) error
}

type Invoker interface {
	Name() string
	ExecMethod(method, path string, data xtypes.LazyData) (xtypes.LazyData, error)
	ContextUser() *claim.UserContext
	ContextUserInfo() (*entities.UserInfo, error)
	ContextUserMessage(opts *UserMessage) error
}
