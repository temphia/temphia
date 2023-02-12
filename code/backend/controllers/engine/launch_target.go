package engine

import (
	"github.com/temphia/temphia/code/backend/engine/invokers"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
)

type TargetLaunchData struct {
	TargetId   int64  `json:"target_id,omitempty"`
	TargetType string `json:"target_type,omitempty"`
}

func (c *Controller) LaunchTarget(uclaim *claim.Session, data TargetLaunchData) (*ExecInstanceOptions, error) {
	return c.launchTarget(uclaim, data)
}

func (c *Controller) launchTarget(uclaim *claim.Session, data TargetLaunchData) (*ExecInstanceOptions, error) {

	target, err := c.corehub.GetTargetApp(uclaim.TenantId, data.TargetType, data.TargetId)
	if err != nil {
		return nil, err
	}

	agent, err := c.corehub.AgentGet(uclaim.TenantId, target.PlugId, target.AgentId)
	if err != nil {
		return nil, err
	}

	token, err := c.signer.SignExecutor(uclaim.TenantId, &claim.Executor{
		TenantId:   uclaim.TenantId,
		UserId:     uclaim.UserID,
		UserGroup:  uclaim.UserGroup,
		DeviceId:   uclaim.DeviceId,
		Type:       "executor",
		SessionId:  uclaim.SessionID,
		ExecId:     c.idgen.Generate().Int64(),
		PlugId:     target.PlugId,
		AgentId:    target.AgentId,
		ExecType:   invokers.TypeUserApp,
		Attributes: map[string]string{},
	})
	if err != nil {
		return nil, err
	}

	return &ExecInstanceOptions{
		Token:        token,
		EntryName:    agent.WebEntry,
		ExecLoader:   agent.WebLoader,
		JSPlugScript: agent.WebScript,
		StyleFile:    agent.WebStyle,
		Plug:         target.PlugId,
		Agent:        target.AgentId,
		ExtScripts:   nil,
	}, nil
}
