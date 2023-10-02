package runtime

import (
	"sync"

	"github.com/rs/zerolog"

	"github.com/temphia/temphia/code/backend/engine/binder"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
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
	binderFactory binder.Factory

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
	r.binderFactory = binder.NewFactory(binder.FactoryOptions{
		App:          r.app,
		Logger:       r.logger,
		Modules:      modules,
		ExecBuilders: r.execBuilders,
		Runtime:      r,
	})

	go r.controlLoop()
	return nil
}

func (r *runtime) InitAgent(tenantId, plug, agent string) error {
	return nil
}

func (r *runtime) Preform(j *job.Job) (*event.Response, error) {
	ns := r.getNS(j.Namespace, true)

	return ns.doWork(j)

}

func (r *runtime) PreformAsync(j *job.AsyncJob) {
	r.jobCh <- j
}

func (r *runtime) ResetAgents(tenantId, plug string, agents []string) {
	ns := r.getNS(tenantId, false)
	if ns == nil {
		return
	}

	ns.destroy(plug, agents)
}

func (r *runtime) ResetBprint(tenantId, bprint string) {

}

func (r *runtime) ListRunning(tenantId string) ([]etypes.RunningExec, error) {
	ns := r.getNS(tenantId, false)
	if ns == nil {
		return nil, easyerr.NotFound("tenant runtime")
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

	ns = r.newNs(tenantId)
	r.ns[tenantId] = ns

	return ns
}
