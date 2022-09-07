package engine

import "github.com/temphia/temphia/code/core/backend/xtypes/models/claim"

type Authd struct {
	PlugId  string
	AgentId string
}

func (c *Controller) launchAuthd(uclaim *claim.Session, data Authd) (string, error) {
	return "", nil
}
