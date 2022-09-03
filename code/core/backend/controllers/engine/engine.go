package engine

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
)

type Controller struct {
	engine etypes.Engine
}

func New(engine etypes.Engine) *Controller {
	return &Controller{
		engine: engine,
	}
}
