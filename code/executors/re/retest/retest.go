package retest

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/executors/re"
	"github.com/temphia/temphia/code/executors/re/langs/python"
)

func TestRe(t *testing.T) {

	builder := re.NewBuilder("re_python", "bash start.sh", python.BootstrapProject)

	exec, err := builder.Instance(etypes.ExecutorOption{
		Binder:   &MB{},
		TenantId: "default0",
		PlugId:   "pp1",
		AgentId:  "aa1",
		File:     "main.py.zip",
	})
	if err != nil {
		panic(err)
	}

	resp, err := exec.Process(&event.Request{
		Id:   "11",
		Name: "ping",
		Data: []byte(`{}`),
	})
	if err != nil {
		panic(err)
	}

	pp.Println("@final", string(resp.Payload))
}

func getTestZip() ([]byte, error) {

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

	return buf.Bytes(), nil

}

type MB struct {
	bindx.Bindings
}

func (m *MB) GetFileWithMeta(file string) (data []byte, version int64, err error) {
	out, err := getTestZip()
	return out, 0, err
}

func (m *MB) Clone() bindx.Core {

	return m
}
