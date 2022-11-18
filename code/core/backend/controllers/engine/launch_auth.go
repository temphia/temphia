package engine

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

type AuthLaunchData struct{}

func (c *Controller) LaunchAuth(uclaim *claim.Session, data AuthLaunchData) (string, error) {
	return c.launchAuth(uclaim, data)
}

func (c *Controller) launchAuth(uclaim *claim.Session, data AuthLaunchData) (string, error) {
	return "", nil
}
