package enginex

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/enginex/bx"
	"github.com/temphia/temphia/code/core/backend/xtypes/enginex/event"
	"github.com/temphia/temphia/code/core/backend/xtypes/enginex/job"
)

type ExecutorBinder interface {
	bx.Bindings

	AttachJob(j *job.Job)
	Execute() (*event.Response, error)
}
