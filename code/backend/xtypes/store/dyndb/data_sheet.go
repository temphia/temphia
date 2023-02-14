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
	TenantId    string       `json:"-"`
	Group       string       `json:"group,omitempty"`
	SheetId     int64        `json:"sheet_id,omitempty"`
	View        string       `json:"view,omitempty"`
	FilterConds []FilterCond `json:"filter_conds,omitempty"`
}

type LoadSheetResp struct {
	Columns    []map[string]any      `json:"columns,omitempty"`
	Cells      []map[string]any      `json:"cells,omitempty"`
	WidgetApps []*entities.TargetApp `json:"widget_apps,omitempty"`
}
