package rfencer

import "github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"

type Env struct {
	Job *job.Job
}

func (e *Env) PyloadGet(query string) any {
	return nil
}

func (e *Env) PayloadModify() bool {
	return false
}
