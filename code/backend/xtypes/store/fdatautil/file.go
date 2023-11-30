package fdatautil

import (
	"io"
	"os"

	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var _ store.FData = (*FileFdata)(nil)

type FileFdata struct {
	fpath string
	file  *os.File
}

func NewFromFile(fpath string) *FileFdata {
	return &FileFdata{
		fpath: fpath,
		file:  nil,
	}
}

func (f *FileFdata) AsBytes() ([]byte, error) {
	return os.ReadFile(f.fpath)
}

func (f *FileFdata) AsReader() (io.Reader, error) {

	file, err := os.Open(f.fpath)
	if err != nil {
		return nil, err
	}

	f.file = file

	return file, nil
}

func (f *FileFdata) Close() error {
	if f.file != nil {
		f.file.Close()
	}

	return os.Remove(f.fpath)

}

func (f *FileFdata) InnerFile() string {
	return f.fpath
}
