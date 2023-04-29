package sheet

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/services/datahub/handle"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/filter"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

var _ dyndb.DataSheetHub = (*Sheet)(nil)

type Sheet struct {
	handle   *handle.Handle
	source   string
	tenantId string
	group    string

	tableHub dyndb.DataTableHub
}

func New(tableHub dyndb.DataTableHub, handle *handle.Handle, source string, tenantId string, group string) *Sheet {
	return &Sheet{
		handle:   handle,
		source:   source,
		tenantId: tenantId,
		group:    group,
		tableHub: tableHub,
	}

}

func (s *Sheet) ListSheetGroup(txid uint32) (*dyndb.ListSheetGroupResp, error) {

	sheetRows, err := s.tableHub.SimpleQuery(txid, dyndb.SimpleQueryReq{
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
	columns, err := s.tableHub.SimpleQuery(txid, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		FilterConds: []dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   filter.FilterEqual,
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

	cursorFilter := filter.FilterGT
	if data.Desc {
		cursorFilter = filter.FilterLT
	}

	cells, err := s.tableHub.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetCellTable,
		FilterConds: []dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   filter.FilterEqual,
				Value:  data.SheetId,
			},
			{
				Column: "rowid",
				Cond:   cursorFilter,
				Value:  data.RowCursorId,
			},
		},
		Desc:    data.Desc,
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

	rowCells := trimIncompleteCells(cells.Rows, colNo, int(count))

	apps, _ := s.handle.CoreHub.ListTargetAppByType(s.tenantId, entities.TargetAppTypeDataSheetWidget, fmt.Sprintf("%s/%s/%d", s.source, s.group, data.SheetId))

	fresp := &dyndb.LoadSheetResp{
		Columns:           columns.Rows,
		Cells:             rowCells,
		WidgetApps:        apps,
		ReverseRefColumns: nil,
	}

	refresp, err := s.tableHub.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		FilterConds: []dyndb.FilterCond{
			{
				Column: "refsheet",
				Cond:   filter.FilterEqual,
				Value:  data.SheetId,
			},
		},
	})

	if err == nil {
		fresp.ReverseRefColumns = refresp.Rows
	}

	return fresp, nil

}

func (s *Sheet) RefQuery(txid uint32, data *dyndb.RefQuerySheet) (*dyndb.QuerySheetResp, error) {

	// fixme => make sure it could access these by checking original ref column.sheet

	if data.TargetSource == "" {
		data.TargetSource = s.source
	}

	var hub dyndb.DataSheetHub

	if data.TargetGroup == "" {
		hub = s
	} else {
		hub = s.handle.MainHub.GetDataSheetHub(data.TargetSource, data.TenantId, data.TargetGroup)
	}

	return hub.Query(txid, &dyndb.QuerySheetReq{
		TenantId:    data.TenantId,
		Group:       data.Group,
		SheetId:     data.TargetSheetId,
		RowCursorId: data.RowCursorId,
	})
}

