package binder

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

var (
	_ bindx.Bindings = (*Binder)(nil)
)

type Binder struct {
	Deps      *Factory
	Namespace string
	PlugId    string
	AgentId   string
	BprintId  string

	Context  context.Context
	Executor etypes.Executor
	Logger   zerolog.Logger

	Job *job.Job

	EventId string
	Resp    []byte

	// lazy loaded
	Resources map[string]*entities.Resource
	Links     map[string]*entities.AgentLink

	executor     etypes.Executor
	ReuseCounter int32
	Epoch        int64

	activeModules    map[int32]etypes.Module
	activeModCounter int32

	invoker InvokerBindings
}

// bindings

func (b *Binder) NewModule(name string, data xtypes.LazyData) (int32, error) {
	return b.selfNewModule(name, data)
}

func (b *Binder) ModuleTicket(name string, opts xtypes.LazyData) (string, error) {
	return b.moduleTicket(name, opts)
}

func (b *Binder) ModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return b.selfModuleExec(mid, method, data)
}

func (b *Binder) InLinks() ([]bindx.Link, error) {
	return b.selfInLinks()
}

func (b *Binder) OutLinks() ([]bindx.Link, error) {
	return b.selfOutLinks()
}

func (b *Binder) LinkExec(name, method string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return b.selfLinkExec(name, method, data)
}

func (b *Binder) LinkExecEmit(name, method string, data xtypes.LazyData) error {
	return nil
}

func (b *Binder) ForkExec(method string, data []byte) error {
	return b.selfForkExec(method, data)
}

func (b *Binder) AsyncLinkExec(name, method string, data xtypes.LazyData) (uint32, error) {
	return 0, nil
}
func (b *Binder) AsyncModuleExec(mid int32, method string, data xtypes.LazyData) (uint32, error) {
	return 0, nil
}
func (b *Binder) AsyncEventPoll(mid int32, eid uint32) (xtypes.LazyData, error) {
	return nil, nil
}
func (b *Binder) AsyncEventWait(mid int32, eid uint32) (xtypes.LazyData, error) {
	return nil, nil
}

func (b *Binder) GetInvoker() bindx.Invoker {
	return nil
}

func (b *Binder) Clone() bindx.Core {
	b2 := &Binder{
		executor:      nil,
		activeModules: make(map[int32]etypes.Module),
		Resources:     nil,
		Links:         nil,
		Deps:          b.Deps,
		Namespace:     b.Namespace,
		PlugId:        b.PlugId,
		AgentId:       b.AgentId,
		BprintId:      b.BprintId,
		Context:       context.Background(),
		Job:           nil,
		EventId:       fmt.Sprintf("cloned%s", b.EventId),
	}

	b2.initLogger()

	return b2
}
