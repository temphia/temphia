package engine

import (
	"github.com/temphia/temphia/code/backend/engine/invokers"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type AdminLaunchData struct {
	PlugId  string `json:"plug_id,omitempty"`
	AgentId string `json:"agent_id,omitempty"`
}

func (c *Controller) LaunchAdmin(uclaim *claim.Session, data AdminLaunchData) (*ExecInstanceOptions, error) {
	return c.launchAdmin(uclaim, data)
}

func (c *Controller) launchAdmin(uclaim *claim.Session, data AdminLaunchData) (*ExecInstanceOptions, error) {

	token, err := c.signer.SignExecutor(uclaim.TenantId, &claim.Executor{
		TenantId:   uclaim.TenantId,
		UserId:     uclaim.UserID,
		UserGroup:  uclaim.UserGroup,
		DeviceId:   uclaim.DeviceId,
		Type:       "executor",
		SessionId:  uclaim.SessionID,
		ExecId:     c.idgen.Generate().Int64(),
		PlugId:     data.PlugId,
		AgentId:    data.AgentId,
		ExecType:   invokers.Admin,
		TargetId:   0,
		Attributes: make(map[string]string),
	})
	if err != nil {
		return nil, err
	}

	agent, err := c.corehub.AgentGet(uclaim.TenantId, data.PlugId, data.AgentId)
	if err != nil {
		return nil, err
	}

	return &ExecInstanceOptions{
		Token:        token,
		EntryName:    agent.WebEntry,
		ExecLoader:   agent.WebLoader,
		JSPlugScript: agent.WebScript,
		StyleFile:    agent.WebStyle,
		Plug:         data.PlugId,
		Agent:        data.AgentId,
		ExtScripts:   nil,
	}, nil

}

func (c *Controller) ExecuteDev(dclaim *claim.PlugDevTkt, plug, agent, action string, body []byte) ([]byte, error) {

	// fixme => check tkt and user perms here
	// fixme => ability to send arbitary invoker type or dev invoker type?

	return c.engine.Execute(etypes.Execution{
		TenantId: dclaim.TenantId,
		PlugId:   plug,
		AgentId:  agent,
		Action:   action,
		Payload:  body,
		Invoker:  nil, //web.NewWeb(ctx, eclaim),
	})
}
