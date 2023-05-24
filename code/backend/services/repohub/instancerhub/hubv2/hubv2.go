package hubv2

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
)

var (
	_ repox.InstancerHubV2 = (*HubV2)(nil)
)

type HubV2 struct {
	pacman repox.RepoBprintOps
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

	for _, item := range schema.Items {

		pp.Println(item)

	}

	// opts.File

	return nil, nil
}

func (h *HubV2) instanceData() {

}

func (h *HubV2) instancePlug() {

}

func (h *HubV2) Upgrade(opts repox.UpdateOptionsV2) error {

	return nil
}

func (h *HubV2) InstanceSheetDirect(opts repox.InstanceSheetOptions) (*xinstance.Response, error) {

	return nil, nil
}
