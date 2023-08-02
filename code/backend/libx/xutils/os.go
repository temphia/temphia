package xutils

import (
	"errors"
	"io"
	"net"
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

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
