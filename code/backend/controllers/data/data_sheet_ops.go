package data

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

func (c *Controller) ListSheetGroup(uclaim *claim.Data) (*ListSheetGroupResp, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	tg, err := dynDb.GetGroup(uclaim.DataGroup)
	if err != nil {
		return nil, err
	}

	ftok, err := c.folderTicket(tg, uclaim)
	if err != nil {
		return nil, err
	}

	sheetRows, err := dynDb.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetTable,
	})
	if err != nil {
		return nil, err
	}

	return &ListSheetGroupResp{
		Sheets:       sheetRows.Rows,
		FolderTicket: ftok,
	}, nil
}

func (c *Controller) LoadSheet(uclaim *claim.Data, data *LoadSheetReq) (*LoadSheetResp, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	columns, err := dynDb.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetColumnTable,
		FilterConds: []*dyndb.FilterCond{
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

	if len(columns.Rows) == 0 {
		return &LoadSheetResp{
			Columns: columns.Rows,
			Cells:   []map[string]any{},
		}, nil
	}

	colNo := len(columns.Rows)
	count := int64((dyndb.DefaultQueryFetchCount * colNo) + colNo)

	cells, err := dynDb.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetCellTable,
		FilterConds: []*dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  data.SheetId,
			},
		},

		Count:   count,
		OrderBy: "rowid",
	})
	if err != nil {
		return nil, err
	}

	if len(cells.Rows) == 0 {
		return &LoadSheetResp{
			Columns: columns.Rows,
			Cells:   cells.Rows,
		}, nil
	}

	rowCells := cells.Rows

	if len(cells.Rows) == int(count) {
		// remove last incomplete cells of a row
		lastrowId := cells.Rows[len(cells.Rows)-1]["rowid"].(int64)
		offset := (len(cells.Rows) - 1) - colNo

		for _, cr := range cells.Rows[:offset] {
			crowId := cr["rowid"].(int64)
			if crowId == lastrowId {
				break
			}
			offset = offset + 1
		}

		rowCells = cells.Rows[:offset]
	}

	return &LoadSheetResp{
		Columns: columns.Rows,
		Cells:   rowCells,
		WidgetApps: []*entities.TargetApp{{
			Id:          1,
			Name:        "test",
			Icon:        "",
			Policy:      "",
			TargetType:  entities.TargetAppTypeDataSheetWidget,
			Target:      "default/group/table",
			ContextType: "todo.1",
			PlugId:      "test1",
			AgentId:     "default",
			TenantId:    uclaim.TenantId,
		}},
	}, nil

}

// sheets

func (c *Controller) ListSheet(uclaim *claim.Data) ([]map[string]any, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	resp, err := dynDb.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: uclaim.TenantId,
		Table:    dyndb.SheetTable,
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

	_, err := dynDb.NewRow(0, dyndb.NewRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetTable,
		Data:     data,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})

	return err
}

func (c *Controller) GetSheet(uclaim *claim.Data, id int64) (map[string]any, error) {
	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDb.GetRow(0, dyndb.GetRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetTable,
		Id:       id,
	})

}

func (c *Controller) UpdateSheet(uclaim *claim.Data, id int64, data map[string]any) error {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	_, err := dynDb.UpdateRow(0, dyndb.UpdateRowReq{
		TenantId: uclaim.TenantId,
		Id:       id,
		Group:    group,
		Table:    dyndb.SheetTable,
		Data:     data,
	})

	return err

}

func (c *Controller) DeleteSheet(uclaim *claim.Data, id int64) error {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	err := dynDb.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetCellTable,
		FilterConds: []*dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  id,
			},
		},
	})
	if err != nil {
		pp.Println("@err while clearing cells")
	}

	dynDb.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetRowTable,
		FilterConds: []*dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  id,
			},
		},
	})

	if err != nil {
		pp.Println("@err while clearing rows")
	}

	dynDb.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetColumnTable,
		FilterConds: []*dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  id,
			},
		},
	})

	if err != nil {
		pp.Println("@err while clearing columns")
	}

	dynDb.DeleteRow(0, dyndb.DeleteRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetTable,
		Id:       id,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})

	return nil
}

// columns

