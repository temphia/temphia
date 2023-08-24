package enginehub

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/launch"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz/httpx"
)

func (e *EngineHub) launchAgent(uclaim *claim.Session, plugId, agentId string) (*launch.Response, error) {

	config := e.app.GetDeps().Confd().(config.Confd).GetConfig()
	domain := fmt.Sprintf("%s-n-%s.%s", plugId, agentId, config.RunnerDomain)

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
	}, nil
}

func (e *EngineHub) launchTarget(uclaim *claim.Session, targetId int64) (*launch.Response, error) {

	return nil, nil
}

func (e *EngineHub) launchEditor(uclaim *claim.Session, plugId, agentId string) (*launch.Response, error) {

	return nil, nil
}
