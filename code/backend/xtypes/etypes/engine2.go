package etypes

import (
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
)

type Engine2 interface {
	GetCache() Ecache

	RPXecute(options Execution) ([]byte, error)
	WebRawXecute(rw http.ResponseWriter, req *http.Request)

	SetRemoteOption(opt any)
	ResetAgent(tenantId, plugId, agentId string) error

	ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error)
	ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error)

	RemotePerform(opt Remote) (xtypes.BeBytes, error)
}

type Remote struct {
	TenantId string
	PlugId   string
	AgentId  string
	Eid      string
	Data     xtypes.BeBytes
}

type Request struct {
	Id      string
	Name    string
	Data    xtypes.BeBytes
	Invoker invoker.Invoker
}

type ExecutorBuilder2 interface {
	New(ExecutorOption) (Executor2, error)
	ServeFile(file string) (xtypes.BeBytes, error)
}

type Executor2 interface {
	RPXecute(options Request) (xtypes.BeBytes, error)
	WebRawXecute(rw http.ResponseWriter, req *http.Request)
	SetRemoteOptions(opts any)
	Reset() error
}
