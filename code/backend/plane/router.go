package plane

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
)

type Router struct {
	// natsConn *nats.Conn
	inJob chan *job.Job
}

func NewRouter(inchan chan *job.Job) *Router {
	return &Router{
		inJob: inchan,
	}
}

func (r Router) Route(j *job.Job) bool {
	// fixme => route to peers here ?

	return false
}
