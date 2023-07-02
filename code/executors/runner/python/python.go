package python

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/temphia/temphia/code/executors/runner/rtypes"
)

// go:embed lib.py
var Lib []byte

// go:embed bootstrap.sh
var Bootstrap []byte

func BootstrapProject(ctx rtypes.BootstrapContext) error {

	pcmd := exec.Command("python", "-m", "venv", ctx.Folder)

	pcmd.Stderr = os.Stderr
	pcmd.Stdout = os.Stdout

	err := pcmd.Run()
	if err != nil {
		return err
	}

	bfile := path.Join(ctx.Folder, "bootstrap.sh")

	// chmod +x bootstrap.sh ?

	err = os.WriteFile(bfile, Bootstrap, 0766)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(ctx.Folder, "action_router.py"), Lib, 0766)
	if err != nil {
		return err
	}

	out, err := ctx.GetFile(ctx.File)
	if err != nil {
		return err
	}

	err = unzip(out, ctx.Folder)
	if err != nil {
		return err
	}

	bcmd := exec.Command("bash", bfile)
	bcmd.Stdin = os.Stdin
	bcmd.Stdout = os.Stdout

	return bcmd.Run()

}

func unzip(src []byte, dest string) error {
	r, err := zip.NewReader(bytes.NewReader(src), int64(len(src)))
	if err == nil {
		return err
	}

	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
