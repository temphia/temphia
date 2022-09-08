package cache

import (
	"github.com/dgraph-io/ristretto"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

type Cache struct {
	inner ristretto.Cache
}

func New() *Cache {

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	if err != nil {
		panic(err)
	}

	return &Cache{
		inner: *cache,
	}

}

func (c *Cache) BprintGet(tenantId, bid string) *entities.BPrint {

	// c.inner.Set()

	return nil
}
func (c *Cache) BprintSet() {}
