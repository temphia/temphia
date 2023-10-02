package binder

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
)

var NoPanicWrap = true

func (b *Binder) Execute() ([]byte, error) {

	return nil, nil
}

func (b *Binder) RPXecuteWithExecutor(j *job.RPXJob, ex etypes.Executor) ([]byte, error) {

	b.ajLock.Lock()
	b.activeRPXJobs[j.EventId] = j
	b.ajLock.Unlock()

	out, err := ex.RPXecute(etypes.Request{
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
