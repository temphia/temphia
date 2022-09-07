package engine

import "github.com/temphia/temphia/code/core/backend/xtypes/models/claim"

type Domain struct {
	PlugId  string
	AgentId string
}

func (c *Controller) launchDomain(uclaim *claim.Session, data Domain) (string, error) {

	return "", nil

	// DomainEditor

}
