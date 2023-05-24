package hubv2

import "github.com/temphia/temphia/code/backend/xtypes/service/repox"

type Handle struct {
	items  map[string]string
	opts   repox.InstanceOptionsV1
	pacman repox.RepoBprintOps
}
