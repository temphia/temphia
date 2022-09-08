package corehub

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

type Cache struct {
	inner bigcache.BigCache
}

func New() *Cache {

	cache, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             16,
		LifeWindow:         10 * time.Minute,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 100 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		panic(err)
	}

	return &Cache{
		inner: *cache,
	}
}
