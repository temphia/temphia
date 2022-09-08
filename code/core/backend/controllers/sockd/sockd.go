package sockd

import "github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"

type Controller struct {
	sockd sockdx.Sockd
}

func New(sockd sockdx.Sockd) *Controller {
	return &Controller{
		sockd: sockd,
	}
}
