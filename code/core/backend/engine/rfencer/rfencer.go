package rfencer

import (
	"sync"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type RFencer struct {
	tenantId string
	corehub  store.CoreHub
	programs map[string]*vm.Program
	pool     sync.Pool
	mlock    sync.Mutex
}

func New(tenantId string, corehub store.CoreHub) *RFencer {
	return &RFencer{
		tenantId: tenantId,
		corehub:  corehub,
		programs: make(map[string]*vm.Program),
		pool: sync.Pool{
			New: func() any {
				return &vm.VM{}
			},
		},
		mlock: sync.Mutex{},
	}
}

func (r *RFencer) get(plugId string) *vm.Program {

	r.mlock.Lock()
	pg, ok := r.programs[plugId]
	r.mlock.Unlock()
	if ok {
		return pg
	}

	plug, err := r.corehub.PlugGet(r.tenantId, plugId)
	if err != nil {
		return nil
	}

	newpg, err := expr.Compile(plug.InvokePolicy, expr.Env(Env{}))
	if err != nil {
		return nil
	}

	r.mlock.Lock()
	oldpg, ok := r.programs[plugId]

	if ok {
		r.mlock.Unlock()
		return oldpg
	}

	r.programs[plugId] = newpg
	r.mlock.Unlock()

	return newpg

}

func (r *RFencer) Execute(job *job.Job) error {
	pg := r.get(job.PlugId)
	if pg == nil {
		return easyerr.NotFound()
	}

	vm := r.pool.Get().(*vm.VM)

	_, err := vm.Run(pg, Env{
		Job: job,
	})

	return err
}
