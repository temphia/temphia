package runtime

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/thoas/go-funk"

	"github.com/temphia/temphia/code/core/backend/engine/binders/standard"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

var (
	_ etypes.Runtime = (*runtime)(nil)
)

type runtime struct {
	close         chan struct{}
	jobCh         chan *job.AsyncJob
	router        etypes.Router
	nodeTags      []string
	app           xtypes.App
	binderFactory standard.Factory

	execBuilders map[string]etypes.ExecutorBuilder
	ns           map[string]*ns
	nlock        sync.Mutex

	//ext services
	signer service.Signer
	syncer store.SyncDB
	logger zerolog.Logger
}

func New(_app xtypes.App, logger zerolog.Logger) *runtime {

	//	reg := _app.Registry().(*registry.Registry)

	deps := _app.GetDeps()

	rt := &runtime{
		close:        make(chan struct{}),
		jobCh:        make(chan *job.AsyncJob),
		router:       nil,
		app:          _app,
		execBuilders: nil,
		nodeTags:     []string{},
		signer:       deps.Signer().(service.Signer),
		syncer:       deps.CoreHub().(store.SyncDB),
		logger:       logger,
		ns:           make(map[string]*ns),
		nlock:        sync.Mutex{},
	}

	return rt
}

func (r *runtime) Run(builders map[string]etypes.ExecutorBuilder, modules map[string]etypes.ModuleBuilder) error {
	r.execBuilders = builders
	r.binderFactory = standard.NewFactory(standard.FactoryOptions{
		App:          r.app,
		Logger:       r.logger,
		Modules:      modules,
		ExecBuilders: r.execBuilders,
		Runtime:      r,
	})

	go r.controlLoop()
	return nil
}

func (r *runtime) Preform(j *job.Job) (*event.Response, error) {
	ns := r.getNS(j.Namespace, true)

	if j.PendingPrePolicy {
		err := ns.fencer.Execute(j)
		if err != nil {
			return nil, err
		}
	}

	if j.NodeTag == "" || funk.ContainsString(r.nodeTags, j.NodeTag) || r.router == nil {
		// fixme => run post_policy here
		return ns.doWork(j)
	}

	return r.router.Route(j)
}

func (r *runtime) PreformAsync(j *job.AsyncJob) {
	r.jobCh <- j
}

func (r *runtime) ResetAgents(tenantId, plug string, agents []string) {

}

func (r *runtime) ResetBprint(tenantId, bprint string) {

}

func (r *runtime) ListRunning(tenantId string) ([]etypes.RunningExec, error) {
	ns := r.getNS(tenantId, false)
	if ns == nil {
		return nil, easyerr.NotFound()
	}

	return ns.listRunning(), nil
}

// private

func (r *runtime) getNS(tenantId string, create bool) *ns {
	r.nlock.Lock()
	defer r.nlock.Unlock()

	ns, ok := r.ns[tenantId]
	if ok {
		return ns
	}

	if r.app.SingleTenant() && tenantId != r.app.TenantId() {
		panic("Wrong Tenant")
	}

	if !create {
		return nil
	}

	ns = r.newNs(tenantId)
	r.ns[tenantId] = ns

	return ns
}
