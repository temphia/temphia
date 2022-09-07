package engine

import "github.com/temphia/temphia/code/core/backend/xtypes/models/claim"

type Widget struct {
	PlugId  string
	AgentId string
}

func (c *Controller) launchWidget(uclaim *claim.Session, data Widget) (string, error) {
	return "", nil
}
