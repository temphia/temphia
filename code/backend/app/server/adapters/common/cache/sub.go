package cache

import (
	"io"

	"gopkg.in/hypirion/go-filecache.v1"
)

type SubCache struct {
	filecache *filecache.Filecache
}

func (s *SubCache) Has(key string) (bool, error) {
	return s.filecache.Has(key)
}

func (s *SubCache) Get(dst io.Writer, key string) error {
	return s.filecache.Get(dst, key)
}
