package hubv2

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

func (h *HubV2) instanceTargetApp(handle Handle, item xbprint.BundleItem) error {

	schema, err := h.readSchema(handle.tenantId, handle.opts.BprintId, item.File)
	if err != nil {
		return err
	}

	for _, s := range schema.Steps {

		switch s.Name {

		case step.PlugStepAddTargetApp:
			// resolve target and plug_id

			data := &entities.TargetApp{}
			err = json.Unmarshal(s.Data, data)
			if err != nil {
				return err
			}

			_, err = h.corehub.AddTargetApp(data)
			if err != nil {
				return err
			}

		case step.PlugStepUpdateTargetApp:
			fallthrough
		case step.PlugStepDeleteTargetApp:
			fallthrough
		default:

			return easyerr.NotImpl()
		}

		pp.Println(s)

	}

	return nil
}

func (h *HubV2) instanceTargetHook(handle Handle, item xbprint.BundleItem) error {

	schema, err := h.readSchema(handle.tenantId, handle.opts.BprintId, item.File)
	if err != nil {
		return err
	}

	for _, s := range schema.Steps {

		pp.Println(s)

		switch s.Name {
		case step.PlugStepAddTargetHook:
			// resolve target and plug_id

			data := &entities.TargetHook{}
			err = json.Unmarshal(s.Data, data)
			if err != nil {
				return err
			}

			_, err = h.corehub.AddTargetHook(data)
			if err != nil {
				return err
			}

		case step.PlugStepUpdateTargetHook:
			fallthrough
		case step.PlugStepDeleteTargetHook:
			fallthrough
		default:
			return easyerr.NotImpl()
		}

	}

	return nil
}
