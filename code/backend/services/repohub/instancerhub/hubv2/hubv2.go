package hubv2

import (
	"encoding/json"

	"github.com/jaevor/go-nanoid"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
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
	pacman repox.RepoBprintOps
	dtable dyndb.DataHub
	syncdb store.SyncDB
}

func New(pacman repox.RepoBprintOps, dtable dyndb.DataHub, syncdb store.SyncDB) *HubV2 {
	return &HubV2{
		pacman: pacman,
		dtable: dtable,
		syncdb: syncdb,
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
		dataSource: "",
		dataGroups: make(map[string]string),
		plugs:      make(map[string]string),
		resources:  make(map[string]string),
		targets:    make(map[string]string),
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

		case xbprint.TypeDataSheet:
			fallthrough
		case xbprint.TypeResource:
			fallthrough
		case xbprint.TypeTarget:
			fallthrough
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