func (c *Controller) ListSheetColumn(uclaim *claim.Data, sid int64) ([]map[string]any, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	resp, err := dynDb.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetColumnTable,
		FilterConds: []*dyndb.FilterCond{
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

	data["sheetid"] = sid

	return dynDb.NewRow(0, dyndb.NewRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetColumnTable,
		Data:     data,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})

}

func (c *Controller) GetSheetColumn(uclaim *claim.Data, sid, cid int64) (map[string]any, error) {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	dynDb.GetRow(0, dyndb.GetRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetColumnTable,
		Id:       cid,
	})

	return nil, nil
}

func (c *Controller) UpdateSheetColumn(uclaim *claim.Data, sid, cid int64, data map[string]any) error {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	dynDb.UpdateRow(0, dyndb.UpdateRowReq{
		TenantId: uclaim.TenantId,
		Id:       cid,
		Group:    group,
		Table:    dyndb.SheetColumnTable,
		Data:     data,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})

	return nil

}

func (c *Controller) DeleteSheetColumn(uclaim *claim.Data, sid, cid int64) error {

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	return dynDb.DeleteRow(0, dyndb.DeleteRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetColumnTable,
		Id:       cid,
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})
}

// row_cells

func (c *Controller) NewRowWithCell(uclaim *claim.Data, sid int64, data map[int64]map[string]any) (map[int64]map[string]any, error) {

	txid := uint32(0)

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	rid, err := dynDb.NewRow(txid, dyndb.NewRowReq{
		TenantId: uclaim.TenantId,
		Group:    group,
		Table:    dyndb.SheetRowTable,
		Data: map[string]any{
			"sheetid": sid,
		},
		ModCtx: dyndb.ModCtx{
			UserId: uclaim.UserID,
		},
	})
	if err != nil {
		return nil, err
	}

	// fixme => batch support

	for cid, cellData := range data {
		cellData["rowid"] = rid
		cellData["sheetid"] = sid
		cellData["colid"] = cid

		cellid, err := dynDb.NewRow(txid, dyndb.NewRowReq{
			TenantId: uclaim.TenantId,
			Group:    group,
			Table:    dyndb.SheetCellTable,
			Data:     cellData,
			ModCtx: dyndb.ModCtx{
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

	pp.Println("@update", data)

	source, group := getTarget(uclaim)
	dynDb := c.dynHub.GetSource(source, uclaim.TenantId)

	for colid, cellData := range data {

		pp.Println("@data", cellData)

		cellId, cellOk := cellData[dyndb.KeyPrimary].(float64)
		version, _ := cellData[dyndb.KeyVersion].(float64)
		if !cellOk {
			cellData["rowid"] = rid
			cellData["sheetid"] = sid
			cellData["colid"] = colid

			_, err := dynDb.NewRow(0, dyndb.NewRowReq{
				TenantId: uclaim.TenantId,
				Group:    group,
				Table:    dyndb.SheetCellTable,
				Data:     cellData,
				ModCtx: dyndb.ModCtx{
					UserId: uclaim.UserID,
				},
			})
			if err != nil {
				return nil, err
			}

		} else {

			delete(cellData, dyndb.KeyPrimary)
			delete(cellData, dyndb.KeyVersion)
			delete(cellData, "rowid")
			delete(cellData, "sheetid")
			delete(cellData, "colid")

			_, err := dynDb.UpdateRow(0, dyndb.UpdateRowReq{
				TenantId: uclaim.TenantId,
				Id:       int64(cellId),
				Group:    group,
				Table:    dyndb.SheetCellTable,
				Data:     cellData,
				Version:  int64(version),
				ModCtx: dyndb.ModCtx{
					UserId: uclaim.UserID,
				},
			})
			if err != nil {
				return nil, err
			}

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
	Sheets       []map[string]any `json:"sheets,omitempty"`
	FolderTicket string           `json:"folder_ticket,omitempty"`
}

type LoadSheetReq struct {
	TenantId    string             `json:"-"`
	Group       string             `json:"group,omitempty"`
	SheetId     int64              `json:"sheet_id,omitempty"`
	View        string             `json:"view,omitempty"`
	FilterConds []dyndb.FilterCond `json:"filter_conds,omitempty"`
}

type LoadSheetResp struct {
	Columns    []map[string]any      `json:"columns,omitempty"`
	Cells      []map[string]any      `json:"cells,omitempty"`
	WidgetApps []*entities.TargetApp `json:"widget_apps,omitempty"`
}
