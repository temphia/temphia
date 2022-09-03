package etypes

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bx"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
)

type ExecutorBinder interface {
	bx.Bindings

	AttachJob(j *job.Job)
	Execute() (*event.Response, error)
}
