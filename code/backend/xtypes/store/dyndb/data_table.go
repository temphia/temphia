package dyndb

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type ModCtx struct {
	UserId    string `json:"user_id,omitempty"`
	UserSign  string `json:"user_sign,omitempty"`
	InitSign  string `json:"init_sign,omitempty"`
	TableName string `json:"table_name,omitempty"`
}

func (m *ModCtx) JSON() ([]byte, error) {
	return json.Marshal(m)
}

type NewRowReq struct {
	TenantId string         `json:"-"`
	Group    string         `json:"group,omitempty"`
	Table    string         `json:"table,omitempty"`
	Data     map[string]any `json:"data,omitempty"`
	ModCtx   ModCtx         `json:"mod_ctx,omitempty"`
	NoReact  bool           `json:"no_react,omitempty"`
}

type NewBatchRowReq struct {
	TenantId string           `json:"-"`
	Group    string           `json:"group,omitempty"`
	Table    string           `json:"table,omitempty"`
	Data     []map[string]any `json:"data,omitempty"`
	ModCtx   ModCtx           `json:"mod_ctx,omitempty"`
	NoReact  bool             `json:"no_react,omitempty"`
}

type GetRowReq struct {
	TenantId  string `json:"-"`
	Group     string `json:"group,omitempty"`
	Table     string `json:"table,omitempty"`
	Id        int64  `json:"id,omitempty"`
	SkipCache bool   `json:"skip_cache,omitempty"`
}

type UpdateRowReq struct {
	TenantId string         `json:"-"`
	Id       int64          `json:"id,omitempty"`
	Version  int64          `json:"version,omitempty"`
	Group    string         `json:"group,omitempty"`
	Table    string         `json:"table,omitempty"`
	Data     map[string]any `json:"data,omitempty"`
	ModCtx   ModCtx         `json:"mod_ctx,omitempty"`
	NoReact  bool           `json:"no_react,omitempty"`
}

type DeleteRowReq struct {
	TenantId string `json:"-"`
	Group    string `json:"group,omitempty"`
	Table    string `json:"table,omitempty"`
	Id       int64  `json:"id,omitempty"`
	ModCtx   ModCtx `json:"mod_ctx,omitempty"`
	NoReact  bool   `json:"no_react,omitempty"`
}

type DeleteRowMultiReq struct {
	TenantId string  `json:"-"`
	Group    string  `json:"group,omitempty"`
	Table    string  `json:"table,omitempty"`
	Ids      []int64 `json:"sid,omitempty"`
	ModCtx   ModCtx  `json:"mod_ctx,omitempty"`
	NoReact  bool    `json:"no_react,omitempty"`
}

type DeleteRowBatchReq struct {
	TenantId    string       `json:"-"`
	Group       string       `json:"group,omitempty"`
	Table       string       `json:"table,omitempty"`
	FilterConds []FilterCond `json:"filter_conds,omitempty"`
	ModCtx      ModCtx       `json:"mod_ctx,omitempty"`
	NoReact     bool         `json:"no_react,omitempty"`
}

type LoadTableReq struct {
	TenantId    string       `json:"-"`
	Table       string       `json:"table,omitempty"`
	Group       string       `json:"group,omitempty"`
	View        string       `json:"view,omitempty"`
	ViewFilters []FilterCond `json:"view_filters,omitempty"`
}

type LoadTableResp struct {
	ReverseRefs   []*entities.Column    `json:"reverse_refs,omitempty"`
	Views         []*entities.DataView  `json:"views,omitempty"`
	DataWidgets   []*entities.TargetApp `json:"data_widgets,omitempty"`
	ActiveView    string                `json:"active_view,omitempty"`
	FolderTickets map[string]string     `json:"folder_tickets,omitempty"`
	UserTickets   map[string]string     `json:"user_tickets,omitempty"`
	QueryResponse *QueryResult          `json:"query_response,omitempty"`
}

type SimpleQueryReq struct {
	TenantId    string       `json:"-"`
	Table       string       `json:"table,omitempty"`
	Group       string       `json:"group,omitempty"`
	Count       int64        `json:"count,omitempty"`
	FilterConds []FilterCond `json:"filter_conds,omitempty"`
	Page        int64        `json:"page,omitempty"`
	Selects     []string     `json:"selects,omitempty"`
	OrderBy     string       `json:"order_by,omitempty"`
	Desc        bool         `json:"desc,omitempty"`
}

