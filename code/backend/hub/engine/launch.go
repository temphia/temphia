package enginehub

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/launch"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (e *EngineHub) launchAgent(uclaim *claim.Session, plugId, agentId string) (*launch.Response, error) {

	domain := fmt.Sprintf("%s-n-%s.%s", plugId, agentId, e.runnerDomain)

	token, err := e.signer.SignExecutor(uclaim.TenantId, &claim.Executor{
		TenantId:  uclaim.TenantId,
		UserId:    uclaim.UserID,
		UserGroup: uclaim.UserGroup,
		DeviceId:  uclaim.DeviceId,
		SessionId: uclaim.SessionID,
		PlugId:    plugId,
		AgentId:   agentId,
	})
	if err != nil {
		return nil, err
	}

	return &launch.Response{
		ApiBaseURL: httpx.ApiBaseURL(domain, uclaim.TenantId),
		Domain:     domain,
		Token:      token,
		PlugId:     plugId,
		AgentId:    agentId,
		TenantId:   uclaim.TenantId,
		StartPage:  "",
	}, nil
}

func (e *EngineHub) launchTarget(uclaim *claim.Session, targetId int64) (*launch.Response, error) {

	targets, err := e.corehub.ListTargetApp(uclaim.TenantId, map[string]any{
		"id": targetId,
	})
	if err != nil {
		return nil, err
	}

	plugId := targets[0].PlugId
	agentId := targets[0].AgentId

	domain := fmt.Sprintf("%s-n-%s.%s", plugId, agentId, e.runnerDomain)

	token, err := e.signer.SignExecutor(uclaim.TenantId, &claim.Executor{
		TenantId:  uclaim.TenantId,
		UserId:    uclaim.UserID,
		UserGroup: uclaim.UserGroup,
		DeviceId:  uclaim.DeviceId,
		SessionId: uclaim.SessionID,
		PlugId:    plugId,
		AgentId:   agentId,
	})
	if err != nil {
		return nil, err
	}

	return &launch.Response{
		ApiBaseURL: httpx.ApiBaseURL(domain, uclaim.TenantId),
		Domain:     domain,
		Token:      token,
		PlugId:     plugId,
		AgentId:    agentId,
		TenantId:   uclaim.TenantId,
		StartPage:  "",
	}, nil

}

func (e *EngineHub) launchEditor(uclaim *claim.Session, plugId, agentId string) (*launch.Response, error) {

	return nil, nil
}
