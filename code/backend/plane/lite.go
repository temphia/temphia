package plane

import (
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/plane/idservice"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

var _ xplane.ControlPlane = (*PlaneLite)(nil)

type PlaneLite struct {
	eventbus *EventBus
	locker   *Locker
	router   *Router
	nodeId   int64
	seq      idservice.IDService
}

func NewLite(CoreHub store.CoreHub) *PlaneLite {

	nodeId := int64(1)

	return &PlaneLite{
		eventbus: NewEventBus(),
		locker:   NewLocker(),
		router:   nil,
		nodeId:   nodeId,
		seq:      *idservice.New(nodeId),
	}
}

// dl and start up stuff

func (p *PlaneLite) Start() error                                         { return nil }
func (p *PlaneLite) Inject(iapp interface{}, config *config.Config) error { return nil }

// liveness and status stuff

func (p *PlaneLite) NotifyStat(stats xplane.NodeStat) error   { return nil }
func (p *PlaneLite) GetAppStatus() (*xplane.AppStatus, error) { return nil, nil }

// router stuff
func (p *PlaneLite) SetJobChan(ic chan *job.Job) {
	p.router = NewRouter(ic)
}

func (p *PlaneLite) GetRouter() xplane.Router {
	return p.router
}

// locker
func (p *PlaneLite) GetLocker() xplane.Locker {
	return nil
}

// sockdrouter
func (p *PlaneLite) GetSockdRouter() xplane.SockdRouter {
	return nil
}

// eventbus
func (p *PlaneLite) GetEventBus() xplane.EventBus { return p.eventbus }

func (p *PlaneLite) GetNodeId() int64 {
	return p.nodeId
}

func (p *PlaneLite) GetIdService() xplane.IDService { return &p.seq }
