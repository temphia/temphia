package etypes

import (
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type Execution struct {
	TenantId string
	PlugId   string
	AgentId  string
	Action   string
	Payload  []byte
	Invoker  invoker.Invoker
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

type RemoteOptions struct{}

type Engine interface {
	Run() error

	GetCache() Ecache

	RPXecute(options Execution) ([]byte, error)
	WebRawXecute(rw http.ResponseWriter, req *http.Request)

	SetRemoteOption(opt RemoteOptions)
	ResetAgent(tenantId, plugId, agentId string) error

	ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error)
	ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error)

	ListExecutors() []string
	ListModules() []string

	RemotePerform(opt Remote) ([]byte, error)
}

type Ecache interface {
	GetAgent(tenantId, plug, agent string) *entities.Agent
	GetPlug(tenantId, plugId string) *entities.Plug
}

/*

type RunningExec struct {
	EventId  string
	BprintId string
	PlugId   string
	AgentId  string
}

ListRunning(tenantId string) ([]RunningExec, error)

*/
