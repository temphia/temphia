package plug

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/xtypes"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/ztrue/tracerr"
)

type PlugInstancer struct {
	app    xtypes.App
	pacman repox.Hub
	syncer store.SyncDB
}

func New(app xtypes.App) xinstance.Instancer {

	deps := app.GetDeps()

	return &PlugInstancer{
		app:    app,
		pacman: deps.RepoHub().(repox.Hub),
		syncer: deps.CoreHub().(store.CoreHub),
	}
}

func (pi *PlugInstancer) Instance(opts xinstance.Options) (*xinstance.Response, error) {

	schemaData := &xbprint.NewPlug{}
	err := opts.Handle.LoadFile(opts.File, schemaData)
	if err != nil {
		return nil, err
	}

	uopts := PlugOptions{}
	if opts.UserData != nil {
		pp.Println("PARSING USER DATA", json.Unmarshal(opts.UserData, &uopts))
	}

	if uopts.Id == "" {
		uopts.Id = xid.New().String()
	}

	resp, err := pi.instance(uopts.Id, opts, schemaData)
	if err != nil {
		return nil, err
	}

	return &xinstance.Response{
		Ok:      true,
		Message: "",
		Slug:    uopts.Id,
		Data:    resp,
	}, nil

}

func (pi *PlugInstancer) instance(pid string, opts xinstance.Options, schema *xbprint.NewPlug) (*PlugResponse, error) {

	plug := &entities.Plug{
		Id:           pid,
		TenantId:     opts.TenantId,
		Name:         schema.Name,
		Live:         true,
		Dev:          true,
		ExtraMeta:    entities.JsonStrMap{},
		InvokePolicy: "",
		BprintId:     opts.BprintId,
	}

	err := pi.syncer.PlugNew(opts.TenantId, plug)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	resp := &PlugResponse{
		Agents:    make([]string, 0),
		Resources: make([]string, 0),
		ErrAgents: make(map[string]string),
	}

	for _, na := range schema.Agents {
		agent := &entities.Agent{
			Id:        na.Name,
			Name:      na.Name,
			Type:      na.Type,
			Executor:  na.Executor,
			IfaceFile: na.IfaceFile,
			WebEntry:  na.WebEntry,
			WebScript: na.WebScript,
			WebStyle:  na.WebStyle,
			WebLoader: na.WebLoader,
			WebFiles:  na.WebFiles,
			EntryFile: na.EntryFile,
			PlugId:    pid,
			ExtraMeta: entities.JsonStrMap{},
			TenantId:  opts.TenantId,
		}

		pp.Println("@agents", agent)

		err := pi.syncer.AgentNew(opts.TenantId, agent)
		if err != nil {
			resp.AddAgentErr(agent.Name, err)
			continue
		}

		for _, rdata := range na.Resources {
			pp.Println("@agent_resources", rdata.Name, rdata)

			if rdata.RefData == nil {
				pobj := opts.Handle.GetPrevObject(rdata.RefName)
				if pobj == nil {
					continue
				}

				pp.Println("@pobj", pobj)

				err := pi.syncer.AgentResourceNew(opts.TenantId, &entities.AgentResource{
					Slug:       rdata.Name,
					PlugId:     pid,
					AgentId:    agent.Id,
					ResourceId: pobj.Slug,
					Actions:    "",
					Policy:     "",
					TenantId:   opts.TenantId,
				})

				if err != nil {
					resp.AddResourceErr(agent.Id, rdata.Name, err)
					pp.Println("@error_creating_agent_resource", err)
					continue
				}
			}

		}

		resp.Agents = append(resp.Agents, agent.Id)
	}

	pp.Println("@resp |>", resp)

	return resp, nil
}
