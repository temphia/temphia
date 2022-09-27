package xutils

import (
	"errors"
	"os"
	"path"

	"github.com/k0kubun/pp"
)

func FileExists(dpath, file string) bool {
	_, err := os.Stat(path.Join(dpath, file))
	if err == nil {
		return true
	}
	return !errors.Is(err, os.ErrNotExist)
}

func CreateIfNotExits(fpath string) error {
	if _, err := os.Stat(fpath); errors.Is(err, os.ErrNotExist) {
		return os.Mkdir(fpath, os.ModePerm)
	}

	return nil
}

func Die(args ...any) {

	pp.Println(args...)
	os.Exit(1)

}
