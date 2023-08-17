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
	tenantId := handle.tenantId

	pid := item.Name + handle.opts.InstanceId

	schema, err := h.readSchema(tenantId, handle.opts.BprintId, item.File)
	if err != nil {
		return err
	}

	return h.processStep(handle, true, pid, "", item.Name, schema.Steps)
}

func (h *HubV2) processStep(handle Handle, new bool, pid, laststep, name string, steps []step.Step) error {
	tenantId := handle.tenantId

	for _, pstep := range steps {

		switch pstep.Name {
		case step.PlugStepNewPlug:
			pschema := xbprint.NewPlug{}
			err := json.Unmarshal(pstep.Data, &pschema)
			if err != nil {
				return err
			}

			err = h.applyNewPlug(tenantId, pid, name, handle, pschema)
			if err != nil {
				return err
			}

		case step.PlugStepNewAgent:

			data := xbprint.NewAgent{}
			err := json.Unmarshal(pstep.Data, &data)
			if err != nil {
				return err
			}

			err = h.applyAddAgent(tenantId, pid, data)
			if err != nil {
				return err
			}

		case step.PlugStepUpdateAgent:
			data := make(map[string]any)
			err := json.Unmarshal(pstep.Data, &data)
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
			data := xbprint.NewInnerLink{}
			err := json.Unmarshal(pstep.Data, &data)
			if err != nil {
				return err
			}
			return h.applyPlugStepAddInnerLink(tenantId, pid, data)

		case step.PlugStepAddRemoveLink:

			agentId := gjson.GetBytes(pstep.Data, "agent_id").String()
			name := gjson.GetBytes(pstep.Data, "name").String()

			return h.applyPlugStepRemoveInnerLink(
				tenantId,
				pid,
				agentId,
				name,
			)

		default:

			return easyerr.NotImpl()
		}

	}

	return nil

}

// apply

func (h *HubV2) applyNewPlug(tenantId, pid, itemid string, handle Handle, data xbprint.NewPlug) error {

	plug := &entities.Plug{
		Id:           pid,
		TenantId:     tenantId,
		Name:         data.Name,
		Live:         false,
		Dev:          true,
		ExtraMeta:    entities.JsonStrMap{},
		BprintId:     handle.opts.BprintId,
		BprintItemId: itemid,
		//		BprintInstanceId: handle.opts.InstanceId,
	}

	err := h.corehub.PlugNew(tenantId, plug)
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

	return h.corehub.AgentUpdate(tenantId, pid, id, data)
}

func (h *HubV2) applyRemoveAgent(tenantId, pid, id string) error {
	return h.corehub.AgentDel(tenantId, pid, id)
}

func (h *HubV2) applyAddAgent(tenantId, pid string, data xbprint.NewAgent) error {

	return h.corehub.AgentNew(tenantId, &entities.Agent{
		Id:         data.Name,
		Name:       data.Name,
		Type:       data.Type,
		Executor:   data.Executor,
		IfaceFile:  data.IfaceFile,
		EntryFile:  data.EntryFile,
		WebOptions: data.WebFiles,
		WebFiles:   data.WebFiles,
		ModVersion: 0,
		PlugId:     pid,
		TenantId:   tenantId,
	})
}

func (h *HubV2) applyPlugStepAddInnerLink(tenantId, pid string, data xbprint.NewInnerLink) error {

	return h.corehub.AgentLinkNew(tenantId, &entities.AgentLink{
		Name:      data.Slug,
		FromPlug:  pid,
		FromAgent: data.From,
		ToPlug:    pid,
		ToAgent:   data.To,
		TenantId:  tenantId,
	})
}

func (h *HubV2) applyPlugStepRemoveInnerLink(tenantId, pid, agentId, name string) error {

	links, err := h.corehub.AgentLinkList(tenantId, pid, agentId)
	if err != nil {
		return err
	}

	for _, al := range links {
		if al.Name == agentId {
			return h.corehub.AgentLinkDel(tenantId, pid, agentId, al.Id)
		}
	}

	return nil
}
