package bindx

import (
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

// this is just for reference

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

type Sockd2 interface {
}

type Next interface {

	// plugkv

	SetBatch(txid uint32) error
	SetBatchLazy([]BatchRecord) error
	LazyBatchFlush() error

	QuickQuery(txid uint32, query *store.PkvQuery) ([]string, error)

	// user

	ListUserPrefix(group string, prefix string) ([]string, error)

	// sockd

	KickFromRoom(connId int64, room string) error
}

type LazyMetrics struct {
	Key      string
	Interval int64
	Value    int64
}

type BatchRecord struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Tag1  string `json:"tag1,omitempty"`
	Tag2  string `json:"tag2,omitempty"`
	Tag3  string `json:"tag3,omitempty"`
	TTL   int    `json:"ttl,omitempty"`
}

type P interface {
	AddLazyMetric()
	GetLazyMetric()
}