func (s *Sheet) Query(txid uint32, data *dyndb.QuerySheetReq) (*dyndb.QuerySheetResp, error) {

	frags := []dyndb.JoinFragment{}

	cursorFilter := filter.FilterGT
	if data.Desc {
		cursorFilter = filter.FilterLT
	}

	for _, fc := range data.FilterConds {

		textFilter := func() {
			frags = append(frags, dyndb.JoinFragment{
				Name:     dyndb.SheetCellTable,
				OnColumn: "rowid",
				Filters: []dyndb.FilterCond{
					{
						Column: "sheetid",
						Cond:   filter.FilterEqual,
						Value:  data.SheetId,
					},
					{
						Column: "value",
						Cond:   fc.Cond,
						Value:  fc.Value,
					},
					{
						Column: "rowid",
						Cond:   cursorFilter,
						Value:  data.RowCursorId,
					},
				},
			})
		}

		numFilter := func() {
			frags = append(frags, dyndb.JoinFragment{
				Name:     dyndb.SheetCellTable,
				OnColumn: "rowid",
				Filters: []dyndb.FilterCond{
					{
						Column: "sheetid",
						Cond:   filter.FilterEqual,
						Value:  data.SheetId,
					},
					{
						Column: "numval",
						Cond:   fc.Cond,
						Value:  fc.Value,
					},
					{
						Column: "rowid",
						Cond:   cursorFilter,
						Value:  data.RowCursorId,
					},
				},
			})
		}

		switch fc.Cond {
		case filter.FilterAround, filter.FilterNotAround:
			panic("Not supported")
		case filter.FilterNumEqual,
			filter.FilterNumNotEqual,
			filter.FilterNumIn,
			filter.FilterNumNotIn,
			filter.FilterLT,
			filter.FilterGT,
			filter.FilterLTE,
			filter.FilterGTE:
			numFilter()
		default:
			textFilter()
		}
	}

	result, err := s.tableHub.MultiJoinQuery(txid, dyndb.MultiJoinReq{
		TenantId: data.TenantId,
		Group:    data.Group,
		Parent:   dyndb.SheetCellTable,
		ParentFilters: []dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   filter.FilterEqual,
				Value:  data.SheetId,
			},
			{
				Column: "rowid",
				Cond:   cursorFilter,
				Value:  data.RowCursorId,
			},
		},
		OnParent:  "rowid",
		Fragments: frags,
		OrderBy:   "rowid",
	})

	if err != nil {
		return nil, err
	}

	cresp, err := s.tableHub.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		FilterConds: []dyndb.FilterCond{
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

	return &dyndb.QuerySheetResp{
		Cells:   result.Rows,
		Columns: cresp.Rows,
	}, nil
}

