package hubv2

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities/resource"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

func (h *HubV2) instanceResource(handle Handle, item xbprint.BundleItem) error {

	schema, err := h.readSchema(handle.tenantId, handle.opts.BprintId, item.File)
	if err != nil {
		return err
	}

	for _, s := range schema.Steps {

		switch s.Type {
		case step.PlugStepNewResourceModule:
			data := xbprint.NewResource{}
			err = json.Unmarshal(s.Data, &data)
			if err != nil {
				return err
			}

			err = h.applyNewResource(item.Name, handle, data)
			if err != nil {
				return err
			}

		case step.PlugStepUpdateResourceModule:
			// data := make(map[string]any)
			// err = json.Unmarshal(s.Data, &data)
			// if err != nil {
			// 	return err
			// }
			// h.syncdb.ResourceUpdate(handle.tenantId, )

			fallthrough

		case step.PlugStepRemoveResourceModule:
			fallthrough
		case step.PlugStepAddResourceLink:
			fallthrough
		case step.PlugStepRemoveResourceLink:
			fallthrough
		default:
			return easyerr.NotImpl()
		}

	}

	return nil
}

func (h *HubV2) applyNewResource(item string, handle Handle, data xbprint.NewResource) error {

	// resolve target here
	target := data.TargetRef

	switch data.Type {
	// case resource.SockRoom:
	// case resource.Folder:
	// case resource.UserGroup:
	// case resource.Module:

	case resource.DataGroup:
		tparts := strings.Split(data.TargetRef, "/")

		if len(tparts) != 2 {
			return easyerr.Error("target has invalid value")
		}

		target = fmt.Sprintf("%s/%s/%s", handle.dataSource, handle.dataGroups[tparts[0]], tparts[1])
	}

	return h.corehub.ResourceNew(handle.tenantId, &entities.Resource{
		TenantId: handle.tenantId,
		Name:     data.Name,
		Type:     data.Type,
		SubType:  data.SubType,
		Target:   target,
		Payload:  data.Payload,
		Policy:   data.Policy,
		// BprintId:         handle.opts.BprintId,
		// BprintItemId:     item,
		// BprintInstanceId: handle.opts.InstanceId,
	})

}
