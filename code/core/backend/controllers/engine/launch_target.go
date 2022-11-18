package engine

import (
	"github.com/temphia/temphia/code/core/backend/engine/invokers"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/vmodels"
)

type TargetLaunchData struct {
	TargetId   int64  `json:"target_id,omitempty"`
	TargetType string `json:"target_type,omitempty"`
}

func (c *Controller) LaunchTarget(uclaim *claim.Session, data TargetLaunchData) (*vmodels.LoaderOptions, error) {
	return c.launchTarget(uclaim, data)
}

func (c *Controller) launchTarget(uclaim *claim.Session, data TargetLaunchData) (*vmodels.LoaderOptions, error) {

	target, err := c.corehub.GetTargetApp(uclaim.TenentId, data.TargetType, data.TargetId)
	if err != nil {
		return nil, err
	}

	agent, err := c.corehub.AgentGet(uclaim.TenentId, target.PlugId, target.AgentId)
	if err != nil {
		return nil, err
	}

	token, err := c.signer.SignExecutor(uclaim.TenentId, &claim.Executor{
		TenentId:   uclaim.TenentId,
		UserId:     uclaim.UserID,
		UserGroup:  uclaim.UserGroup,
		DeviceId:   "",
		Type:       "executor",
		SessionId:  uclaim.SessionID,
		ExecId:     0,
		PlugId:     target.PlugId,
		AgentId:    target.AgentId,
		ExecType:   invokers.TypeUserApp,
		Attributes: map[string]string{},
	})
	if err != nil {
		return nil, err
	}

	return &vmodels.LoaderOptions{
		BaseURL:      "",
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
