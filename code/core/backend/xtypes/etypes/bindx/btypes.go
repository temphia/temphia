package bindx

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type Bindings interface {
	BindCore

	PlugKVBindingsGet() BindPlugKV
	SockdBindingsGet() BindSockd
	UserBindingsGet() BindUser
	CabinetBindingsGet() BindCabinet
	SelfBindingsGet() BindSelf
	NodeCacheGet() BindNodeCache
}

type BindCore interface {
	Log(msg string)
	LazyLog(msgs []string)
	Sleep(int32)
	GetFileWithMeta(file string) (data []byte, version int64, err error)
	GetApp() any
}

type BindPlugKV interface {
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

type BindCabinet interface {
	AddFile(bucket string, file string, contents []byte) error
	ListFolder(bucket string) ([]string, error)
	GetFile(bucket string, file string) ([]byte, error)
	DeleteFile(bucket string, file string) error
	GenerateTicket(bucket string, ticket *CabTicket) (string, error)
}

type BindSockd interface {
	SendDirect(room string, connId int64, payload []byte) error
	SendDirectBatch(room string, conns []int64, payload []byte) error
	SendBroadcast(room string, ignores []int64, payload []byte) error
	SendTagged(room string, tags []string, ignores []int64, payload []byte) error
	RoomUpdateTags(room string, opts sockdx.UpdateTagOptions) error
}

type BindUser interface {
	ListUsers(group string) ([]string, error)
	MessageUser(group, user, message string, encrypted bool) error
	GetUser(group, user string) (*entities.UserInfo, error)

	MessageCurrentUser(title, message string, encrypted bool) error
	CurrentUser() (*entities.UserInfo, error)
}

type BindNet interface {
	HttpRaw(*HttpRequest) *HttpResponse
	HttpRawBatch([]*HttpRequest) []*HttpResponse

	HttpQuickGet(url string, headers map[string]string) ([]byte, error)
	HttpQuickPost(url string, headers map[string]string, data []byte) ([]byte, error)
	HttpFormPost(url string, headers map[string]string, data []byte) ([]byte, error)

	HttpJsonGet(url string, headers map[string]string) ([]byte, error)
	HttpJsonPost(url string, headers map[string]string, data []byte) ([]byte, error)
}

type BindNodeCache interface {
	Put(key string, value []byte, expire int64) error
	PutCAS(key string, value []byte, version, expire int64) error
	Get(key string) (data []byte, version int64, expire int64, err error)
	Expire(key string) error
}

type BindSelf interface {
	SelfGetFile(file string) ([]byte, error)
	SelfAddFile(file string, data []byte) error
	SelfUpdateFile(file string, data []byte) error

	SelfListResources() ([]*Resource, error)
	SelfGetResource(name string) (*Resource, error)

	SelfInLinks() ([]Link, error)
	SelfOutLinks() ([]Link, error)

	SelfLinkExec(name, method string, data xtypes.LazyData, async, detached bool) (xtypes.LazyData, error)
	SelfModuleExec(name, method, path string, data xtypes.LazyData) (xtypes.LazyData, error)

	SelfForkExec(method string, data []byte) error
}

// future

type BindLocker interface {
	// fixme => nested key lock

	SelfLockWait(key string) error
	SelfLock(key string, expiry int) error
	SelfLockRenew(key string, expiry int) error
	SelfUnLock(key string) error

	ResourceLockWait(resource string, key string) error
	ResourceLock(resource string, key string) error
	ResourceLockRenew(resource string, key string) error
	ResourceUnLock(resource string, key string) error
}
