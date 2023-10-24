package evgoja

import (
	"net/http"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

type EvGoja struct {
	evLoop *eventloop.EventLoop
	rt     *goja.Runtime

	bindx    bindx.Bindings
	tenantId string
	plugId   string
	agentId  string
}

func (g *EvGoja) RPXecute(r etypes.Request) (xtypes.BeBytes, error) {
	return g.rPXecute(r)
}

func (g *EvGoja) WebRawXecute(rw http.ResponseWriter, req *http.Request) {
	g.webRawXecute(rw, req)
}

func (g *EvGoja) Close() error {

	return nil
}
