package instancers

import (
	"fmt"
	"strconv"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xinstancer"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"gopkg.in/yaml.v2"
)

var _ xinstancer.Instancer = (*instancer)(nil)

type instancer struct {
	corehub store.CoreHub
	cabinet store.CabinetHub
	datahub dyndb.DataHub
}

func New(corehub store.CoreHub, cabinet store.CabinetHub, datahub dyndb.DataHub) *instancer {
	return &instancer{
		corehub: corehub,
		cabinet: cabinet,
		datahub: datahub,
	}
}

var (
	gFunc = xtypes.GetSlugGenerator(5)
)

func (i *instancer) Instance(opts xinstancer.Options) (*xinstancer.Response, error) {
	as, err := i.loadAppSchema(opts.TenantId, opts.BprintId)
	if err != nil {
		return nil, err
	}

	if opts.PlugId == "" {
		opts.PlugId = gFunc()
	}

	err = i.corehub.PlugNew(opts.TenantId, &entities.Plug{
		Id:               opts.PlugId,
		Name:             as.Name,
		BprintId:         opts.BprintId,
		InstancedObjects: entities.JsonStrMap{},
		BprintItemId:     "",
		StepHead:         "",
		TenantId:         opts.TenantId,
	})
	if err != nil {
		return nil, err
	}

	resp, err := i.runStep(as, opts)
	if err != nil {
		return nil, err
	}

	// update plug here

	err = i.corehub.PlugUpdate(opts.TenantId, opts.PlugId, map[string]any{
		"instanced_objects": opts.InstancedIds,
		"step_head":         resp.Items,
	})
	if err != nil {

		return nil, err
	}

	return resp, nil
}

