package ncache

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes/service"
)

type Binding struct {
	ncache   service.NodeCache
	tenantId string
	plugId   string
}

const (
	PlugSpace = "plug"
)

func New(ncache service.NodeCache, tenantId string, plugId string) Binding {

	return Binding{
		ncache:   ncache,
		tenantId: tenantId,
		plugId:   plugId,
	}
}

func (b *Binding) Put(key string, value []byte, expire int64) error {
	return b.ncache.Put(b.tenantId, PlugSpace, b.plugKey(key), value, expire)
}

func (b *Binding) PutCAS(key string, value []byte, version, expire int64) error {
	return b.ncache.PutCAS(b.tenantId, PlugSpace, b.plugKey(key), value, version, expire)
}

func (b *Binding) Get(key string) ([]byte, int64, int64, error) {
	return b.ncache.Get(b.tenantId, PlugSpace, b.plugKey(key))
}

func (b *Binding) Expire(key string) error {
	return b.ncache.Expire(b.tenantId, PlugSpace, b.plugKey(key))
}

// private

func (b *Binding) plugKey(key string) string {
	return fmt.Sprintf("%s__%s", b.plugId, key)
}
