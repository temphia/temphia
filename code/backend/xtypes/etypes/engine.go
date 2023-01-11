package etypes

import "github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"

type Execution struct {
	TenantId string
	PlugId   string
	AgentId  string
	Action   string
	Payload  []byte
	Invoker  invoker.Invoker
}

type Engine interface {
	Run() error
	GetRuntime() Runtime

	Execute(options Execution) ([]byte, error)

	ServeAgentFile(tenantId, plugId, agentId, file string) ([]byte, error)
	ServeExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error)
}
