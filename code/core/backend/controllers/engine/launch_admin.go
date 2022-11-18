package engine

import (
	"github.com/temphia/temphia/code/core/backend/engine/invokers"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

type AdminLaunchData struct {
	PlugId  string
	AgentId string
}

func (c *Controller) LaunchAdmin(uclaim *claim.Session, data AdminLaunchData) (string, error) {
	return c.launchAdmin(uclaim, data)
}

func (c *Controller) launchAdmin(uclaim *claim.Session, data AdminLaunchData) (string, error) {

	return c.signer.SignExecutor(uclaim.TenentId, &claim.Executor{
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

}
