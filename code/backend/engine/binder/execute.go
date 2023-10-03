package binder

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
)

var NoPanicWrap = true

func (b *Binder) RPXecute(j *job.RPXJob) ([]byte, error) {

	b.ajLock.Lock()
	b.activeRPXJobs[j.EventId] = j
	b.ajLock.Unlock()

	out, err := b.Executor.RPXecute(etypes.Request{
		Id:      j.EventId,
		Name:    j.Name,
		Data:    j.Payload,
		Invoker: j.Invoker,
	})

	b.ajLock.Lock()
	delete(b.activeRPXJobs, j.EventId)
	b.ajLock.Unlock()

	return out, err
}

func (b *Binder) WebRawXecute(j *job.RawWebJob) {
	b.Executor.WebRawXecute(j.Writer, j.Request)
}
