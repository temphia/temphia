package etypes

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
)

type ExecutorBinder interface {
	bindx.Bindings

	AttachJob(j *job.Job)
	Execute() (*event.Response, error)
}
