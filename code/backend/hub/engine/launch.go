package enginehub

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/engine/invokers"
	"github.com/temphia/temphia/code/backend/engine/invokers/bundled"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

func (e *EngineHub) launchTarget(uclaim *claim.Session, data etypes.TargetLaunchData) (*etypes.LaunchOptions, error) {

	target, err := e.corehub.GetTargetApp(uclaim.TenantId, data.TargetType, data.TargetId)
	if err != nil {
		return nil, err
	}

	agent, err := e.corehub.AgentGet(uclaim.TenantId, target.PlugId, target.AgentId)
	if err != nil {
		return nil, err
	}

	token, err := e.signer.SignExecutor(uclaim.TenantId, &claim.Executor{
		TenantId:   uclaim.TenantId,
		UserId:     uclaim.UserID,
		UserGroup:  uclaim.UserGroup,
		DeviceId:   uclaim.DeviceId,
		Type:       "executor",
		SessionId:  uclaim.SessionID,
		ExecId:     e.idgen.Generate().Int64(),
		PlugId:     target.PlugId,
		AgentId:    target.AgentId,
		ExecType:   invokers.UserApp,
		TargetId:   target.Id,
		Attributes: map[string]string{},
	})
	if err != nil {
		return nil, err
	}

	return &etypes.LaunchOptions{
		Token:        token,
		EntryName:    agent.WebFiles["entry"],
		ExecLoader:   agent.WebFiles["loader"],
		JSPlugScript: agent.WebFiles["script"],
		StyleFile:    agent.WebFiles["style"],
		Plug:         target.PlugId,
		Agent:        target.AgentId,
		ExtScripts:   nil,
	}, nil
}

// admin

func (e *EngineHub) launchAdmin(uclaim *claim.Session, data etypes.AdminLaunchData) (*etypes.LaunchOptions, error) {

	token, err := e.signer.SignExecutor(uclaim.TenantId, &claim.Executor{
		TenantId:   uclaim.TenantId,
		UserId:     uclaim.UserID,
		UserGroup:  uclaim.UserGroup,
		DeviceId:   uclaim.DeviceId,
		Type:       "executor",
		SessionId:  uclaim.SessionID,
		ExecId:     e.idgen.Generate().Int64(),
		PlugId:     data.PlugId,
		AgentId:    data.AgentId,
		ExecType:   invokers.Admin,
		TargetId:   0,
		Attributes: make(map[string]string),
	})
	if err != nil {
		return nil, err
	}

	agent, err := e.corehub.AgentGet(uclaim.TenantId, data.PlugId, data.AgentId)
	if err != nil {
		return nil, err
	}

	return &etypes.LaunchOptions{
		Token:        token,
		EntryName:    agent.WebFiles["entry"],
		ExecLoader:   agent.WebFiles["loader"],
		JSPlugScript: agent.WebFiles["script"],
		StyleFile:    agent.WebFiles["style"],
		Plug:         data.PlugId,
		Agent:        data.AgentId,
		ExtScripts:   nil,
	}, nil

}

func (e *EngineHub) executeDev(dclaim *claim.UserContext, plug, agent, action string, body []byte) ([]byte, error) {
	// fixme => ability to send arbitary invoker type or dev invoker type?

	return e.engine.Execute(etypes.Execution{
		TenantId: dclaim.TenantId,
		PlugId:   plug,
		AgentId:  agent,
		Action:   action,
		Payload:  body,
		Invoker:  bundled.NewAdmin(dclaim),
	})
}

func (e *EngineHub) launchTargetDomain(tenantId, host, plugId, agentId string) (*etypes.LaunchDomainOptions, error) {

	agent, err := e.corehub.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		return nil, err
	}

	aburl := httpx.ApiBaseURL(host, tenantId)

	data := etypes.BootData{
		ApiBaseURL: aburl,
		PlugId:     plugId,
		AgentId:    agentId,
		// EntryName:  agent.WebEntry,
		// ExecLoader: agent.WebLoader,
	}

	bdata, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	return &etypes.LaunchDomainOptions{
		ApiBaseURL: aburl,
		// ScriptFile:   agent.WebScript,
		// StyleFile:    agent.WebStyle,
		LoaderScript: "",
		PlugId:       plugId,
		AgentId:      agentId,
		// ExecLoader:   agent.WebLoader,
		BootData:   string(bdata),
		ExtScripts: map[string]string{},
		Name:       agent.Name,
	}, nil
}
