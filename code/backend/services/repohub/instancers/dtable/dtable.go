package dtable

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/services/repohub/seeder"
	"github.com/temphia/temphia/code/backend/xtypes"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

var _ xinstance.Instancer = (*dtabeInstancer)(nil)

type dtabeInstancer struct {
	app     xtypes.App
	pacman  repox.Hub
	cabhub  store.CabinetHub
	coreHub store.CoreHub
	dynhub  dyndb.DataHub
}

func New(app xtypes.App) *dtabeInstancer {

	deps := app.GetDeps()

	return &dtabeInstancer{
		app:     app,
		pacman:  deps.RepoHub().(repox.Hub),
		cabhub:  deps.Cabinet().(store.CabinetHub),
		coreHub: deps.CoreHub().(store.CoreHub),
		dynhub:  deps.DataHub().(dyndb.DataHub),
	}
}

func (di *dtabeInstancer) Instance(opts xinstance.Options) (*xinstance.Response, error) {

	schemaData := &xbprint.NewTableGroup{}

	err := opts.Handle.LoadFile(opts.File, schemaData)
	if err != nil {
		return nil, err
	}

	err = Validate(schemaData)
	if err != nil {
		return nil, err
	}

	dopts, err := di.extractUserOptions(opts.TenantId, opts.Automatic, opts.UserData, schemaData)
	if err != nil {
		return nil, err
	}

	resp, err := di.instance(opts.TenantId, dopts, schemaData)
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

func (di *dtabeInstancer) instance(tenantId string, opts *DataGroupRequest, schema *xbprint.NewTableGroup) (*DataGroupResponse, error) {

	var dhub dyndb.DynSource

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
			table.SyncType = dyndb.DynSyncTypeEventAndData
			table.ActivityType = dyndb.DynActivityTypeStrict
			continue
		}
		if tableOpts.SyncType == "" {
			tableOpts.SyncType = dyndb.DynSyncTypeEventAndData
		}

		if tableOpts.ActivityType == "" {
			tableOpts.ActivityType = dyndb.DynActivityTypeStrict
		}

		table.SyncType = tableOpts.SyncType
		table.ActivityType = tableOpts.ActivityType
	}

	err := dhub.NewGroup(tenantId, schema)
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

			err = dhub.NewView(tenantId, &entities.DataView{
				Id:          0,
				Name:        view.Name,
				Count:       view.Count,
				FilterConds: &view.FilterConds,
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
	case dyndb.DynSeedTypeData:
		err := seeder.DataSeed()
		if err != nil {
			resp.SeedError = err.Error()
		}

	case dyndb.DynSeedTypeAutogen:
		err = seeder.GeneratedSeed(dyndb.DefaultSeedNo)
		if err != nil {
			resp.SeedError = err.Error()
		}
	}

	return resp, err

}
