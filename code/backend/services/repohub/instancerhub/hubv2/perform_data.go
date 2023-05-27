package hubv2

import (
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

func (h *HubV2) instanceData(handle Handle, item xbprint.BundleItem) error {

	dsrc := h.dtable.DefaultSource(handle.tenantId)

	schema, err := h.readSchema(handle.tenantId, handle.opts.BprintId, item.File)
	if err != nil {
		return err
	}

	gslug := item.Name + handle.opts.InstanceId

	err = dsrc.MigrateSchema(handle.tenantId, step.MigrateOptions{
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
