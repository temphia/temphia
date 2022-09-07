package engine

import "github.com/temphia/temphia/code/core/backend/xtypes/models/claim"

type Data struct {
	PlugId  string
	AgentId string
}

func (c *Controller) launchData(uclaim *claim.Session, data Data) (string, error) {
	return "", nil
}