func (s *Sheet) ListSheet(txid uint32) ([]map[string]any, error) {
	resp, err := s.tableHub.SimpleQuery(txid, dyndb.SimpleQueryReq{
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
	_, err := s.tableHub.NewRow(txid, dyndb.NewRowReq{
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
	return s.tableHub.GetRow(txid, dyndb.GetRowReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetTable,
		Id:       id,
	})
}

func (s *Sheet) UpdateSheet(txid uint32, id int64, userId string, data map[string]any) error {
	_, err := s.tableHub.UpdateRow(0, dyndb.UpdateRowReq{
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
	_, err := s.tableHub.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetCellTable,
		FilterConds: []dyndb.FilterCond{
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

	s.tableHub.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetRowTable,
		FilterConds: []dyndb.FilterCond{
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

	s.tableHub.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		FilterConds: []dyndb.FilterCond{
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

	return s.tableHub.DeleteRow(0, dyndb.DeleteRowReq{
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

	resp, err := s.tableHub.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		FilterConds: []dyndb.FilterCond{
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

	return s.tableHub.NewRow(0, dyndb.NewRowReq{
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
	return s.tableHub.GetRow(0, dyndb.GetRowReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		Id:       cid,
	})
}

func (s *Sheet) UpdateSheetColumn(txid uint32, sid, cid int64, userId string, data map[string]any) error {

	_, err := s.tableHub.UpdateRow(txid, dyndb.UpdateRowReq{
		TenantId: s.tenantId,
		Id:       cid,
		Group:    s.group,
		Version:  -1,
		Table:    dyndb.SheetColumnTable,
		Data:     data,
		ModCtx: dyndb.ModCtx{
			UserId: userId,
		},
	})

	return err
}

func (s *Sheet) DeleteSheetColumn(txid uint32, sid, cid int64, userId string) error {

	_, err := s.tableHub.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetCellTable,
		FilterConds: []dyndb.FilterCond{
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

	return s.tableHub.DeleteRow(txid, dyndb.DeleteRowReq{
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

	rid, err := s.tableHub.NewRow(txid, dyndb.NewRowReq{
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

	finalCells := make([]map[string]any, 0)

	for cid, cellData := range data {
		cellData["rowid"] = rid
		cellData["sheetid"] = sid
		cellData["colid"] = cid

		finalCells = append(finalCells, cellData)

	}

	_, err = s.tableHub.NewBatchRows(txid, dyndb.NewBatchRowReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetCellTable,
		Data:     finalCells,
		ModCtx: dyndb.ModCtx{
			UserId:   userId,
			AltIdent: fmt.Sprint(rid),
		},
	})

	return nil, err
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

			_, err := s.tableHub.NewRow(0, dyndb.NewRowReq{
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

			_, err := s.tableHub.UpdateRow(0, dyndb.UpdateRowReq{
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

func (s *Sheet) DeleteRowWithCell(txid uint32, sid, rid int64, userId string) error {

	_, err := s.tableHub.DeleteRowBatch(txid, dyndb.DeleteRowBatchReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetCellTable,
		FilterConds: []dyndb.FilterCond{
			{
				Column: "rowid",
				Cond:   "equal",
				Value:  rid,
			},
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  sid,
			},
		},
	})
	if err != nil {
		return err
	}

	return s.tableHub.DeleteRow(txid, dyndb.DeleteRowReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetRowTable,
		Id:       rid,
		ModCtx: dyndb.ModCtx{
			UserId:   userId,
			UserSign: "",
		},
	})
}

func (s *Sheet) GetRowRelations(txid uint32, sid, rid, refsheet, refcol int64) (*dyndb.Relation, error) {

	cresp, err := s.tableHub.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		FilterConds: []dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  refsheet,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	refresp, err := s.tableHub.JoinQuery(0, dyndb.JoinReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Parent:   dyndb.SheetCellTable,
		Child:    dyndb.SheetCellTable,
		OnParent: "rowid",
		OnChild:  "rowid",
		ParentFilters: []dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  refsheet,
			},
			{
				Column: "colid",
				Cond:   "equal",
				Value:  refcol,
			},
		},
		ChildFilters: []dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  refsheet,
			},
		},
		OrderBy: "rowid",
	})

	if err != nil {
		return nil, err
	}

	return &dyndb.Relation{
		SheetId: refsheet,
		Columns: cresp.Rows,
		Cells:   refresp.Rows,
	}, nil
}

func (s *Sheet) FTSQuery(txid uint32, req *dyndb.FTSQuerySheet) (*dyndb.QuerySheetResp, error) {

	cresp, err := s.tableHub.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Table:    dyndb.SheetColumnTable,
		FilterConds: []dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  req.SheetId,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	refresp, err := s.tableHub.JoinQuery(0, dyndb.JoinReq{
		TenantId: s.tenantId,
		Group:    s.group,
		Parent:   dyndb.SheetCellTable,
		Child:    dyndb.SheetCellTable,
		OnParent: "rowid",
		OnChild:  "rowid",
		ParentFilters: []dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  req.SheetId,
			},
			{
				Column: "value",
				Cond:   filter.FilterContains,
				Value:  req.SearchTerm,
			},
		},
		ChildFilters: []dyndb.FilterCond{
			{
				Column: "sheetid",
				Cond:   "equal",
				Value:  req.SheetId,
			},
		},
		OrderBy: "rowid",
	})

	if err != nil {
		return nil, err
	}

	return &dyndb.QuerySheetResp{
		Cells:   refresp.Rows,
		Columns: cresp.Rows,
	}, nil
}

func trimIncompleteCells(rows []map[string]any, colno, count int) []map[string]any {
	if len(rows) == int(count) {
		// remove last incomplete cells of a row
		lastrowId := rows[len(rows)-1]["rowid"].(int64)
		offset := (len(rows) - 1) - colno
		tail := rows[offset:]
		for _, cr := range tail {
			crowId := cr["rowid"].(int64)
			if crowId == lastrowId {
				break
			}
			offset = offset + 1
		}
		return rows[:offset]
	}

	return rows
}
