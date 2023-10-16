package noop

import (
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

// noop executor that just logs everything and does nothing

type Builder struct {
}

func (b *Builder) New(opts etypes.ExecutorOption) (etypes.Executor, error) {

	return &Noop{
		tenantId: opts.TenantId,
		plugId:   opts.PlugId,
		agentId:  opts.AgentId,
	}, nil
}

func (b *Builder) ServeFile(file string) (xtypes.BeBytes, error) {

	pp.Println("@serving file", file)

	return nil, nil
}

func (b *Builder) SetRemoteOptions(opts etypes.RemoteOptions) error {
	return nil
}

type Noop struct {
	tenantId string
	plugId   string
	agentId  string
}

func (n *Noop) RPXecute(r etypes.Request) (xtypes.BeBytes, error) {

	pp.Println("@rpx_execute", r)

	return nil, nil

}

func (n *Noop) WebRawXecute(rw http.ResponseWriter, req *http.Request) {
	pp.Println("@web_raw_execute", req.URL.Path)
}

func (n *Noop) Reset() error {

	return nil
}
