package data

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func (c *Controller) ListSheetGroup(uclaim *claim.Data) (*ListSheetGroupResp, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	sheetRows, err := dynDb.SimpleQuery(0, store.SimpleQueryReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    store.SheetTable,
	})
	if err != nil {
		return nil, err
	}

	return &ListSheetGroupResp{
		Sheets: sheetRows.Rows,
	}, nil

}

func (c *Controller) LoadSheet(uclaim *claim.Data, data *LoadSheetReq) (*LoadSheetResp, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	columns, err := dynDb.SimpleQuery(0, store.SimpleQueryReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    store.SheetColumnTable,
		FilterConds: []*store.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  data.SheetId,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	cells, err := dynDb.SimpleQuery(0, store.SimpleQueryReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    store.SheetCellTable,
		FilterConds: []*store.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  data.SheetId,
			},
		},

		OrderBy: "rowid",
	})

	if err != nil {
		return nil, err
	}

	return &LoadSheetResp{
		Columns: columns.Rows,
		Cells:   cells.Rows,
	}, nil

}

// sheets

func (c *Controller) ListSheet(uclaim *claim.Data) ([]map[string]any, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	resp, err := dynDb.SimpleQuery(0, store.SimpleQueryReq{
		TenantId: uclaim.TenantId,
		Table:    store.SheetTable,
		Group:    group,
	})
	if err != nil {
		return nil, err
	}

	return resp.Rows, nil
}

func (c *Controller) NewSheet(uclaim *claim.Data, data map[string]any) error {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	_, err := dynDb.NewRow(0, store.NewRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    store.SheetTable,
		Data:     data,
		ModCtx: store.ModCtx{
			UserId: uclaim.UserID,
		},
	})

	return err
}

func (c *Controller) GetSheet(uclaim *claim.Data, id int64) (map[string]any, error) {
	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDb.GetRow(0, store.GetRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    store.SheetTable,
		Id:       id,
	})

}

func (c *Controller) UpdateSheet(uclaim *claim.Data, id int64, data map[string]any) error {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	_, err := dynDb.UpdateRow(0, store.UpdateRowReq{
		TenantId: uclaim.TenantId,
		Id:       id,
		Group:    group,
		Table:    store.SheetTable,
		Data:     data,
	})

	return err

}

func (c *Controller) DeleteSheet(uclaim *claim.Data, id int64) error {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	dynDb.DeleteRows(0, store.DeleteRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    store.SheetTable,
		Id:       []int64{id},
		ModCtx: store.ModCtx{
			UserId: uclaim.UserID,
		},
	})

	return nil
}

// columns

func (c *Controller) ListSheetColumn(uclaim *claim.Data, sid int64) ([]map[string]any, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	resp, err := dynDb.SimpleQuery(0, store.SimpleQueryReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    store.SheetColumnTable,
		FilterConds: []*store.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  sid,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return resp.Rows, nil
}

func (c *Controller) NewSheetColumn(uclaim *claim.Data, sid int64, data map[string]any) (int64, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDb.NewRow(0, store.NewRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    store.SheetColumnTable,
		Data:     data,
		ModCtx: store.ModCtx{
			UserId: uclaim.UserID,
		},
	})

}

func (c *Controller) GetSheetColumn(uclaim *claim.Data, sid, cid int64) (map[string]any, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	dynDb.GetRow(0, store.GetRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    store.SheetColumnTable,
		Id:       cid,
	})

	return nil, nil
}

func (c *Controller) UpdateSheetColumn(uclaim *claim.Data, sid, cid int64, data map[string]any) error {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	dynDb.UpdateRow(0, store.UpdateRowReq{
		TenantId: uclaim.TenantId,
		Id:       cid,
		Group:    group,
		Table:    store.SheetColumnTable,
		Data:     data,
		ModCtx: store.ModCtx{
			UserId: uclaim.UserID,
		},
	})

	return nil

}

func (c *Controller) DeleteSheetColumn(uclaim *claim.Data, sid, cid int64) error {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	dynDb.DeleteRows(0, store.DeleteRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    store.SheetColumnTable,
		Id:       []int64{cid},
		ModCtx: store.ModCtx{
			UserId: uclaim.UserID,
		},
	})

	return nil
}

// row_cells

func (c *Controller) NewRowWithCell(uclaim *claim.Data, sid int64, data []map[string]any) (map[int64]map[string]any, error) {

	txid := uint32(0)

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	rid, err := dynDb.NewRow(txid, store.NewRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    store.SheetRowTable,
		Data:     map[string]any{},
		ModCtx: store.ModCtx{
			UserId: uclaim.UserID,
		},
	})
	if err != nil {
		return nil, err
	}

	// fixme => batch support

	for _, cellData := range data {
		cellData["rowid"] = rid

		cellid, err := dynDb.NewRow(txid, store.NewRowReq{
			TenantId: uclaim.TenantId,
			Group:    group,
			Table:    store.SheetCellTable,
			Data:     cellData,
			ModCtx: store.ModCtx{
				UserId: uclaim.UserID,
			},
		})
		if err != nil {
			return nil, err
		}
		pp.Println(cellid)

	}

	return nil, nil

}

func (c *Controller) UpdateRowWithCell(uclaim *claim.Data, sid, rid int64, data map[int64]map[string]any) (map[int64]map[string]any, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	for cellId, cellData := range data {

		_, err := dynDb.UpdateRow(0, store.UpdateRowReq{
			TenantId: uclaim.TenantId,
			Id:       cellId,
			Group:    group,
			Table:    store.SheetRowTable,
			Data:     cellData,
			ModCtx: store.ModCtx{
				UserId: uclaim.UserID,
			},
		})
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// models

type ListSheetGroupReq struct {
	TenantId string `json:"-"`
	Group    string `json:"group,omitempty"`
}

type ListSheetGroupResp struct {
	Sheets []map[string]any `json:"sheets,omitempty"`
}

type LoadSheetReq struct {
	TenantId    string             `json:"-"`
	Group       string             `json:"group,omitempty"`
	SheetId     int64              `json:"sheet_id,omitempty"`
	View        string             `json:"view,omitempty"`
	FilterConds []store.FilterCond `json:"filter_conds,omitempty"`
}

type LoadSheetResp struct {
	Columns []map[string]any `json:"columns,omitempty"`
	Cells   []map[string]any `json:"cells,omitempty"`
}
