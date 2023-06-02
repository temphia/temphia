package binder

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/temphia/temphia/code/backend/engine/binder/handle"
	plugkv "github.com/temphia/temphia/code/backend/engine/modules/plugkv2"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"

	"gitlab.com/mr_balloon/golib"
)

var (
	_ bindx.Bindings = (*Binder)(nil)
)

type Binder struct {
	Handle       *handle.Handle
	executor     etypes.Executor
	ReuseCounter int32
	Epoch        int64

	// specific bind impl
	plugKV plugkv.Binding
	sockd  SockdBinding

	self    SelfBindings
	invoker InvokerBindings
}

func (b *Binder) AttachJob(j *job.Job) {
	b.Handle.Job = j
	b.ReuseCounter = b.ReuseCounter + 1
	b.Handle.InitLogger()
	b.Handle.Context = context.Background()
	b.Handle.EventId = j.EventId
	b.Handle.Resp = nil

	// build specific binds
	b.plugKV = plugkv.New(b.Handle)
	b.sockd = SockdBinding{
		sockd: b.Handle.Deps.Sockd,
	}

	b.self = NewSelfBindings(b.Handle, b)
	b.invoker = NewInvoker(b.Handle)
}

func (b *Binder) SetExec(exec etypes.Executor) {
	b.executor = exec
}

var NoPanicWrap = true

func (b *Binder) Execute() (*event.Response, error) {
	b.logInfo().Msg(logid.BinderEventProcessStart)
	b.logDebug().Interface("job_req", b.Handle.Job).Msg(logid.BinderEventRequestDebug)

	var eresp *event.Response
	var err error

	if NoPanicWrap {
		eresp, err = b.executor.Process(b.Handle.Job.AsEvent())
	} else {
		perr := golib.PanicWrapper(func() {
			eresp, err = b.executor.Process(b.Handle.Job.AsEvent())
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
		if b.Handle.Resp != nil {
			eresp.Payload = b.Handle.Resp
		}
	}

	b.logDebug().Interface("resp", b.Handle.Job).Msg(logid.BinderEventResponseDebug)

	return eresp, nil
}

// bindings

func (b *Binder) PlugKVBindingsGet() any { return nil }

func (b *Binder) SelfBindingsGet() bindx.Self   { return &b.self }
func (b *Binder) InvokerGet() bindx.Invoker     { return &b.invoker }
func (b *Binder) SockdBindingsGet() bindx.Sockd { return &b.sockd }

// private

func (b *Binder) logInfo() *zerolog.Event {
	return b.Handle.Logger.Info()
}

func (b *Binder) logErr() *zerolog.Event {
	return b.Handle.Logger.Info()
}

func (b *Binder) logDebug() *zerolog.Event {
	return b.Handle.Logger.Info()
}
