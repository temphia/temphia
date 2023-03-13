package binder

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/temphia/temphia/code/backend/engine/binder/handle"
	"github.com/temphia/temphia/code/backend/engine/binder/specific/cab"
	"github.com/temphia/temphia/code/backend/engine/binder/specific/ncache"
	"github.com/temphia/temphia/code/backend/engine/binder/specific/net"
	"github.com/temphia/temphia/code/backend/engine/binder/specific/plugkv"
	"github.com/temphia/temphia/code/backend/engine/binder/specific/self"
	"github.com/temphia/temphia/code/backend/engine/binder/specific/sockd"
	"github.com/temphia/temphia/code/backend/engine/binder/specific/user"
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
	plugKV  plugkv.Binding
	sockd   sockd.Binding
	cabinet cab.Binding
	net     net.Binding
	ncache  ncache.Binding
	user    user.Binding
	self    self.Binding
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
	b.sockd = sockd.New(b.Handle)
	b.cabinet = cab.New(b.Handle)
	b.net = net.New()
	b.ncache = ncache.New(b.Handle)
	b.user = user.New(b.Handle)
	b.self = self.New(b.Handle)
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

func (b *Binder) PlugKVBindingsGet() bindx.PlugKV { return &b.plugKV }
func (b *Binder) SelfBindingsGet() bindx.Self     { return &b.self }
func (b *Binder) NetGet() bindx.Net               { return &b.net }

func (b *Binder) SockdBindingsGet() bindx.Sockd     { return &b.sockd }
func (b *Binder) CabinetBindingsGet() bindx.Cabinet { return &b.cabinet }
func (b *Binder) NodeCacheGet() bindx.NodeCache     { return &b.ncache }
func (b *Binder) UserBindingsGet() bindx.User       { return &b.user }

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
