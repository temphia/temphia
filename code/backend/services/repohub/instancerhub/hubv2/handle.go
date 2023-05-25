package hubv2

import "github.com/temphia/temphia/code/backend/xtypes/service/repox"

type Handle struct {
	dataSource string
	dataGroups map[string]string
	plugs      map[string]string
	opts       repox.InstanceOptionsV2
}
