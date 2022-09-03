package dtable

import (
	"encoding/json"

	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints/instancer"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/instance"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type dtabeInstancer struct {
	app     xtypes.App
	pacman  service.Pacman
	cabhub  store.CabinetHub
	coreHub store.CoreHub
	dynhub  store.DynHub
}

func New(app xtypes.App) *dtabeInstancer {

	deps := app.GetDeps()

	return &dtabeInstancer{
		app:     app,
		pacman:  deps.Pacman().(service.Pacman),
		cabhub:  deps.Cabinet().(store.CabinetHub),
		coreHub: deps.CoreHub().(store.CoreHub),
		dynhub:  deps.DynHub().(store.DynHub),
	}
}

func (di *dtabeInstancer) Instance(opts instancer.Options) (interface{}, error) {

	schemaData := &bprints.NewTableGroup{}
	err := di.pacman.ParseInstanceFile(opts.TenantId, opts.Bid, opts.File, schemaData)
	if err != nil {
		return nil, err
	}

	dopts := &instance.DataGroupRequest{}
	err = json.Unmarshal(opts.Data, dopts)
	if err != nil {
		return nil, err
	}

	return di.instance(opts.TenantId, opts.File, dopts, schemaData)
}

func (di *dtabeInstancer) instance(tenantId, file string, opts *instance.DataGroupRequest, schema *bprints.NewTableGroup) (*instance.DataGroupResponse, error) {

	var dhub store.DynSource

	if opts.DyndbSource == "" {
		dhub = di.dynhub.DefaultSource(tenantId)
	} else {
		dhub = di.dynhub.GetSource(opts.DyndbSource, tenantId)
	}

	if opts.GroupName != "" {
		schema.Name = opts.GroupName
	}

	if opts.GroupSlug != "" {
		schema.Slug = opts.GroupSlug
	}

	if opts.CabinetFolder != "" {
		schema.CabinetFolder = opts.CabinetFolder
	}

	if opts.CabinetSource != "" {
		schema.CabinetSource = opts.CabinetSource
	}

	for _, table := range schema.Tables {
		tableOpts, ok := opts.TableOptions[table.Slug]
		if !ok {
			continue
		}
		table.SyncType = tableOpts.SyncType
		table.ActivityType = tableOpts.ActivityType
	}

	err := dhub.NewGroup(schema)
	if err != nil {
		return nil, err
	}

	resp := &instance.DataGroupResponse{
		Source:     dhub.Name(),
		GroupSlug:  opts.GroupSlug,
		GroupName:  opts.GroupName,
		ViewErrors: make(map[string]string),
	}

	for _, tbl := range schema.Tables {
		for _, view := range tbl.Views {

			err = dhub.NewView(&entities.DataView{
				Id:          0,
				Name:        view.Name,
				Count:       view.Count,
				FilterConds: view.FilterConds,
				Selects:     view.Selects,
				MainColumn:  view.MainColumn,
				SearchTerm:  view.SearchTerm,
				TableID:     tbl.Slug,
				GroupID:     schema.Slug,
				TenantID:    tenantId,
				ExtraMeta:   nil,
			})
			if err != nil {
				resp.ViewErrors[tbl.Name+"/"+view.Name] = err.Error()
			}
		}
	}

	seeder := Seeder{
		tg:     schema,
		model:  nil,
		pacman: di.pacman,
		source: dhub,
		tenant: tenantId,
		group:  opts.GroupSlug,
		userId: opts.UserId,
	}

	switch opts.SeedType {
	case store.DynSeedTypeData:
		err := seeder.dataSeed()
		if err != nil {
			resp.SeedError = err.Error()
		}

	case store.DynSeedTypeAutogen:
		err = seeder.generatedSeed(200)
		if err != nil {
			resp.SeedError = err.Error()
		}
	}

	return resp, err

}
