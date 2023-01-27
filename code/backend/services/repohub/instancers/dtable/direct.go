package dtable

import "github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"

func (di *dtabeInstancer) DirectInstance(tenantId string, opts *DataGroupRequest, schema *xbprint.NewTableGroup) (*DataGroupResponse, error) {
	return di.instance(tenantId, opts, schema)
}
