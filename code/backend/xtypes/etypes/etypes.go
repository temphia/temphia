package etypes

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
)

type ExecutorBinder interface {
	bindx.Bindings

	GetApp() any
	GetModuleBuilder(name string) (any, error)
	GetLogger() *zerolog.Logger

	AttachJob(j *job.Job)
	Execute() (*event.Response, error)
}
