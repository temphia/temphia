package hubv2

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

var (
	_ repox.InstancerHubV2 = (*HubV2)(nil)
)

type HubV2 struct {
	pacman repox.RepoBprintOps
	dtable dyndb.DataHub
}

func New(pacman repox.RepoBprintOps, dtable dyndb.DataHub) *HubV2 {
	return &HubV2{
		pacman: pacman,
		dtable: dtable,
	}
}

func (h *HubV2) Instance(opts repox.InstanceOptionsV2) (*repox.InstanceResponseV2, error) {

	out, err := h.pacman.BprintGetBlob(opts.UserSession.TenantId, opts.BprintId, "schema.json")
	if err != nil {
		return nil, err
	}

	schema := xbprint.BundleV1{}
	err = json.Unmarshal(out, &schema)
	if err != nil {
		return nil, err
	}

	handle := Handle{
		items: make(map[string]string),
		opts:  opts,
	}

	for _, item := range schema.Items {

		switch item.Type {
		case xbprint.TypeDataGroup:
			err = h.instanceData(handle, item)
			if err != nil {
				return nil, err
			}

		case xbprint.TypePlug:
			err = h.instancePlug(handle, item)
			if err != nil {
				return nil, err
			}

		default:
			return nil, easyerr.NotImpl()
		}
	}

	return nil, nil
}

func (h *HubV2) Upgrade(opts repox.UpdateOptionsV2) error {

	return nil
}

func (h *HubV2) InstanceSheetDirect(opts repox.InstanceSheetOptions) (*xinstance.Response, error) {

	return nil, nil
}
