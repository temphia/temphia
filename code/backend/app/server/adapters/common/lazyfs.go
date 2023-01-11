package common

import (
	"io/fs"
)

type LazyFSOptions struct {
	Prefix  string
	Folder  string
	Tenant  string
	Handler func(tenantId, folder, file string) ([]byte, error)
}

type LazyFS struct {
	tenantId string
	folder   string
	prefix   string
	handler  func(tenantId, folder, file string) ([]byte, error)
}

func NewLazyFS(opts LazyFSOptions) *LazyFS {
	return &LazyFS{
		prefix:   opts.Prefix,
		handler:  opts.Handler,
		tenantId: opts.Tenant,
		folder:   opts.Folder,
	}
}

func (f *LazyFS) Open(name string) (fs.File, error) {
	return &File{
		name:   name,
		parent: f,
	}, nil
}
