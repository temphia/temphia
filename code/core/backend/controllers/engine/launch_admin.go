package engine

import (
	"github.com/temphia/temphia/code/core/backend/engine/invokers"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/vmodels"
)

type AdminLaunchData struct {
	PlugId  string
	AgentId string
}

func (c *Controller) LaunchAdmin(uclaim *claim.Session, data AdminLaunchData) (*vmodels.LoaderOptions, error) {
	return c.launchAdmin(uclaim, data)
}

func (c *Controller) launchAdmin(uclaim *claim.Session, data AdminLaunchData) (*vmodels.LoaderOptions, error) {

	token, err := c.signer.SignExecutor(uclaim.TenentId, &claim.Executor{
		TenentId:   uclaim.TenentId,
		UserId:     uclaim.UserID,
		UserGroup:  uclaim.UserGroup,
		DeviceId:   uclaim.DeviceId,
		Type:       "executor",
		SessionId:  uclaim.SessionID,
		ExecId:     0,
		PlugId:     data.PlugId,
		AgentId:    data.AgentId,
		ExecType:   invokers.TypeWebAdmin,
		Attributes: make(map[string]string),
	})
	if err != nil {
		return nil, err
	}

	agent, err := c.corehub.AgentGet(uclaim.TenentId, data.PlugId, data.AgentId)
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
		Plug:         data.PlugId,
		Agent:        data.AgentId,
		ExtScripts:   nil,
	}, nil

}
