package binder

import (
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/engine/invokers/forked"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
)

func (b *SelfBindings) selfForkExec(method string, data []byte) error {

	newEventId := xid.New().String()
	b.root.logInfo().
		Str("new_event_id", newEventId).
		Msg(logid.BinderExecutionForked)

	go func() {
		b.runtime.Preform(&job.Job{
			PlugId:            b.handle.PlugId,
			AgentId:           b.handle.AgentId,
			EventId:           newEventId,
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
