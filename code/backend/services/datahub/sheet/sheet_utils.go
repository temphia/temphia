package sheet

import (
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (s *Sheet) ExportSheets(txid uint32, opts dyndb.ExportOptions) (*dyndb.ExportData, error) {

	fresp := dyndb.ExportData{
		Source:    s.source,
		Group:     s.group,
		Date:      time.Now(),
		SheetData: make(map[int64]dyndb.SheetData),
	}

	sraw, err := s.ListSheet(txid)
	if err != nil {
		return nil, err
	}

	sheets := make([]dyndb.Sheet, 0, len(sraw))

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &sheets,
	})

	err = decoder.Decode(sraw)
	if err != nil {
		return nil, err
	}

	for _, sid := range opts.Sheets {

		sdata := dyndb.SheetData{
			Id:      sid,
			Name:    "",
			Columns: nil,
			Cells:   make(map[int64]map[int64]dyndb.SheetCell),
		}

		rowCursor := int64(0)

		for {
			resp, err := s.Query(txid, &dyndb.QuerySheetReq{
				TenantId:    opts.TenantId,
				Group:       opts.Group,
				SheetId:     sid,
				RowCursorId: rowCursor,
			})
			if err != nil {
				return nil, easyerr.Wrap("err querying data for export", err)
			}

			rowslen := len(resp.Cells)
			if rowslen == 0 {
				break
			}

			for _, v := range resp.Cells {

				if nval, ok := v["numval"]; ok {
					switch n := nval.(type) {
					case string:
						inum, err := strconv.ParseInt(n, 10, 64)
						if err != nil {
							panic(err)
						}
						v["numval"] = inum
					}
				}

				cell := dyndb.SheetCell{}
				err := mapstructure.Decode(v, &cell)
				if err != nil {
					return nil, err
				}

				rowid := cell.RowId
				colid := cell.ColId

				cell.RowId = 0
				cell.ColId = 0
				cell.SheetId = 0

				rowCells, ok := sdata.Cells[rowid]
				if !ok {
					rowCells = make(map[int64]dyndb.SheetCell)
					sdata.Cells[rowid] = rowCells
				}

				rowCells[colid] = cell
				if rowid > rowCursor {
					rowCursor = rowid
				}

			}
		}

		cols, err := s.ListSheetColumn(txid, sid)
		if err != nil {
			return nil, err
		}

		ocols := make([]dyndb.SheetColumn, 0, len(cols))

		err = mapstructure.Decode(cols, &ocols)
		if err != nil {
			return nil, err
		}
		sdata.Columns = ocols

		for _, psheet := range sheets {

			if psheet.Id == sid {
				sdata.Name = psheet.Name
				break
			}

		}

		fresp.SheetData[sid] = sdata

	}

	return &fresp, nil

}

func (s *Sheet) ImportSheets(txid uint32, opts dyndb.ImportOptions, data *dyndb.ExportData) error {

	/*



		allcols, err := s.tableHub.SimpleQuery(0, dyndb.SimpleQueryReq{
			TenantId:    s.tenantId,
			Group:       s.group,
			Table:       dyndb.SheetColumnTable,
			FilterConds: []dyndb.FilterCond{},
		})

		if err != nil {
			return err
		}

		cols := make(map[int64][]Column)

		for _, v := range allcols.Rows {
			col := Column{}
			err := mapstructure.Decode(v, &col)
			if err != nil {
				return easyerr.Wrap("err parsing column", err)
			}

			siblings, ok := cols[col.SheetId]
			if !ok {
				siblings = make([]Column, 0, 5)
			}

			cols[col.SheetId] = append(siblings, col)
		}

		// find exec_order for columns, if x sheet has reference to y
		// then y has to be seeded/imported first

		order := findExecOrder(cols)

		for _, sid := range order {

			frows := make(map[int64]map[string]any)

			for _, row := range data.Rows[sid] {
				// fixme => implement this
				rowid := row[dyndb.KeyPrimary]

				delete(row, dyndb.KeyPrimary)
				delete(row, dyndb.KeyVersion)

				pp.Println(rowid)

				s.NewRowWithCell(txid, sid, "", nil)

			}

			pp.Println("@sid", sid, frows)
		}
	*/

	return nil
}

// private

func findExecOrder(cols map[int64][]dyndb.SheetColumn) []int64 {
	visited := make(map[int64]bool)
	order := make([]int64, 0)

	var visit func(sheetId int64)
	visit = func(sheetId int64) {
		if visited[sheetId] {
			return
		}

		visited[sheetId] = true

		if colsOnSheet, ok := cols[sheetId]; ok {
			for _, col := range colsOnSheet {
				// If the column has a reference to another sheet, process that sheet first
				if col.Ctype == "reference" && col.RefSheet != 0 {
					visit(col.RefSheet)
				}
			}
		}

		order = append(order, sheetId)
	}

	for sheetId := range cols {
		if !visited[sheetId] {
			visit(sheetId)
		}
	}
	return order
}
