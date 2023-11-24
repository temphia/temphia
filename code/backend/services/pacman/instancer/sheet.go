package instancers

import (
	"encoding/json"
	"fmt"

	_ "embed"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xinstancer"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

//go:embed _sheet_schema.json
var sheetSchema []byte

func (i *instancer) InstanceSheetDirect(opts xinstancer.SheetOptions) (*xinstancer.Response, error) {
	tenantId := opts.UserContext.TenantId
	gslug := gFunc()
	schemaData := opts.Template

	// fixme => inside same txn
	txnId := uint32(0)

	parsedSchema := &xpackage.NewTableGroup{}

	err := json.Unmarshal(sheetSchema, parsedSchema)
	if err != nil {
		return nil, err
	}

	parsedSchema.Slug = gslug
	parsedSchema.Name = opts.Template.Name
	parsedSchema.Description = opts.Template.Info

	err = i.datahub.GetDynDB().NewGroup(tenantId, parsedSchema)
	if err != nil {
		return nil, err
	}

	dtable := i.datahub.GetDataTableHub(tenantId, gslug)

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
				case xpackage.SheetColTypeNumber:
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

	return &xinstancer.Response{
		StepHead: "",
		Items:    map[string]string{},
	}, nil
}
