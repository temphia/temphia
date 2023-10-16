package engine

import (
	"path"

	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/engine/binder"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
)

func (e *Engine) rPXecute(opts etypes.RPXecuteOptions) ([]byte, error) {

	b := e.getBinding(opts.TenantId, opts.PlugId, opts.AgentId)

	j := &job.RPXJob{
		EventId:   "",
		Name:      opts.Action,
		Namespace: opts.TenantId,
		Payload:   opts.Payload,
		Invoker:   opts.Invoker,
		NodeTag:   "",
	}

	return b.RPXecute(j)

}

func (e *Engine) webRawXecute(opts etypes.WebRawXecuteOptions) {

	b := e.getBinding(opts.TenantId, opts.PlugId, opts.AgentId)

	b.WebRawXecute(&job.RawWebJob{
		EventId: "",
		Writer:  opts.Writer,
		Request: opts.Request,
	})
}

// private

func (e *Engine) getBinding(tenantId, plugId, agentId string) *binder.Binder {

	key := plugId + agentId

	e.rLock.RLock()
	b := e.running[key]
	e.rLock.RUnlock()

	if b != nil {
		return b
	}

	agent := e.ecache.GetAgent(tenantId, plugId, agentId)
	plug := e.ecache.GetPlug(tenantId, plugId)

	b = e.binderFactory.New(binder.BinderOptions{
		Namespace: tenantId,
		PlugId:    plugId,
		AgentId:   agentId,
		BprintId:  plug.BprintId,
		Epoch:     0,
	})

	eb := e.execbuilders[agent.Executor]

	cd := e.app.GetDeps().Confd().(config.Confd)
	ety := e.rundb.get(plugId, agentId)
	if ety == nil {
		iid, _ := xutils.GenerateRandomString(5)
		ety = &entry{
			BprintId: plug.BprintId,
			RunId:    iid,
			PlugId:   plugId,
			AgentId:  agentId,
		}

		e.rundb.set(ety)
		err := e.rundb.flush()
		if err != nil {
			return nil
		}
	}

	// check mod_hash if it changed? then generate new runid

	ex, err := eb.New(etypes.ExecutorOption{
		Binder:        b,
		TenantId:      tenantId,
		PlugId:        plugId,
		AgentId:       agentId,
		File:          agent.EntryFile,
		ExecType:      agent.Executor,
		BprintId:      plug.BprintId,
		DefaultRunner: "",
		RunFolder:     path.Join(cd.RootDataFolder(), ety.RunId),
		WebOptions:    agent.WebOptions,
	})

	if err != nil {
		return nil
	}

	b.Executor = ex

	e.rLock.Lock()
	e.running[key] = b
	e.rLock.Unlock()

	return b

}
