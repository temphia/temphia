package engine

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
)

func (e *Engine) serveAgentFile(tenantId, plugId, agentId, file string) ([]byte, error) {

	plug, err := e.syncer.PlugGet(tenantId, plugId)
	if err != nil {
		e.logger.Debug().
			Str("tenant_id", tenantId).
			Str("plug_id", plugId).
			Str("agent_id", agentId).
			Err(err).
			Str("file", file).
			Msg(logid.EngineServePlugLoadError)

		return nil, err
	}

	agent, err := e.syncer.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		e.logger.Debug().
			Str("tenant_id", tenantId).
			Str("plug_id", plugId).
			Str("agent_id", agentId).
			Err(err).
			Str("file", file).
			Msg(logid.EngineServeAgentLoadError)

		return nil, err

	}

	actualFile := agent.WebFiles[file]
	if actualFile == "" {
		e.logger.Debug().
			Str("tenant_id", tenantId).
			Str("plug_id", plugId).
			Str("agent_id", agentId).
			Str("file", file).
			Msg(logid.EngineServeEmptyMappings)
		return nil, easyerr.NotFound("web file")
	}

	out, err := e.pacman.BprintGetBlob(tenantId, plug.BprintId, actualFile)
	if err != nil {
		e.logger.Debug().
			Str("tenant_id", tenantId).
			Str("plug_id", plugId).
			Str("agent_id", agentId).
			Err(err).
			Str("file", file).
			Msg(logid.EngineServeBprintErr)
		return nil, err
	}

	return out, err

}

func (e *Engine) serveExecutorFile(tenantId, plugId, agentId, file string) ([]byte, error) {

	agent, err := e.syncer.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		e.logger.Debug().
			Str("tenant_id", tenantId).
			Str("plug_id", plugId).
			Str("agent_id", agentId).
			Str("file", file).
			Msg(logid.EngineExecServeAgentLoadError)

	}

	builder, ok := e.execbuilders[agent.Executor]
	if !ok {
		e.logger.Debug().
			Str("tenant_id", tenantId).
			Str("plug_id", plugId).
			Str("agent_id", agentId).
			Str("file", file).
			Msg(logid.EngineExecServeExecBuilderNotFound)
	}

	out, err := builder.ExecFile(file)
	if err != nil {
		e.logger.Debug().
			Str("tenant_id", tenantId).
			Str("plug_id", plugId).
			Str("agent_id", agentId).
			Err(err).
			Str("file", file).
			Msg(logid.EngineExecServeExecBuilderErr)
		return nil, err
	}

	return out, nil
}
