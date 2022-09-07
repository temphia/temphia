package engine

import "github.com/temphia/temphia/code/core/backend/xtypes/models/claim"

type User struct {
	PlugId  string
	AgentId string
}

func (c *Controller) launchUser(uclaim *claim.Session, data User) (string, error) {
	return "", nil

}
