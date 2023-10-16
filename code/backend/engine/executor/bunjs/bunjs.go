package bunjs

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os/exec"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type BunJS struct {
	tenantId  string
	plugId    string
	agentId   string
	addr      string
	rPXPrefix string
	cmd       *exec.Cmd
	proxy     *httputil.ReverseProxy
}

func (b *BunJS) RPXecute(r etypes.Request) (xtypes.BeBytes, error) {

	rw := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/", b.rPXPrefix), bytes.NewReader(r.Data))
	if err != nil {
		return nil, err
	}

	b.proxy.ServeHTTP(rw, req)

	if rw.Code != http.StatusOK {
		return nil, errors.New(rw.Body.String())
	}

	return rw.Body.Bytes(), nil

}

func (b *BunJS) WebRawXecute(rw http.ResponseWriter, req *http.Request) {
	pp.Println("@web_raw_execute", req.URL.Path)
	b.proxy.ServeHTTP(rw, req)
}

func (b *BunJS) Reset() error {

	return nil
}
