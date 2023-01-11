package etypes

import (
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
)

type ExecutorBinder interface {
	bindx.Bindings

	AttachJob(j *job.Job)
	Execute() (*event.Response, error)
}
