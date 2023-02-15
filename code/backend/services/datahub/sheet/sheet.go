package sheet

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/services/datahub/handle"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

var _ dyndb.DataSheetHub = (*Sheet)(nil)

type Sheet struct {
	inner    dyndb.DynDB
	handle   *handle.Handle
	source   string
	tenantId string
	group    string
}

func New(inner dyndb.DynDB, handle *handle.Handle, source string, tenantId string, group string) *Sheet {
	return &Sheet{
		inner:    inner,
		handle:   handle,
		source:   source,
		tenantId: tenantId,
		group:    group,
	}

}

func (s *Sheet) ListSheetGroup(txid uint32) (*dyndb.ListSheetGroupResp, error) {

	sheetRows, err := s.inner.SimpleQuery(txid, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetTable,
	})
	if err != nil {
		return nil, err
	}

	return &dyndb.ListSheetGroupResp{
		Sheets:       sheetRows.Rows,
		FolderTicket: "",
	}, nil
}

func (s *Sheet) LoadSheet(txid uint32, data *dyndb.LoadSheetReq) (*dyndb.LoadSheetResp, error) {
	columns, err := s.inner.SimpleQuery(txid, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Group:    s.group,
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
		return &dyndb.LoadSheetResp{
			Columns: columns.Rows,
			Cells:   []map[string]any{},
		}, nil
	}

	colNo := len(columns.Rows)
	count := int64((dyndb.DefaultQueryFetchCount * colNo) + colNo)

	cells, err := s.inner.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Group:    s.group,
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
		return &dyndb.LoadSheetResp{
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

	return &dyndb.LoadSheetResp{
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
			TenantId:    s.tenantId,
		}},
	}, nil

}

func (s *Sheet) ListSheet(txid uint32) ([]map[string]any, error) {
	resp, err := s.inner.SimpleQuery(txid, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Table:    dyndb.SheetTable,
		Group:    s.group,
	})

	if err != nil {
		return nil, err
	}

	return resp.Rows, nil
}

func (s *Sheet) NewSheet(txid uint32, userId string, data map[string]any) error {
	_, err := s.inner.NewRow(txid, dyndb.NewRowReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetTable,
		Data:     data,
		ModCtx: dyndb.ModCtx{
			UserId: userId,
		},
	})
	return err
}

func (s *Sheet) GetSheet(txid uint32, id int64) (map[string]any, error) {
	return s.inner.GetRow(txid, dyndb.GetRowReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetTable,
		Id:       id,
	})
}

func (s *Sheet) UpdateSheet(txid uint32, id int64, userId string, data map[string]any) error {
	_, err := s.inner.UpdateRow(0, dyndb.UpdateRowReq{
		TenantId: s.tenantId,
		Id:       id,
		Group:    s.group,
		Table:    dyndb.SheetTable,
		Data:     data,
		ModCtx: dyndb.ModCtx{
			UserId: userId,
		},
	})

	return err

}

func (s *Sheet) DeleteSheet(txid uint32, id int64, userId string) error {
	err := s.inner.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId: s.tenantId,
		Group:    s.group,
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

	s.inner.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId: s.tenantId,
		Group:    s.group,
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

	s.inner.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId: s.tenantId,
		Group:    s.group,
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

	return s.inner.DeleteRow(0, dyndb.DeleteRowReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetTable,
		Id:       id,
		ModCtx: dyndb.ModCtx{
			UserId: userId,
		},
	})
}

func (s *Sheet) ListSheetColumn(txid uint32, sid int64) ([]map[string]any, error) {

	resp, err := s.inner.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Group:    s.group,
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

func (s *Sheet) NewSheetColumn(txid uint32, sid int64, userId string, data map[string]any) (int64, error) {
	data["sheetid"] = sid

	return s.inner.NewRow(0, dyndb.NewRowReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		Data:     data,
		ModCtx: dyndb.ModCtx{
			UserId: userId,
		},
	})
}

func (s *Sheet) GetSheetColumn(txid uint32, sid, cid int64) (map[string]any, error) {
	return s.inner.GetRow(0, dyndb.GetRowReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		Id:       cid,
	})
}

func (s *Sheet) UpdateSheetColumn(txid uint32, sid, cid int64, userId string, data map[string]any) error {

	_, err := s.inner.UpdateRow(txid, dyndb.UpdateRowReq{
		TenantId: s.tenantId,
		Id:       cid,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		Data:     data,
		ModCtx: dyndb.ModCtx{
			UserId: userId,
		},
	})

	return err
}

func (s *Sheet) DeleteSheetColumn(txid uint32, sid, cid int64, userId string) error {

	err := s.inner.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetCellTable,
		FilterConds: []*dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  sid,
			},
			{
				Column: "colid",
				Cond:   "equal",
				Value:  cid,
			},
		},
	})

	if err != nil {
		pp.Println("@err while clearing cells")
	}

	return s.inner.DeleteRow(txid, dyndb.DeleteRowReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		Id:       cid,
		ModCtx: dyndb.ModCtx{
			UserId: userId,
		},
	})
}

func (s *Sheet) NewRowWithCell(txid uint32, sid int64, userId string, data map[int64]map[string]any) (map[int64]map[string]any, error) {

	rid, err := s.inner.NewRow(txid, dyndb.NewRowReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetRowTable,
		Data: map[string]any{
			"sheetid": sid,
		},
		ModCtx: dyndb.ModCtx{
			UserId: userId,
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

		cellid, err := s.inner.NewRow(txid, dyndb.NewRowReq{
			TenantId: s.tenantId,
			Group:    s.group,
			Table:    dyndb.SheetCellTable,
			Data:     cellData,
			ModCtx: dyndb.ModCtx{
				UserId: userId,
			},
		})
		if err != nil {
			return nil, err
		}
		pp.Println(cellid)

	}

	return nil, nil
}

func (s *Sheet) UpdateRowWithCell(txid uint32, sid, rid int64, userId string, data map[int64]map[string]any) (map[int64]map[string]any, error) {

	for colid, cellData := range data {

		pp.Println("@data", cellData)

		cellId, cellOk := cellData[dyndb.KeyPrimary].(float64)
		version, _ := cellData[dyndb.KeyVersion].(float64)
		if !cellOk {
			cellData["rowid"] = rid
			cellData["sheetid"] = sid
			cellData["colid"] = colid

			_, err := s.inner.NewRow(0, dyndb.NewRowReq{
				TenantId: s.tenantId,
				Group:    s.group,
				Table:    dyndb.SheetCellTable,
				Data:     cellData,
				ModCtx: dyndb.ModCtx{
					UserId: userId,
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

			_, err := s.inner.UpdateRow(0, dyndb.UpdateRowReq{
				TenantId: s.tenantId,
				Id:       int64(cellId),
				Group:    s.group,
				Table:    dyndb.SheetCellTable,
				Data:     cellData,
				Version:  int64(version),
				ModCtx: dyndb.ModCtx{
					UserId: userId,
				},
			})
			if err != nil {
				return nil, err
			}
		}
	}

	return nil, nil
}
