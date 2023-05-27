package hubv2

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/tidwall/gjson"
)

func (h *HubV2) instancePlug(handle Handle, item xbprint.BundleItem) error {
	tenantId := handle.opts.UserSession.TenantId

	out, err := h.pacman.BprintGetBlob(tenantId, handle.opts.BprintId, item.File)
	if err != nil {
		return err
	}

	steps := step.Schema{}
	err = json.Unmarshal(out, &steps)
	if err != nil {
		return err
	}

	pid := item.Name + handle.opts.InstanceId

	for _, pstep := range steps.Steps {

		switch pstep.Name {
		case step.PlugStepNewPlug:
			pschema := xbprint.NewPlug{}
			err = json.Unmarshal(pstep.Data, &pschema)
			if err != nil {
				return err
			}

			err = h.applyNewPlug(tenantId, pid, item.Name, handle, pschema)
			if err != nil {
				return err
			}

		case step.PlugStepNewAgent:

			data := xbprint.NewAgent{}
			err = json.Unmarshal(pstep.Data, &data)
			if err != nil {
				return err
			}

			err = h.applyAddAgent(tenantId, pid, data)
			if err != nil {
				return err
			}

		case step.PlugStepUpdateAgent:
			data := make(map[string]any)
			err = json.Unmarshal(pstep.Data, &data)
			if err != nil {
				return err
			}

			return h.applyUpdateAgent(tenantId, pid, data)
		case step.PlugStepRemoveAgent:
			return h.applyRemoveAgent(
				tenantId,
				pid,
				gjson.GetBytes(pstep.Data, "id").String(),
			)
		case step.PlugStepAddInnerLink:
		case step.PlugStepAddRemoveLink:
		default:

			return easyerr.NotImpl()
		}

	}

	return nil
}

// apply

func (h *HubV2) applyNewPlug(tenantId, pid, itemid string, handle Handle, data xbprint.NewPlug) error {

	plug := &entities.Plug{
		Id:               pid,
		TenantId:         tenantId,
		Name:             data.Name,
		Live:             false,
		Dev:              true,
		ExtraMeta:        entities.JsonStrMap{},
		InvokePolicy:     "",
		BprintId:         handle.opts.BprintId,
		BprintItemId:     itemid,
		BprintInstanceId: handle.opts.InstanceId,
	}

	err := h.syncdb.PlugNew(tenantId, plug)
	if err != nil {
		return err
	}

	for _, agent := range data.Agents {
		err = h.applyAddAgent(tenantId, pid, agent)
		if err != nil {
			return err
		}
	}

	// set live to true ?

	return nil
}

func (h *HubV2) applyUpdateAgent(tenantId, pid string, data map[string]any) error {

	id := data["id"].(string)
	delete(data, "id")

	return h.syncdb.AgentUpdate(tenantId, pid, id, data)
}

func (h *HubV2) applyRemoveAgent(tenantId, pid, id string) error {
	return h.syncdb.AgentDel(tenantId, pid, id)
}

func (h *HubV2) applyAddAgent(tenantId, pid string, data xbprint.NewAgent) error {

	return h.syncdb.AgentNew(tenantId, &entities.Agent{
		Id:         data.Name,
		Name:       data.Name,
		Type:       data.Type,
		Executor:   data.Executor,
		IfaceFile:  data.IfaceFile,
		EntryFile:  data.EntryFile,
		WebEntry:   data.WebEntry,
		WebScript:  data.WebScript,
		WebStyle:   data.WebStyle,
		WebLoader:  data.WebLoader,
		WebFiles:   data.WebFiles,
		ModVersion: 0,
		PlugId:     pid,
		TenantId:   tenantId,
	})
}
