package ncache

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/engine/binder/handle"
	"github.com/temphia/temphia/code/backend/xtypes/service"
)

type Binding struct {
	ncache service.NodeCache
	handle *handle.Handle
}

const (
	PlugSpace = "plug"
)

func New(handle *handle.Handle) Binding {

	return Binding{
		ncache: handle.Deps.NodeCache,
		handle: handle,
	}
}

func (b *Binding) Put(key string, value []byte, expire int64) error {
	return b.ncache.Put(b.handle.Namespace, PlugSpace, b.plugKey(key), value, expire)
}

func (b *Binding) PutCAS(key string, value []byte, version, expire int64) error {
	return b.ncache.PutCAS(b.handle.Namespace, PlugSpace, b.plugKey(key), value, version, expire)
}

func (b *Binding) Get(key string) ([]byte, int64, int64, error) {
	return b.ncache.Get(b.handle.Namespace, PlugSpace, b.plugKey(key))
}

func (b *Binding) Expire(key string) error {
	return b.ncache.Expire(b.handle.Namespace, PlugSpace, b.plugKey(key))
}

// private

func (b *Binding) plugKey(key string) string {
	return fmt.Sprintf("%s__%s", b.handle.PlugId, key)
}
