package sheet

import (
	_ "embed"
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/services/repohub/instancers/dtable"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
	"github.com/temphia/temphia/code/backend/xtypes/store"
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
	pacman  repox.Hub
	cabhub  store.CabinetHub
	coreHub store.CoreHub
	dynhub  store.DataHub

	dataInstancer DataInstancer
}

func New(app xtypes.App) *SheetInstancer {

	deps := app.GetDeps()

	return &SheetInstancer{
		app:           app,
		pacman:        deps.RepoHub().(repox.Hub),
		cabhub:        deps.Cabinet().(store.CabinetHub),
		coreHub:       deps.CoreHub().(store.CoreHub),
		dynhub:        deps.DataHub().(store.DataHub),
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

	dropts.SeedType = ""
	dropts.GroupName = schemaData.Name
	resp, err := s.dataInstancer.DirectInstance(opts.TenantId, dropts, &parsedSchema)
	if err != nil {
		return nil, err
	}

	source := s.dynhub.GetSource(resp.Source, opts.TenantId)

	// fixme => inside same txn

	txnId := uint32(0)

	sheetsIdx := make(map[string]int64)
	colsIdx := make(map[string]map[string]int64)

	for sidx := range schemaData.Sheets {
		sheet := &schemaData.Sheets[sidx]
		idx, err := source.NewRow(uint32(txnId), store.NewRowReq{
			TenantId: opts.TenantId,
			Group:    resp.GroupSlug,
			Table:    "sheets",
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

			cid, err := source.NewRow(txnId, store.NewRowReq{
				TenantId: opts.TenantId,
				Group:    resp.GroupSlug,
				Table:    "scols",
				Data: map[string]any{
					"name":    column.Name,
					"ctype":   column.Ctype,
					"sheetid": idx,
					"color":   column.Color,
					"opts":    extraopts,
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

			rowid, err := source.NewRow(uint32(txnId), store.NewRowReq{
				TenantId: opts.TenantId,
				Group:    resp.GroupSlug,
				Table:    "srows",
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
					cellData["num_value"] = seedCellData
				default:
					cellData["value"] = seedCellData
				}

				cellid, err := source.NewRow(uint32(txnId), store.NewRowReq{
					TenantId: opts.TenantId,
					Group:    resp.GroupSlug,
					Table:    "scells",
					Data:     cellData,
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

	return nil, nil
}
