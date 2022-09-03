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

func Die(args ...interface{}) {

	pp.Println(args...)
	os.Exit(1)

}
