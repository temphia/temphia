package hubv2

import "github.com/temphia/temphia/code/backend/xtypes/service/repox"

type Handle struct {
	tenantId   string
	dataSource string
	dataGroups map[string]string
	plugs      map[string]string
	resources  map[string]string
	targets    map[string]int64

	opts repox.InstanceOptionsV2
}
