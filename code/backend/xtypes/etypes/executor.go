package etypes

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
)

type BuilderFactory func(app any) (ExecutorBuilder, error)

type ExecutorOption struct {
	Binder   bindx.Bindings
	TenantId string
	PlugId   string
	AgentId  string
	File     string
	ExecType string
	EnvVars  map[string]any
}

type Executor interface {
	Process(*event.Request) (*event.Response, error)
}

type ExecutorBuilder interface {
	Instance(ExecutorOption) (Executor, error)
	ExecFile(file string) ([]byte, error)
}

type ExecBuilderFunc func(ExecutorOption) (Executor, error)

func (e ExecBuilderFunc) Instance(opts ExecutorOption) (Executor, error) {
	return e(opts)
}

func (e ExecBuilderFunc) ExecFile(file string) ([]byte, error) { return []byte(``), nil }
