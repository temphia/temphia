package plane

import (
	"github.com/temphia/temphia/code/backend/plane/idservice"
	"github.com/temphia/temphia/code/backend/plane/msgbus"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

var _ xplane.ControlPlane = (*PlaneLite)(nil)

type PlaneLite struct {
	locker *Locker
	router *Router
	nodeId int64
	seq    idservice.IDService

	msgbus xplane.MsgBus
}

func NewLite(coreHub store.CoreHub) *PlaneLite {

	nodeId := int64(1)

	return &PlaneLite{
		locker: NewLocker(),
		router: nil,
		nodeId: nodeId,
		seq:    *idservice.New(nodeId),
		msgbus: msgbus.New(nodeId, coreHub),
	}
}

// dl and start up stuff

func (p *PlaneLite) Start() error {
	go p.msgbus.Start()
	return nil
}

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

func (p *PlaneLite) GetMsgBus() xplane.MsgBus {
	return p.msgbus
}

func (p *PlaneLite) GetNodeId() int64 {
	return p.nodeId
}

func (p *PlaneLite) GetIdService() xplane.IDService { return &p.seq }
