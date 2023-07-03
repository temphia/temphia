package python

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/executors/runner"
)

func TestPyRunner(t *testing.T) {

	runner := runner.New(&runner.Options{
		BootstrapFunc: BootstrapProject,
		Runcmd:        "bash start.sh",
		EntryFile:     "main.py.zip",
		GetFile: func(name string) ([]byte, error) {

			var buf bytes.Buffer

			z := zip.NewWriter(&buf)
			defer z.Close()

			for _, fk := range []string{"main.py", "start.sh"} {

				rfile, err := os.Open(path.Join("testdata", fk))
				if err != nil {
					return nil, err
				}

				defer rfile.Close()

				wfile, err := z.Create(fk)
				if err != nil {
					return nil, err
				}

				if _, err := io.Copy(wfile, rfile); err != nil {
					return nil, err
				}
			}

			return nil, nil
		},
	})

	err := runner.Init()
	if err != nil {
		pp.Println(err)
		t.Fatal(err)
	}

}
