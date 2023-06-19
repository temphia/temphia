package runtime

import (
	"sync"

	"github.com/temphia/temphia/code/backend/engine/binder"
	"github.com/temphia/temphia/code/backend/engine/rfencer"
	"github.com/temphia/temphia/code/backend/engine/runtime/rpool"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type ns struct {
	runtime  *runtime
	tenantId string
	running  map[string]*binder.Binder
	rlock    sync.Mutex // only using as pointer(no copy after first use) so its fine
	pool     rpool.Pool
	fencer   rfencer.RFencer
}

func (r *runtime) newNs(tenantId string) *ns {
	n := &ns{
		runtime:  r,
		tenantId: tenantId,
		running:  map[string]*binder.Binder{},
		rlock:    sync.Mutex{},
		pool:     rpool.NewPool(),
		fencer:   rfencer.New(tenantId, r.app.GetDeps().CoreHub().(store.CoreHub)),
	}

	return n
}

func (n *ns) doWork(j *job.Job) (*event.Response, error) {
	binder, err := n.getBinder(j)
	if err != nil {
		return nil, err
	}

	n.countUp(binder)
	resp, err := binder.Execute()
	n.countDown(binder)

	n.setBinder(j, binder)
	return resp, err
}

func (n *ns) destroy(plug string, agents []string) {
	n.pool.Destroy(plug, agents)
}

func (n *ns) listRunning() []etypes.RunningExec {
	re := make([]etypes.RunningExec, len(n.running))
	for _, b := range n.running {

		re = append(re, etypes.RunningExec{
			EventId:  b.EventId,
			BprintId: b.BprintId,
			PlugId:   b.PlugId,
			AgentId:  b.AgentId,
		})

	}
	return re
}

func (n *ns) countUp(b *binder.Binder) {
	n.rlock.Lock()
	defer n.rlock.Unlock()

	n.running[b.EventId] = b
}

func (n *ns) countDown(b *binder.Binder) {
	n.rlock.Lock()
	defer n.rlock.Unlock()

	delete(n.running, b.EventId)
}
