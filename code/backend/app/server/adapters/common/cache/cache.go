package cache

import (
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"gopkg.in/hypirion/go-filecache.v1"
)

var (
	_ httpx.GlobalCache = (*Cache)(nil)
)

type Cache struct {
	subCaches map[string]*SubCache
	mlock     sync.Mutex
}

func New() *Cache {
	return &Cache{
		subCaches: make(map[string]*SubCache),
		mlock:     sync.Mutex{},
	}
}

func (c *Cache) GetSubCache(key string, loader httpx.CacheLoader) (httpx.SubCache, error) {

	c.mlock.Lock()
	sc := c.subCaches[key]
	c.mlock.Unlock()

	if sc != nil {
		return sc, nil
	}

	fc, err := filecache.New(filecache.MB*100, loader)
	if err != nil {
		return nil, err
	}

	sc = &SubCache{
		filecache: fc,
	}

	c.mlock.Lock()
	if _sc, ok := c.subCaches[key]; ok {
		c.mlock.Unlock()
		return _sc, nil
	}

	c.subCaches[key] = sc
	c.mlock.Unlock()

	return sc, nil
}
