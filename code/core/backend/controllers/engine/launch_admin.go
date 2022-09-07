package engine

import "github.com/temphia/temphia/code/core/backend/xtypes/models/claim"

type Admin struct {
	PlugId  string
	AgentId string
}

func (c *Controller) launchAdmin(uclaim *claim.Session, data Admin) (string, error) {
	return "", nil
}