type JoinReq struct {
	TenantId      string       `json:"-"`
	Group         string       `json:"group,omitempty"`
	Parent        string       `json:"parent,omitempty"`
	Child         string       `json:"child,omitempty"`
	OnParent      string       `json:"on_parent,omitempty"`
	OnChild       string       `json:"on_child,omitempty"`
	ParentFilters []FilterCond `json:"parent_ft,omitempty"`
	ChildFilters  []FilterCond `json:"child_ft,omitempty"`
	OrderBy       string       `json:"order_by,omitempty"`
	Desc          bool         `json:"desc,omitempty"`
}

type MultiJoinReq struct {
	TenantId      string         `json:"-"`
	Group         string         `json:"group,omitempty"`
	Parent        string         `json:"parent,omitempty"`
	OnParent      string         `json:"on_parent,omitempty"`
	ParentFilters []FilterCond   `json:"parent_ft,omitempty"`
	Fragments     []JoinFragment `json:"fragments,omitempty"`
	OrderBy       string         `json:"order_by,omitempty"`
	Desc          bool           `json:"desc,omitempty"`
}

type JoinFragment struct {
	Name     string       `json:"name,omitempty"`
	OnColumn string       `json:"on_column,omitempty"`
	Filters  []FilterCond `json:"filters,omitempty"`
}

type MultiJoinResult struct {
	Rows    []map[string]any            `json:"rows"`
	Columns map[string]*entities.Column `json:"columns,omitempty"`
}

type JoinResult struct {
	Rows       []map[string]any            `json:"rows"`
	ParentCols map[string]*entities.Column `json:"parent_columns,omitempty"`
	ChildCols  map[string]*entities.Column `json:"child_columns,omitempty"`
}

type FTSQueryReq struct {
	TenantId     string       `json:"-"`
	Table        string       `json:"table,omitempty"`
	Group        string       `json:"group,omitempty"`
	SearchTerm   string       `json:"search_term,omitempty"`
	SearchColumn string       `json:"search_column,omitempty"`
	Count        int64        `json:"count,omitempty"`
	Filters      []FilterCond `json:"filters,omitempty"`
	Page         int64        `json:"page,omitempty"`
	UsePattern   bool         `json:"use_pattern,omitempty"`
}

type RefLoadReq struct {
	Column      string `json:"column,omitempty"`
	Type        string `json:"type,omitempty"`
	Target      string `json:"target,omitempty"`
	Object      string `json:"object,omitempty"`
	CursorRowId int64  `json:"cursor_row_id,omitempty"`
}

type RevRefLoadReq struct {
	CurrentTable string `json:"current_table,omitempty"`
	TargetTable  string `json:"target_table,omitempty"`
	TargetColumn string `json:"column,omitempty"`
	CurrentItem  any    `json:"current_item,omitempty"`
	CursorRowId  int64  `json:"cursor_row_id,omitempty"`
	Count        int    `json:"count,omitempty"`
}

type RefResolveReq struct {
	Column string `json:"column,omitempty"`
	Type   string `json:"type,omitempty"`
	Target string `json:"target,omitempty"`
	Object string `json:"object,omitempty"`
	RowIds []any  `json:"row_ids,omitempty"`
}

type FilterCond struct {
	Column string `json:"column,omitempty"`
	Cond   string `json:"cond,omitempty"`
	Value  any    `json:"value,omitempty"`
	Target string `json:"target,omitempty"`
}

type TemplateQueryReq struct {
	TenantId  string                    `json:"-"`
	Group     string                    `json:"group,omitempty"`
	Fragments map[string]map[string]any `json:"fragments,omitempty"`
	Name      string                    `json:"name,omitempty"`
}

type SqlQueryReq struct {
	NoTransform bool   `json:"no_transform,omitempty"`
	Raw         bool   `json:"raw,omitempty"`
	Group       string `json:"group,omitempty"`
	QStr        string `json:"qstr,omitempty"`
}

type SqlQueryResult struct {
	Records any                         `json:"records,omitempty"`
	Columns map[string]*entities.Column `json:"columns,omitempty"`
}

type QueryResult struct {
	Count   int64                       `json:"count,omitempty"`
	Page    int64                       `json:"page,omitempty"`
	Rows    []map[string]any            `json:"rows"`
	Columns map[string]*entities.Column `json:"columns,omitempty"`
	Final   bool                        `json:"final,omitempty"`
}

type LoadDgroupResp struct {
	Tables       []*entities.Table `json:"tables,omitempty"`
	FolderTicket string            `json:"folder_ticket,omitempty"`
}
