package bindx

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Bindings interface {
	Core

	PlugKVBindingsGet() PlugKV
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
}

type Self interface {
	SelfListResources() ([]*Resource, error)
	SelfGetResource(name string) (*Resource, error)

	SelfInLinks() ([]Link, error)
	SelfOutLinks() ([]Link, error)

	SelfLinkExec(name, method string, data xtypes.LazyData, async, detached bool) (xtypes.LazyData, error)
	SelfNewModule(name string, data xtypes.LazyData) (int32, error)
	SelfModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error)
	SelfModuleTicket(name string, opts xtypes.LazyData) (string, error)

	SelfForkExec(method string, data []byte) error
}

type PlugKV interface {
	Set(txid uint32, key, value string, opts *store.SetOptions) error
	Update(txid uint32, key, value string, opts *store.UpdateOptions) error
	Get(txid uint32, key string) (*entities.PlugKV, error)
	Del(txid uint32, key string) error
	DelBatch(txid uint32, keys []string) error
	Query(txid uint32, query *store.PkvQuery) ([]*entities.PlugKV, error)

	NewTxn() (uint32, error)
	RollBack(txid uint32) error
	Commit(txid uint32) error
}

type Invoker interface {
	Name() string
	ExecMethod(method, path string, data xtypes.LazyData) (xtypes.LazyData, error)
	ContextUser() *claim.UserContext
	ContextUserInfo() (*entities.UserInfo, error)
	ContextUserMessage(opts *UserMessage) error
}
