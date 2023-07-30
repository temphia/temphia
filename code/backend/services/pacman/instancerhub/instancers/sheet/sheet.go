package sheet

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/services/pacman/instancerhub/instancers/dtable"
	"github.com/temphia/temphia/code/backend/xtypes"

	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

var _ xinstance.Instancer = (*SheetInstancer)(nil)

//go:embed _sheet_schema.json
var sheetSchema []byte

var parsedSchema xbprint.NewTableGroup

func init() {

	err := json.Unmarshal(sheetSchema, &parsedSchema)
	if err != nil {
		panic("cannot Unmarshal _sheet_schema.json" + err.Error())
	}

}

type SheetInstancer struct {
	app     xtypes.App
	pacman  repox.Pacman
	cabhub  store.CabinetHub
	coreHub store.CoreHub
	dynhub  dyndb.DataHub

	dataInstancer DataInstancer
}

func New(app xtypes.App) *SheetInstancer {

	deps := app.GetDeps()

	return &SheetInstancer{
		app:           app,
		pacman:        deps.RepoHub().(repox.Pacman),
		cabhub:        deps.Cabinet().(store.CabinetHub),
		coreHub:       deps.CoreHub().(store.CoreHub),
		dynhub:        deps.DataHub().(dyndb.DataHub),
		dataInstancer: dtable.New(app),
	}

}

func (s *SheetInstancer) Instance(opts xinstance.Options) (*xinstance.Response, error) {
	schemaData := &xbprint.NewSheetGroup{}

	err := opts.Handle.LoadFile(opts.File, schemaData)
	if err != nil {
		return nil, err
	}

	extractor := dtable.ExtractUserOptions(s.cabhub, s.coreHub, s.dynhub)

	dropts, err := extractor(opts.TenantId, true, nil, &parsedSchema)
	if err != nil {
		return nil, err
	}

	return s.instanceInner(opts.TenantId, schemaData, dropts)
}

func (s *SheetInstancer) instanceInner(tenantId string, schemaData *xbprint.NewSheetGroup, dropts *dtable.DataGroupRequest) (*xinstance.Response, error) {

	dropts.SeedType = ""
	dropts.GroupName = schemaData.Name

	resp, err := s.dataInstancer.DirectInstance(tenantId, dropts, &xbprint.NewTableGroup{
		Name:          schemaData.Name,
		Slug:          parsedSchema.Slug,
		Description:   schemaData.Info,
		Tables:        parsedSchema.Tables,
		ExecOrder:     parsedSchema.ExecOrder,
		Renderer:      parsedSchema.Renderer,
		CabinetSource: parsedSchema.CabinetSource,
		CabinetFolder: parsedSchema.CabinetFolder,
	})

	if err != nil {
		return nil, err
	}

	return s.instance(resp.Source, tenantId, resp.GroupSlug, schemaData)
}

func (s *SheetInstancer) instance(source, tenantId, gslug string, schemaData *xbprint.NewSheetGroup) (*xinstance.Response, error) {

	dtable := s.dynhub.GetDataTableHub(tenantId, gslug)

	// fixme => inside same txn

	txnId := uint32(0)

	sheetsIdx := make(map[string]int64)
	colsIdx := make(map[string]map[string]int64)

	for sidx := range schemaData.Sheets {
		sheet := &schemaData.Sheets[sidx]
		idx, err := dtable.NewRow(uint32(txnId), dyndb.NewRowReq{
			TenantId: tenantId,
			Group:    gslug,
			Table:    dyndb.SheetTable,
			Data: map[string]any{
				"name": sheet.Name,
			},
		})
		if err != nil {
			return nil, err
		}
		sheetsIdx[sheet.Name] = idx
		currColsIndx := make(map[string]int64, len(sheet.Columns))
		colsIdx[sheet.Name] = currColsIndx

		for cidx := range sheet.Columns {
			column := &sheet.Columns[cidx]
			// fixme => reference ctype reverse lookup here

			extraopts := "{}"
			if column.ExtaOptions != nil {
				out, err := json.Marshal(column.ExtaOptions)
				if err == nil {
					extraopts = string(out)
				}

			}

			cid, err := dtable.NewRow(txnId, dyndb.NewRowReq{
				TenantId: tenantId,
				Group:    gslug,
				Table:    dyndb.SheetColumnTable,
				Data: map[string]any{
					"name":      column.Name,
					"ctype":     column.Ctype,
					"sheetid":   idx,
					"color":     column.Color,
					"extraopts": extraopts,
					"opts":      column.Options,
				},
			})
			if err != nil {
				return nil, err
			}

			currColsIndx[column.Name] = cid

		}

	}

	// seed data

	for sidx := range schemaData.Sheets {
		sheet := &schemaData.Sheets[sidx]

		currColsIndx := colsIdx[sheet.Name]

		for _, row := range sheet.SeedData {

			rowid, err := dtable.NewRow(uint32(txnId), dyndb.NewRowReq{
				TenantId: tenantId,
				Group:    gslug,
				Table:    dyndb.SheetRowTable,
				Data: map[string]any{
					"sheetid": sheetsIdx[sheet.Name],
				},
			})

			if err != nil {
				return nil, err
			}

			for cidx := range sheet.Columns {
				column := &sheet.Columns[cidx]

				seedCellData, ok := row[column.Name]
				if !ok {
					continue
				}

				cellData := map[string]any{
					"sheetid": sheetsIdx[sheet.Name],
					"colid":   currColsIndx[column.Name],
					"rowid":   rowid,
				}

				switch column.Ctype {
				case xbprint.SheetColTypeNumber:
					cellData["numval"] = seedCellData
				default:
					cellData["value"] = seedCellData
				}

				cellid, err := dtable.NewRow(uint32(txnId), dyndb.NewRowReq{
					TenantId: tenantId,
					Group:    gslug,
					Table:    dyndb.SheetCellTable,
					Data:     cellData,
					ModCtx: dyndb.ModCtx{
						InitSign: fmt.Sprint(rowid),
					},
				})
				if err != nil {
					pp.Println("err creating cell", err)
					continue
				}

				pp.Println("created cell %d", cellid)
			}

		}

	}

	// resp.Source

	return &xinstance.Response{
		Ok:             true,
		Type:           xbprint.TypeDataSheet,
		Slug:           gslug,
		ResourceTarget: fmt.Sprintf("%s/%s", source, gslug),
	}, nil

}

func (s *SheetInstancer) DirectInstance(tenantId, source, gslug string, template *xbprint.NewSheetGroup) (*xinstance.Response, error) {

	resp, err := dtable.ExtractUserOptions(s.cabhub, s.coreHub, s.dynhub)(tenantId, true, nil, &parsedSchema)
	if err != nil {
		return nil, err
	}

	return s.instanceInner(tenantId, template, resp)

}
