package datahub

import "github.com/temphia/temphia/code/backend/xtypes/store"

func (d *dynSource) ListSheetGroup(opts store.ListSheetGroupReq) (*store.ListSheetGroupResp, error) {

	ddb := d.dynDB()

	sheetRows, err := ddb.SimpleQuery(0, store.SimpleQueryReq{
		TenantId: opts.TenantId,
		Group:    opts.Group,
		Table:    "sheets",
	})
	if err != nil {
		return nil, err
	}

	return &store.ListSheetGroupResp{
		Sheets: sheetRows.Rows,
	}, nil
}

func (d *dynSource) LoadSheet(opts store.LoadSheetReq) (*store.LoadSheetResp, error) {

	ddb := d.dynDB()

	columns, err := ddb.SimpleQuery(0, store.SimpleQueryReq{
		TenantId: opts.TenantId,
		Group:    opts.Group,
		Table:    "scols",
		FilterConds: []*store.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  opts.SheetId,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	cells, err := ddb.SimpleQuery(0, store.SimpleQueryReq{
		TenantId: opts.TenantId,
		Group:    opts.Group,
		Table:    "scells",
		FilterConds: []*store.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  opts.SheetId,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return &store.LoadSheetResp{
		Columns: columns.Rows,
		Cells:   cells.Rows,
	}, nil
}
