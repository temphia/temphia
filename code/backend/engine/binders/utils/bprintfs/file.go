package bprintfs

import (
	"io"
	"io/fs"
	"time"

	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
)

type File struct {
	name      string
	b         bindx.Bindings
	dataCache []byte
	offset    int64
}

func (s *File) Stat() (fs.FileInfo, error) { return s, nil }
func (s *File) Read(b []byte) (int, error) {
	if s.dataCache == nil {
		err := s.fillData()
		if err != nil {
			return 0, err
		}
	}

	if s.offset >= int64(len(s.dataCache)) {
		return 0, io.EOF
	}

	if s.offset < 0 {
		return 0, &fs.PathError{Op: "read", Path: s.name, Err: fs.ErrInvalid}
	}

	n := copy(b, s.dataCache[s.offset:])
	s.offset += int64(n)
	return n, nil
}

func (s *File) Close() error {
	s.b = nil
	s.dataCache = nil
	return nil
}

func (s *File) fillData() error {
	data, _, err := s.b.GetFileWithMeta(s.name)
	if err != nil {
		return err
	}
	s.dataCache = data
	return nil
}

// fs.fsinfo
func (s *File) Name() string { return s.name }
func (s *File) Size() int64 {
	if s.dataCache == nil {
		err := s.fillData()
		if err != nil {
			return 0
		}
	}

	return int64(len(s.dataCache))
}
func (s *File) Mode() fs.FileMode  { return fs.FileMode(0666) }
func (s *File) ModTime() time.Time { return time.Time{} }
func (s *File) IsDir() bool        { return false }
func (s *File) Sys() any           { return nil }
