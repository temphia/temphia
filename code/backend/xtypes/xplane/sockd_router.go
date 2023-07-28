package xplane

import "github.com/temphia/temphia/code/backend/xtypes/etypes/job"

type Router interface {
	Route(j *job.Job) bool
}

type SockdRouter interface {
	Publish(tenantId, room string, tags map[string]string, rawData []byte) error
	Broadcast(tenantId, room string, rawData []byte) error
	SendSession(tenantId, session string, rawData []byte) error
}
