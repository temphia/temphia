package dyndb

import (
	"time"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

const (
	SheetTable       = "sheets"
	SheetColumnTable = "scols"
	SheetRowTable    = "srows"
	SheetCellTable   = "scells"
)

type ListSheetGroupReq struct {
	TenantId string `json:"-"`
	Group    string `json:"group,omitempty"`
}

type ListSheetGroupResp struct {
	Sheets       []map[string]any `json:"sheets,omitempty"`
	FolderTicket string           `json:"folder_ticket,omitempty"`
}

type LoadSheetReq struct {
	TenantId    string `json:"-"`
	Group       string `json:"group,omitempty"`
	SheetId     int64  `json:"sheet_id,omitempty"`
	View        string `json:"view,omitempty"`
	RowCursorId int64  `json:"row_cursor_id,omitempty"`
	Desc        bool   `json:"desc,omitempty"`
}

type FTSQuerySheet struct {
	TenantId   string `json:"-"`
	Group      string `json:"group,omitempty"`
	SheetId    int64  `json:"sheet_id,omitempty"`
	SearchTerm string `json:"search_term,omitempty"`
	Count      int32  `json:"count,omitempty"`
	ColumnId   int64  `json:"column_id,omitempty"`
	ColumnType string `json:"column_type,omitempty"`
}

type RefQuerySheet struct {
	TenantId      string `json:"-"`
	Group         string `json:"group,omitempty"`
	SheetId       int64  `json:"sheet_id,omitempty"`
	ColumnId      int64  `json:"column_id,omitempty"`
	RowCursorId   int64  `json:"row_cursor_id,omitempty"`
	TargetSource  string `json:"target_source,omitempty"`
	TargetGroup   string `json:"target_group,omitempty"`
	TargetSheetId int64  `json:"target_sheet_id,omitempty"`
}

type QuerySheetReq struct {
	TenantId    string       `json:"-"`
	Group       string       `json:"group,omitempty"`
	SheetId     int64        `json:"sheet_id,omitempty"`
	View        string       `json:"view,omitempty"`
	FilterConds []FilterCond `json:"filter_conds,omitempty"`
	RowCursorId int64        `json:"row_cursor_id,omitempty"`
	Desc        bool         `json:"desc,omitempty"`
}

type QuerySheetResp struct {
	Cells   []map[string]any `json:"cells,omitempty"`
	Columns []map[string]any `json:"columns,omitempty"`
}

type LoadSheetResp struct {
	Columns           []map[string]any      `json:"columns,omitempty"`
	Cells             []map[string]any      `json:"cells,omitempty"`
	WidgetApps        []*entities.TargetApp `json:"widget_apps,omitempty"`
	ReverseRefColumns []map[string]any      `json:"reverse_ref_cols,omitempty"`
}

type Relation struct {
	SheetId int64            `json:"sheet_id,omitempty"`
	Columns []map[string]any `json:"columns,omitempty"`
	Cells   []map[string]any `json:"cells,omitempty"`
}

type ExportOptions struct {
	TenantId string  `json:"-"`
	Group    string  `json:"group,omitempty"`
	Sheets   []int64 `json:"sheets,omitempty"`
}

type ImportOptions struct {
	TenantId string  `json:"-"`
	Group    string  `json:"group,omitempty"`
	Sheets   []int64 `json:"sheets,omitempty"`
}

type ExportData struct {
	Source    string              `json:"source,omitempty"`
	Group     string              `json:"group,omitempty"`
	Date      time.Time           `json:"date,omitempty"`
	SheetData map[int64]SheetData `json:"sheet_date,omitempty"`
}

type SheetData struct {
	Id      int64                    `json:"id,omitempty"`
	Name    string                   `json:"name,omitempty"`
	Columns []Column                 `json:"columns,omitempty"`
	Cells   map[int64]map[int64]Cell `json:"cells,omitempty"`
}

type Column struct {
	Id        int64  `json:"__id,omitempty"`
	Version   int64  `json:"__version,omitempty"`
	Name      string `json:"name,omitempty"`
	Ctype     string `json:"ctype,omitempty"`
	Color     string `json:"color,omitempty"`
	Options   string `json:"opts,omitempty"`
	RefColumn int64  `json:"refcolumn,omitempty"`
	RefSheet  int64  `json:"refsheet,omitempty"`
	SheetId   int64  `json:"sheetid,omitempty"`
	ExtraOpts any    `json:"extraopts,omitempty"`
}

type Cell struct {
	Id      int64  `json:"__id,omitempty"`
	Version int64  `json:"__version,omitempty"`
	ColId   int64  `json:"colid,omitempty"`
	RowId   int64  `json:"rowid,omitempty"`
	SheetId int64  `json:"sheetid,omitempty"`
	Value   string `json:"value,omitempty"`
	NumVal  int64  `json:"numval,omitempty"`
	Color   string `json:"color,omitempty"`
}
