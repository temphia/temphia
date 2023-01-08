package plug

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints/instancer"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/ztrue/tracerr"
)

type PlugInstancer struct {
	app    xtypes.App
	pacman repox.Hub
	syncer store.SyncDB
}

func New(app xtypes.App) instancer.Instancer {

	deps := app.GetDeps()

	return &PlugInstancer{
		app:    app,
		pacman: deps.RepoHub().(repox.Hub),
		syncer: deps.CoreHub().(store.CoreHub),
	}
}

func (pi *PlugInstancer) Instance(opts instancer.Options) (any, error) {

	popts := instancer.Plug{}
	err := json.Unmarshal(opts.Data, &popts)
	if err != nil {
		return nil, err
	}

	schemaData := &bprints.PlugNew{}
	err = pi.pacman.ParseInstanceFile(opts.TenantId, opts.Bid, opts.File, schemaData)
	if err != nil {
		return nil, err
	}

	if popts.NewPlugId == "" {
		popts.NewPlugId = xid.New().String()
	}

	return pi.instance(opts.TenantId, opts.Bid, popts, schemaData)
}

func (pi *PlugInstancer) instance(tenantId, bid string, opts instancer.Plug, schema *bprints.PlugNew) (any, error) {

	plug := &entities.Plug{
		Id:           opts.NewPlugId,
		TenantId:     tenantId,
		Name:         opts.NewPlugName,
		Live:         true,
		Dev:          true,
		ExtraMeta:    nil,
		Owner:        "",
		InvokePolicy: "",
		BprintId:     bid,
	}

	err := pi.syncer.PlugNew(tenantId, plug)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	resp := &PlugResponse{
		Agents:       make([]string, 0),
		Resources:    make([]string, 0),
		ErrAgents:    make(map[string]string),
		ErrResources: make(map[string]string),
	}

	for aname, aopts := range opts.AgentOptions {

		ahint, ok := schema.AgentHints[aname]
		if !ok {
			return nil, easyerr.NotFound()
		}

		agent := &entities.Agent{
			Id:        aname,
			Name:      aopts.Name,
			Type:      ahint.Type,
			Executor:  ahint.Executor,
			IfaceFile: ahint.IfaceFile,
			WebEntry:  ahint.WebEntry,
			WebScript: ahint.WebScript,
			WebStyle:  ahint.WebStyle,
			WebLoader: ahint.WebLoader,
			WebFiles:  ahint.WebFiles,
			EnvVars:   entities.JsonStrMap{},
			PlugId:    opts.NewPlugId,
			ExtraMeta: entities.JsonStrMap{},
			TenantId:  tenantId,
		}

		err := pi.syncer.AgentNew(tenantId, agent)
		if err != nil {
			pp.Println(err)
			resp.ErrAgents[aname] = err.Error()
			continue
		}

		resp.Agents = append(resp.Agents, aname)
	}

	for _, res := range opts.Resources {

		if res.Id == "" {
			res.Id = xid.New().String()
		}

		resource := &entities.Resource{
			Name:      res.Name,
			TenantId:  tenantId,
			Type:      res.Type,
			SubType:   res.SubType,
			Payload:   res.Payload,
			Policy:    res.Policy,
			PlugId:    opts.NewPlugId,
			Id:        res.Id,
			ExtraMeta: nil,
		}

		err := pi.syncer.ResourceNew(tenantId, resource)
		if err != nil {
			resp.ErrResources[res.Name] = err.Error()
			continue
		}

		resp.Resources = append(resp.Resources, res.Id)
	}

	for aname, aopts := range opts.AgentOptions {
		for resKey, resId := range aopts.Resources {

			err = pi.syncer.AgentResourceNew(tenantId, &entities.AgentResource{
				Id:         0,
				Name:       resKey,
				PlugId:     opts.NewPlugId,
				AgentId:    aname,
				ResourceId: resId,
				Actions:    "",
				Policy:     "",
				TenantId:   tenantId,
			})

			pp.Println(err)
		}
	}

	return resp, nil
}
