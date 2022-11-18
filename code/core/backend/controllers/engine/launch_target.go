package engine

import "github.com/temphia/temphia/code/core/backend/xtypes/models/claim"

type TargetLaunchData struct {
	TargetId int64
}

func (c *Controller) LaunchTarget(uclaim *claim.Session, data TargetLaunchData) (string, error) {
	return c.launchTarget(uclaim, data)
}

func (c *Controller) launchTarget(uclaim *claim.Session, data TargetLaunchData) (string, error) {
	return "", nil
}
