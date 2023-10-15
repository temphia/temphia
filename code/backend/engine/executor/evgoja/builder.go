package evgoja

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type Builder struct {
}

func (b *Builder) New(opts etypes.ExecutorOption) (etypes.Executor, error) {
	evl := eventloop.NewEventLoop()

	g := &EvGoja{
		bindx:    opts.Binder,
		tenantId: opts.TenantId,
		plugId:   opts.PlugId,
		agentId:  opts.AgentId,
		evLoop:   evl,
	}

	registry := require.NewRegistry()

	evl.Run(func(r *goja.Runtime) {
		g.rt = r

		registry.Enable(r)
		registry.RegisterNativeModule("temphia", g.temphiaBindings)
	})

	evl.Start()

	return nil, nil
}

func (b *Builder) ServeFile(file string) (xtypes.BeBytes, error) {

	return nil, nil
}
