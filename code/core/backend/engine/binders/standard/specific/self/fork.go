package self

import (
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/core/backend/engine/invoker/forked"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
)

func (b *Binding) selfForkExec(method string, data []byte) error {

	// fixme => log here

	go func() {
		b.runtime.Preform(&job.Job{
			PlugId:            b.handle.PlugId,
			AgentId:           b.handle.AgentId,
			EventId:           xid.New().String(),
			EventAction:       method,
			Namespace:         b.handle.Namespace,
			Payload:           data,
			PendingPrePolicy:  true,
			PendingPostPolicy: true,
			Loaded:            b.handle.Job.Loaded,
			Plug:              b.handle.Job.Plug,
			Agent:             b.handle.Job.Agent,
			Invoker:           forked.New(b.handle.EventId),
		})
	}()

	return nil

}
