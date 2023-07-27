package xutils

import (
	"errors"
	"io"
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

func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}

	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