func (i *instancer) runStep(as *xpackage.AppSchema, opts xinstancer.Options) (*xinstancer.Response, error) {

	for _, step := range as.Steps {

		getIntId := func() int64 {
			id, _ := strconv.ParseInt(opts.InstancedIds[step.ObjectId], 10, 64)
			return id
		}

		switch step.Type {
		case xinstancer.PlugStepNewAgent:
			agent := &entities.Agent{}
			err := step.DataAs(agent)
			if err != nil {
				return nil, err
			}

			err = i.corehub.AgentNew(opts.TenantId, agent)
			if err != nil {
				return nil, err
			}

			opts.InstancedIds[step.ObjectId] = agent.Id

		case xinstancer.PlugStepUpdateAgent:
			id := opts.InstancedIds[step.ObjectId]
			data := make(map[string]any)
			err := step.DataAs(data)
			if err != nil {
				return nil, err
			}

			err = i.corehub.AgentUpdate(opts.TenantId, opts.PlugId, id, data)
			if err != nil {
				return nil, err
			}
		case xinstancer.PlugStepRemoveAgent:
			id := opts.InstancedIds[step.ObjectId]
			err := i.corehub.AgentDel(opts.TenantId, opts.PlugId, id)
			if err != nil {
				return nil, err
			}
		case xinstancer.PlugStepNewResourceModule:
			res := &entities.Resource{}
			err := step.DataAs(res)
			if err != nil {
				return nil, err
			}

			res.Id = gFunc()
			res.OwnedByPlug = opts.PlugId
			// fixme => logic bashed on different resource types ?

			err = i.corehub.ResourceNew(opts.TenantId, res)
			if err != nil {
				return nil, err
			}

			opts.InstancedIds[step.ObjectId] = res.Id

		case xinstancer.PlugStepUpdateResourceModule:
			id := opts.InstancedIds[step.ObjectId]
			data := make(map[string]any)
			err := step.DataAs(data)
			if err != nil {
				return nil, err
			}

			err = i.corehub.ResourceUpdate(opts.TenantId, id, data)
			if err != nil {
				return nil, err
			}

		case xinstancer.PlugStepRemoveResourceModule:
			id := opts.InstancedIds[step.ObjectId]
			err := i.corehub.ResourceDel(opts.TenantId, id)
			if err != nil {
				return nil, err
			}

		case xinstancer.PlugStepAddTargetApp:
			target := &entities.TargetApp{}
			err := step.DataAs(target)
			if err != nil {
				return nil, err
			}

			target.AgentId = opts.InstancedIds[target.AgentId]
			target.OwnedByPlug = opts.PlugId
			target.TenantId = opts.TenantId

			// fixme => logic bashed on different targetapp types ?

			tid, err := i.corehub.AddTargetApp(target)
			if err != nil {
				return nil, err
			}

			opts.InstancedIds[step.ObjectId] = fmt.Sprint(tid)

		case xinstancer.PlugStepUpdateTargetApp:

			id := getIntId()

			data := make(map[string]any)
			err := step.DataAs(data)
			if err != nil {
				return nil, err
			}

			ts, err := i.corehub.ListTargetApp(opts.TenantId, map[string]any{
				"id": id,
			})
			if err != nil {
				return nil, err
			}
			if len(ts) == 0 {
				return nil, easyerr.Error("targetapp not found")
			}

			ttype := ts[0].TargetType
			err = i.corehub.UpdateTargetApp(opts.TenantId, ttype, id, data)
			if err != nil {
				return nil, err
			}

		case xinstancer.PlugStepDeleteTargetApp:
			id := getIntId()
			ts, err := i.corehub.ListTargetApp(opts.TenantId, map[string]any{
				"id": id,
			})
			if err != nil {
				return nil, err
			}
			if len(ts) == 0 {
				return nil, easyerr.Error("targetapp not found")
			}
			ttype := ts[0].TargetType
			err = i.corehub.RemoveTargetApp(opts.TenantId, ttype, id)
			if err != nil {
				return nil, err
			}
		case xinstancer.PlugStepAddTargetHook:
			target := &entities.TargetHook{}
			err := step.DataAs(target)
			if err != nil {
				return nil, err
			}

			target.AgentId = opts.InstancedIds[target.AgentId]
			target.OwnedByPlug = opts.PlugId
			target.TenantId = opts.TenantId

			// fixme => logic bashed on different targetapp types ?

			tid, err := i.corehub.AddTargetHook(target)
			if err != nil {
				return nil, err
			}

			opts.InstancedIds[step.ObjectId] = fmt.Sprint(tid)

		case xinstancer.PlugStepUpdateTargetHook:
			id := getIntId()

			data := make(map[string]any)
			err := step.DataAs(data)
			if err != nil {
				return nil, err
			}

			ts, err := i.corehub.ListTargetHook(opts.TenantId, map[string]any{
				"id": id,
			})
			if err != nil {
				return nil, err
			}
			if len(ts) == 0 {
				return nil, easyerr.Error("targethook not found")
			}

			ttype := ts[0].TargetType
			err = i.corehub.UpdateTargetHook(opts.TenantId, ttype, id, data)
			if err != nil {
				return nil, err
			}

		case xinstancer.PlugStepDeleteTargetHook:
			id := getIntId()
			ts, err := i.corehub.ListTargetApp(opts.TenantId, map[string]any{
				"id": id,
			})
			if err != nil {
				return nil, err
			}
			if len(ts) == 0 {
				return nil, easyerr.Error("targethook not found")
			}
			ttype := ts[0].TargetType
			err = i.corehub.RemoveTargetHook(opts.TenantId, ttype, id)
			if err != nil {
				return nil, err
			}

		case xinstancer.PlugStepRunDataMigration:
			schema, err := i.readMigration(opts.TenantId, opts.BprintId, step.File)
			if err != nil {
				return nil, err
			}

			dyndb := i.datahub.GetDynDB()

			dyndb.MigrateSchema(opts.TenantId, schema)

			pp.Println("@run_migration_here")

		default:

			panic("not implemented")
		}

	}

	return nil, nil
}

func (i *instancer) Upgrade(opts xinstancer.Options) error {

	as, err := i.loadAppSchema(opts.TenantId, opts.NextBprintId)
	if err != nil {
		return err
	}

	if opts.PlugId == "" {
		opts.PlugId = gFunc()
	}

	plug, err := i.corehub.PlugGet(opts.TenantId, opts.PlugId)
	if err != nil {
		return err
	}

	opts.InstancedIds = plug.InstancedObjects

	// fixme => pass last step_head

	i.runStep(as, opts)

	return nil

}

func (i *instancer) InstanceSheetDirect(opts xinstancer.SheetOptions) (*xinstancer.Response, error) {

	return nil, nil
}

// private

func (i *instancer) readMigration(tenantId, bprintid, file string) (xpackage.MigrateOptions, error) {

	return xpackage.MigrateOptions{}, nil
}

func (i *instancer) loadAppSchema(tenantId, bprintid string) (*xpackage.AppSchema, error) {

	as := &xpackage.AppSchema{}
	out, err := i.loadBprintFile("", "app.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(out, as)
	if err != nil {
		return nil, err
	}

	return as, nil
}

func (i *instancer) loadBprintFile(folder, file string) ([]byte, error) {

	return nil, nil
}
