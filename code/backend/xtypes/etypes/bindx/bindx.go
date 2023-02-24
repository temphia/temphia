package bindx

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Bindings interface {
	Core

	PlugKVBindingsGet() PlugKV
	SockdBindingsGet() Sockd
	UserBindingsGet() User
	CabinetBindingsGet() Cabinet
	SelfBindingsGet() Self
	NodeCacheGet() NodeCache
	NetGet() Net
}

type Core interface {
	Log(msg string)
	LazyLog(msgs []string)
	Sleep(int32)
	GetFileWithMeta(file string) (data []byte, version int64, err error)
	GetApp() any
}

type PlugStateTkt struct {
	KeyPrefix string `json:"key_prefix,omitempty"`
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

	Ticket(opts *PlugStateTkt) (string, error)
}

type Cabinet interface {
	AddFile(bucket string, file string, contents []byte) error
	ListFolder(bucket string) ([]string, error)
	GetFile(bucket string, file string) ([]byte, error)
	DeleteFile(bucket string, file string) error

	GenerateTicket(bucket string, ticket *CabTicket) (string, error)
}

type Sockd interface {
	SendDirect(room string, connId int64, payload []byte) error
	SendDirectBatch(room string, conns []int64, payload []byte) error
	SendBroadcast(room string, ignores []int64, payload []byte) error
	SendTagged(room string, tags []string, ignores []int64, payload []byte) error
	RoomUpdateTags(room string, opts sockdx.UpdateTagOptions) error
}

type User interface {
	ListUser(group string) ([]string, error)
	MessageUser(group, user string, opts *UserMessage) error
	GetUser(group, user string) (*entities.UserInfo, error)

	MessageCurrentUser(opts *UserMessage) error
	CurrentUser() (*entities.UserInfo, error)
}

type Net interface {
	HttpRaw(*HttpRequest) *HttpResponse
	HttpRawBatch([]*HttpRequest) []*HttpResponse

	HttpQuickGet(url string, headers map[string]string) ([]byte, error)
	HttpQuickPost(url string, headers map[string]string, data []byte) ([]byte, error)
	HttpFormPost(url string, headers map[string]string, data []byte) ([]byte, error)

	HttpJsonGet(url string, headers map[string]string) ([]byte, error)
	HttpJsonPost(url string, headers map[string]string, data []byte) ([]byte, error)
}

type NodeCache interface {
	Put(key string, value []byte, expire int64) error
	PutCAS(key string, value []byte, version, expire int64) error
	Get(key string) (data []byte, version int64, expire int64, err error)
	Expire(key string) error
}

type Self interface {
	SelfAddFile(file string, data []byte) error
	SelfUpdateFile(file string, data []byte) error

	SelfAddDataFile(file string, data []byte) error
	SelfUpdateDataFile(file string, data []byte) error
	SelfGetDataFile(file string) ([]byte, error)
	SelfListDataFiles() (map[string]string, error)
	SelfDeleteDataFile(file string) error

	SelfListResources() ([]*Resource, error)
	SelfGetResource(name string) (*Resource, error)

	SelfInLinks() ([]Link, error)
	SelfOutLinks() ([]Link, error)

	SelfLinkExec(name, method string, data xtypes.LazyData, async, detached bool) (xtypes.LazyData, error)
	SelfModuleExec(name, method, path string, data xtypes.LazyData) (xtypes.LazyData, error)

	SelfForkExec(method string, data []byte) error
}
