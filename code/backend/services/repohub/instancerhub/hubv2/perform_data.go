package hubv2

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

func (h *HubV2) instanceData(handle Handle, item xbprint.BundleItem) error {
	tenantId := handle.opts.UserSession.TenantId
	dsrc := h.dtable.DefaultSource(tenantId)

	out, err := h.pacman.BprintGetBlob(tenantId, handle.opts.BprintId, item.File)
	if err != nil {
		return err
	}

	schema := step.Schema{}
	err = json.Unmarshal(out, &schema)
	if err != nil {
		return err
	}

	gslug := item.Name + handle.opts.InstanceId

	err = dsrc.MigrateSchema(tenantId, step.MigrateOptions{
		Steps:            schema.Steps,
		New:              true,
		Gslug:            gslug,
		BprintId:         handle.opts.BprintId,
		BprintItemId:     item.Name,
		BprintInstanceId: handle.opts.InstanceId,
	})
	if err != nil {
		return err
	}

	handle.dataGroups[item.Name] = gslug

	return nil
}
