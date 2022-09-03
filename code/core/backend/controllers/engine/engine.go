package engine

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/enginex"
)

type Controller struct {
	engine enginex.Engine
}

func New(engine enginex.Engine) *Controller {
	return &Controller{
		engine: engine,
	}
}
