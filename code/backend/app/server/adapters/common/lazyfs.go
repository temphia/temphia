package common

import (
	"io/fs"
)

type LazyFSOptions struct {
	Tenant  string
	Handler func(tenantId, file string) ([]byte, error)
	Files   map[string]struct{}
}

type LazyFS struct {
	tenantId string
	handler  func(tenantId, file string) ([]byte, error)
	files    map[string]struct{}
}

func NewLazyFS(opts LazyFSOptions) *LazyFS {
	return &LazyFS{
		handler:  opts.Handler,
		tenantId: opts.Tenant,
		files:    opts.Files,
	}
}

func (f *LazyFS) Open(name string) (fs.File, error) {
	return &File{
		name:   name,
		parent: f,
	}, nil
}
