package runtime

import (
	"sync"

	"github.com/temphia/temphia/code/core/backend/engine/binders/standard"
	"github.com/temphia/temphia/code/core/backend/engine/rfencer"
	"github.com/temphia/temphia/code/core/backend/engine/runtime/rpool"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type ns struct {
	runtime  *runtime
	tenantId string
	running  map[string]*standard.Binder
	rlock    sync.Mutex // only using as pointer(no copy after first use) so its fine
	pool     rpool.Pool
	fencer   rfencer.RFencer
}

func (r *runtime) newNs(tenantId string) *ns {
	n := &ns{
		runtime:  r,
		tenantId: tenantId,
		running:  map[string]*standard.Binder{},
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

func (n *ns) listRunning() []etypes.RunningExec {
	re := make([]etypes.RunningExec, len(n.running))
	for _, b := range n.running {

		re = append(re, etypes.RunningExec{
			EventId:  b.Handle.EventId,
			BprintId: b.Handle.BprintId,
			PlugId:   b.Handle.PlugId,
			AgentId:  b.Handle.AgentId,
		})

	}
	return re
}

func (n *ns) countUp(b *standard.Binder) {
	n.rlock.Lock()
	defer n.rlock.Unlock()

	n.running[b.Handle.EventId] = b
}

func (n *ns) countDown(b *standard.Binder) {
	n.rlock.Lock()
	defer n.rlock.Unlock()

	delete(n.running, b.Handle.EventId)
}
