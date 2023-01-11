package bprintfs

import (
	"io/fs"

	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
)

var _ fs.FS = (*FS)(nil)

type FS struct {
	b bindx.Bindings

	// files []string
}

func New(b bindx.Bindings) *FS {
	return &FS{b: b}
}

func (s *FS) Open(name string) (fs.File, error) {
	return &File{
		name: name,
		b:    s.b,
	}, nil
}
