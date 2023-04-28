package sheet

import (
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (s *Sheet) ExportSheets(txid uint32, opts dyndb.ExportOptions) (*dyndb.ExportData, error) {

	/*

		fresp := dyndb.ExportData{
			Source:  s.source,
			Group:   s.group,
			Date:    time.Now(),
			Rows:    make(map[int64][]map[string]any),
			Columns: make(map[int64][]map[string]any),
		}

		for _, sid := range opts.Sheets {

			rows := make([]map[string]any, 0)

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

				rows = append(rows, resp.Cells...)

				switch rid := resp.Cells[rowslen-1][dyndb.KeyPrimary].(type) {
				case float64:
					rowCursor = int64(rid)
				case int64:
					rowCursor = rid
				default:
					panic("row id should be float64 or int64")
				}

			}
			fresp.Rows[sid] = rows

			cols, err := s.ListSheetColumn(txid, sid)
			if err != nil {
				return nil, err
			}
			fresp.Columns[sid] = cols
		}

		return &fresp, nil

	*/

	return nil, nil
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

func findExecOrder(cols map[int64][]dyndb.Column) []int64 {
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
