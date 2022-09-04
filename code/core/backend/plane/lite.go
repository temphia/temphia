package plane

import (
	"github.com/bwmarrin/snowflake"

	"github.com/temphia/temphia/code/core/backend/app/config"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

var _ xplane.ControlPlane = (*PlaneLite)(nil)

type PlaneLite struct {
	eventbus    *EventBus
	locker      *Locker
	router      *Router
	idGenerator *snowflake.Node
	nodeId      int64
}

func NewLite() *PlaneLite {

	nodeId := int64(1)

	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		panic(err)
	}

	return &PlaneLite{
		eventbus:    NewEventBus(),
		locker:      NewLocker(),
		router:      nil,
		idGenerator: node,
		nodeId:      nodeId,
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

func (p *PlaneLite) NewUId() int64 {
	id := p.idGenerator.Generate()
	return id.Int64()
}

func (p *PlaneLite) GetNodeId() int64 {
	return p.nodeId
}
