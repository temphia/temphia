package bindx

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Bindings interface {
	Core

	PlugKVBindingsGet() PlugKV
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
	ExecMethod(method string, data xtypes.LazyData) (xtypes.LazyData, error)
	UserContext() *claim.UserContext
	UserInfo() (*entities.UserInfo, error)
	UserMessage(opts *UserMessage) error
}
