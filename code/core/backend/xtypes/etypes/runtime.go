package etypes

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
)

type RunningExec struct {
	EventId  string
	BprintId string
	PlugId   string
	AgentId  string
}

type Runtime interface {
	Run(map[string]ExecutorBuilder, map[string]ModuleBuilder) error

	Preform(j *job.Job) (*event.Response, error)
	PreformAsync(j *job.AsyncJob)

	ResetAgents(tenantId, plug string, agents []string)
	ResetBprint(tenantId, bprint string)
	ListRunning(tenantId string) ([]RunningExec, error)
}

type Router interface {
	Route(j *job.Job) bool
}
