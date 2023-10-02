package etypes

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
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

	InitAgent(tenantId, plug, agent string) error

	ResetAgents(tenantId, plug string, agents []string)
	ResetBprint(tenantId, bprint string)
	ListRunning(tenantId string) ([]RunningExec, error)
}

type Router interface {
	Route(j *job.Job) (*event.Response, error)
}
