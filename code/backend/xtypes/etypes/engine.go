package etypes

import (
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
)

type Execution struct {
	TenantId string
	PlugId   string
	AgentId  string
	Action   string
	Payload  []byte
	Invoker  invoker.Invoker
}

type REOptions struct {
}

type Engine interface {
	Run() error
	GetRuntime() Runtime
	ListExecutors() []string
	ListModules() []string

	RPXecute(options Execution) ([]byte, error)
	WebRawXecute(rw http.ResponseWriter, req *http.Request)
	SetREOption(opt REOptions)

	ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error)
	ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error)
}
