package etypes

import (
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

type BuilderFactory func(app any) (ExecutorBuilder, error)

type ExecutorOption struct {
	Binder   bindx.Bindings
	TenantId string
	PlugId   string
	AgentId  string
	File     string
	ExecType string
	EnvVars  map[string]string
}

type ExecutorIface struct {
	Methods     map[string]*Method    `json:"methods,omitempty"`
	Events      map[string]*EventType `json:"events,omitempty"`
	Schemas     map[string]*ValueType `json:"schemas,omitempty"`
	Bindings    map[string]*Method    `json:"bindings,omitempty"`
	Definations map[string]any        `json:"definations,omitempty"`
}

type ExecutorBuilder interface {
	New(ExecutorOption) (Executor, error)
	ServeFile(file string) (xtypes.BeBytes, error)
	SetRemoteOptions(opts RemoteOptions) error
}

type Executor interface {
	RPXecute(r Request) (xtypes.BeBytes, error)
	WebRawXecute(rw http.ResponseWriter, req *http.Request)
	Reset() error
}

type ExecBuilderFunc func(ExecutorOption) (Executor, error)

func (e ExecBuilderFunc) New(opts ExecutorOption) (Executor, error) {
	return e(opts)
}

func (e ExecBuilderFunc) ServeFile(file string) (xtypes.BeBytes, error) {
	return nil, nil
}
