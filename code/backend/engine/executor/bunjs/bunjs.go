package bunjs

import (
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type BunJS struct {
	tenantId string
	plugId   string
	agentId  string
	addr     string
}

func (b *BunJS) RPXecute(r etypes.Request) (xtypes.BeBytes, error) {

	pp.Println("@rpx_execute", r)

	return nil, nil

}

func (b *BunJS) WebRawXecute(rw http.ResponseWriter, req *http.Request) {
	pp.Println("@web_raw_execute", req.URL.Path)
}

func (b *BunJS) Reset() error {

	return nil
}
