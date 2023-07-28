package engine

import (
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
)

func (e *Engine) execute(opts etypes.Execution) ([]byte, error) {
	e.logger.Debug().Str("tenant_id", opts.TenantId).
		Str("plug_id", opts.PlugId).
		Str("agent_id", opts.AgentId).
		Str("action", opts.Action).
		Msg(logid.EngineExecAction)

	j := &job.Job{
		Namespace:         opts.TenantId,
		PlugId:            opts.PlugId,
		AgentId:           opts.AgentId,
		EventId:           xid.New().String(),
		EventAction:       opts.Action,
		Payload:           opts.Payload,
		PendingPrePolicy:  true,
		PendingPostPolicy: true,
		Plug:              nil,
		Agent:             nil,
		Loaded:            false,
		Invoker:           opts.Invoker,
	}

	eresp, err := e.runtime.Preform(j)
	if err != nil {
		e.logger.Debug().Str("tenant_id", opts.TenantId).
			Str("plug_id", opts.PlugId).
			Str("agent_id", opts.AgentId).
			Str("action", opts.Action).
			Err(err).
			Msg(logid.EngineExecActionErr)

		return nil, err
	}

	return eresp.Payload, nil

}
