package sheet

import (
	"github.com/temphia/temphia/code/backend/services/repohub/instancers/dtable"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

type DataInstancer interface {
	DirectInstance(tenantId string, opts *dtable.DataGroupRequest, schema *xbprint.NewTableGroup) (*dtable.DataGroupResponse, error)
}
