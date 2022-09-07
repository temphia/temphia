package etypes

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
)

type Execution struct {
	TenantId string
	PlugId   string
	AgentId  string
	Action   string
	Payload  []byte
	Invoker  job.Invoker
}

type Engine interface {
	Run() error
	GetRuntime() Runtime

	Execute(options Execution) ([]byte, error)

	ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error)
	ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error)
}
