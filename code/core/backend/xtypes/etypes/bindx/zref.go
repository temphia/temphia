package bindx

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

	QuickGet(txid uint32, key string) (string, error)
	QuickSet(txid uint32, key, valye string) error

	// user

	ListUserPrefix(group string, prefix string) ([]string, error)

	// sockd

	KickFromRoom(connId int64, room string) error
}
