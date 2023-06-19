package binder

import (
	"context"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"gitlab.com/mr_balloon/golib"
)

func (b *Binder) AttachJob(j *job.Job) {
	b.Job = j
	b.ReuseCounter = b.ReuseCounter + 1
	b.initLogger()
	b.Context = context.Background()
	b.EventId = j.EventId
	b.Resp = nil

	// b.invoker = NewInvoker(b.Handle) // fixme
}

func (b *Binder) SetExec(exec etypes.Executor) {
	b.executor = exec
}

var NoPanicWrap = true

func (b *Binder) Execute() (*event.Response, error) {
	b.logInfo().Msg(logid.BinderEventProcessStart)
	b.logDebug().Interface("job_req", b.Job).Msg(logid.BinderEventRequestDebug)

	var eresp *event.Response
	var err error

	if NoPanicWrap {
		eresp, err = b.executor.Process(b.Job.AsEvent())
	} else {
		perr := golib.PanicWrapper(func() {
			eresp, err = b.executor.Process(b.Job.AsEvent())
		})

		if perr != nil {
			b.logErr().Err(perr).Msg(logid.BinderExecutePanicked)
			return nil, perr
		}
	}

	if err != nil {
		b.logErr().Err(err).Msg(logid.BinderExecuteErr)
		return nil, err
	}

	b.logErr().Msg(logid.BinderEventProcessOK)

	if eresp == nil {
		eresp = &event.Response{}
	}

	if eresp.Payload == nil {
		if b.Resp != nil {
			eresp.Payload = b.Resp
		}
	}

	b.logDebug().Interface("resp", b.Job).Msg(logid.BinderEventResponseDebug)

	return eresp, nil
}
