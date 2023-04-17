package dyndb

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

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
}

type QuerySheetReq struct {
	TenantId    string `json:"-"`
	Group       string `json:"group,omitempty"`
	SheetId     int64  `json:"sheet_id,omitempty"`
	View        string `json:"view,omitempty"`
	RowCursorId int64  `json:"row_cursor_id,omitempty"`
	Desc        bool   `json:"desc,omitempty"`
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
