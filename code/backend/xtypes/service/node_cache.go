package service

type NodeCache interface {
	Put(tenantId, space, key string, value []byte, expire int64) error
	PutCAS(tenantId, space, key string, value []byte, version, expire int64) error
	Get(tenantId, space, key string) (data []byte, version int64, expire int64, err error)
	Expire(tenantId, space, key string) error
}
