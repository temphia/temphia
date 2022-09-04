package runtime

import (
	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/core/backend/engine/binders/standard"
	"github.com/temphia/temphia/code/core/backend/engine/runtime/loader"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"

	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"

	"github.com/ztrue/tracerr"
)

func (n *ns) getBinder(j *job.Job) (*standard.Binder, error) {
	if !j.Loaded {
		err := n.loadJob(j)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}
	}

	excbinder, left := n.pool.Borrow(j.PlugId, j.AgentId)
	if excbinder != nil {
		pp.Println("@left_in_pool", left)

		excbinder.AttachJob(j)

		return excbinder, nil
	}

	eb, ok := n.runtime.execBuilders[j.Agent.Executor]
	if !ok {
		return nil, easyerr.Error("Executor builder not found")
	}

	bind := n.runtime.binderFactory.New(standard.BinderOptions{
		Namespace: j.Namespace,
		PlugId:    j.PlugId,
		AgentId:   j.PlugId,
		BprintId:  j.Plug.BprintId,
		Epoch:     0,
	})

	bind.AttachJob(j)

	exec, err := eb.Instance(etypes.ExecutorOption{
		Binder:   bind,
		PlugId:   j.PlugId,
		AgentId:  j.AgentId,
		ExecType: j.Agent.Executor,
		TenantId: j.Namespace,
		EnvVars:  nil,
		File:     "",
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	bind.SetExec(exec)
	return bind, nil
}

func (n *ns) setBinder(j *job.Job, bind *standard.Binder) {
	n.pool.Return(bind)
}

func (n *ns) loadJob(j *job.Job) error {

	data, err := loader.Load(n.runtime.syncer, j.Namespace, j.PlugId, j.AgentId)

	if err != nil {
		return tracerr.Wrap(err)
	}
	j.Agent = data.Agent
	j.Plug = data.Plug
	j.Loaded = true
	return nil
}
