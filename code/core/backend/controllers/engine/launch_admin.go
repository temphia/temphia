package engine

import (
	"github.com/temphia/temphia/code/core/backend/engine/invoker"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

type Admin struct {
	PlugId  string
	AgentId string
}

func (c *Controller) launchAdmin(uclaim *claim.Session, data Admin) (string, error) {

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
		ExecType:   invoker.TypeWebAdmin,
		Attributes: make(map[string]string),
	})

}
