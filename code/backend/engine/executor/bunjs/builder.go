package bunjs

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type Builder struct {
}

func (b *Builder) New(opts etypes.ExecutorOption) (etypes.Executor, error) {

	return &BunJS{
		tenantId: opts.TenantId,
		plugId:   opts.PlugId,
		agentId:  opts.AgentId,
		addr:     "",
	}, nil
}

func (b *Builder) ServeFile(file string) (xtypes.BeBytes, error) {

	pp.Println("@serving file", file)

	return nil, nil
}

func (b *Builder) SetRemoteOptions(opts etypes.RemoteOptions) error {

	return nil
}
