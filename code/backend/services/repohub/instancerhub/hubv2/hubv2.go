package hubv2

import (
	"encoding/json"

	"github.com/jaevor/go-nanoid"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

var (
	_        repox.InstancerHubV2 = (*HubV2)(nil)
	gFunc, _                      = nanoid.CustomASCII("abcdefghijklmnopqrstuvwxyz1234567890", 5)
)

type HubV2 struct {
	pacman  repox.RepoBprintOps
	dtable  dyndb.DataHub
	corehub store.CoreHub
}

func New(pacman repox.RepoBprintOps, dtable dyndb.DataHub, corehub store.CoreHub) *HubV2 {
	return &HubV2{
		pacman:  pacman,
		dtable:  dtable,
		corehub: corehub,
	}
}

func (h *HubV2) Instance(opts repox.InstanceOptionsV2) (*repox.InstanceResponseV2, error) {

	if opts.InstanceId == "" {
		opts.InstanceId = gFunc()
	}

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
		tenantId:   opts.UserSession.TenantId,
		dataSource: "",
		dataGroups: make(map[string]string),
		plugs:      make(map[string]string),
		resources:  make(map[string]string),
		targets:    make(map[string]int64),
		opts:       opts,
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
		case xbprint.TypeResource:
			err = h.instanceResource(handle, item)
			if err != nil {
				return nil, err
			}

		case xbprint.TypeTargetApp:
			err = h.instanceTargetApp(handle, item)
			if err != nil {
				return nil, err
			}

		case xbprint.TypeTargetHook:
			err = h.instanceTargetHook(handle, item)
			if err != nil {
				return nil, err
			}

		case xbprint.TypeDataSheet:
			fallthrough

		default:
			return nil, easyerr.NotImpl()
		}
	}

	return nil, nil
}

func (h *HubV2) Upgrade(opts repox.UpdateOptionsV2) error {

	// h.corehub.ListTargetApp(opts.UserSession.TenantId, map[string]any{
	// 	"bprint_id":          opts.BprintId,
	// 	"bprint_instance_id": opts.InstanceId,
	// })

	return nil
}

func (h *HubV2) InstanceSheetDirect(opts repox.InstanceSheetOptions) (*xinstance.Response, error) {

	return nil, nil
}

// private

func (h *HubV2) readSchema(tenantId, bprintid, file string) (*step.Schema, error) {
	out, err := h.pacman.BprintGetBlob(tenantId, bprintid, file)
	if err != nil {
		return nil, err
	}

	steps := step.Schema{}
	err = json.Unmarshal(out, &steps)
	if err != nil {
		return nil, err
	}

	return &steps, nil
}
