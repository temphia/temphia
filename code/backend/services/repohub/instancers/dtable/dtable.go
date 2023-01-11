package dtable

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/services/repohub/seeder"
	"github.com/temphia/temphia/code/backend/xtypes"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var _ xinstance.Instancer = (*dtabeInstancer)(nil)

type dtabeInstancer struct {
	app     xtypes.App
	pacman  repox.Hub
	cabhub  store.CabinetHub
	coreHub store.CoreHub
	dynhub  store.DataHub
}

func New(app xtypes.App) *dtabeInstancer {

	deps := app.GetDeps()

	return &dtabeInstancer{
		app:     app,
		pacman:  deps.RepoHub().(repox.Hub),
		cabhub:  deps.Cabinet().(store.CabinetHub),
		coreHub: deps.CoreHub().(store.CoreHub),
		dynhub:  deps.DataHub().(store.DataHub),
	}
}

func (di *dtabeInstancer) Instance(opts xinstance.Options) (*xinstance.Response, error) {

	schemaData := &xbprint.NewTableGroup{}

	err := opts.Handle.LoadFile(opts.File, schemaData)
	if err != nil {
		return nil, err
	}

	// fixme in automatic mode data could be nil

	dopts := &DataGroupRequest{}
	err = json.Unmarshal(opts.UserData, dopts)
	if err != nil {
		return nil, err
	}

	resp, err := di.instance(opts.TenantId, opts.File, dopts, schemaData)
	if err != nil {
		return nil, err
	}

	return &xinstance.Response{
		Ok:      true,
		Message: "",
		Slug:    resp.GroupSlug,
		Data:    resp,
	}, nil

}

func (di *dtabeInstancer) instance(tenantId, file string, opts *DataGroupRequest, schema *xbprint.NewTableGroup) (*DataGroupResponse, error) {

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

	pp.Println("@table    ===================|>", schema, opts)

	for _, table := range schema.Tables {
		tableOpts, ok := opts.TableOptions[table.Slug]
		if !ok {
			table.SyncType = store.DynSyncTypeEventAndData
			table.ActivityType = store.DynActivityTypeStrict
			continue
		}
		if tableOpts.SyncType == "" {
			tableOpts.SyncType = store.DynSyncTypeEventAndData
		}

		if tableOpts.ActivityType == "" {
			tableOpts.ActivityType = store.DynActivityTypeStrict
		}

		table.SyncType = tableOpts.SyncType
		table.ActivityType = tableOpts.ActivityType
	}

	err := dhub.NewGroup(schema)
	if err != nil {
		return nil, err
	}

	resp := &DataGroupResponse{
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

	seeder := seeder.New(schema, di.pacman, dhub, tenantId, opts.GroupSlug, opts.UserId)

	switch opts.SeedType {
	case store.DynSeedTypeData:
		err := seeder.DataSeed()
		if err != nil {
			resp.SeedError = err.Error()
		}

	case store.DynSeedTypeAutogen:
		err = seeder.GeneratedSeed(200)
		if err != nil {
			resp.SeedError = err.Error()
		}
	}

	return resp, err

}
