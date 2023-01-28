package datahub

import "github.com/temphia/temphia/code/backend/xtypes/store"

type ListSheetGroupReq struct {
	TenantId string `json:"-"`
	Table    string `json:"table,omitempty"`
	Group    string `json:"group,omitempty"`
}

type ListSheetGroupResp struct {
	Sheets []map[string]any `json:"sheets,omitempty"`
}

type LoadSheetReq struct {
	TenantId    string             `json:"-"`
	Table       string             `json:"table,omitempty"`
	Group       string             `json:"group,omitempty"`
	SheetId     string             `json:"sheet_id,omitempty"`
	View        string             `json:"view,omitempty"`
	FilterConds []store.FilterCond `json:"filter_conds,omitempty"`
}

type LoadSheetResp struct {
	Columns []map[string]any `json:"columns,omitempty"`
	Rows    []map[string]any `json:"rows,omitempty"`
	Cells   []map[string]any `json:"cells,omitempty"`
}

func (d *dynSource) ListSheetGroup(opts ListSheetGroupReq) (*ListSheetGroupResp, error) {

	return nil, nil
}

func (d *dynSource) LoadSheetGroup(opts LoadSheetReq) (*LoadSheetResp, error) {

	return nil, nil
}
