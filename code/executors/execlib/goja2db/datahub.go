package goja2db

import "github.com/temphia/temphia/code/backend/xtypes/store/dyndb"

type tableQueryOptions struct {
	Source      string             `json:"source,omitempty"`
	Group       string             `json:"group,omitempty"`
	Table       string             `json:"table,omitempty"`
	Count       int64              `json:"count,omitempty"`
	FilterConds []dyndb.FilterCond `json:"filter_conds,omitempty"`
	Page        int64              `json:"page,omitempty"`
	Selects     []string           `json:"selects,omitempty"`
	OrderBy     string             `json:"order_by,omitempty"`
	Desc        bool               `json:"desc,omitempty"`
}

func (ctx *goja2db) tableQuery(opts tableQueryOptions) (any, any) {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.SimpleQuery(0, dyndb.SimpleQueryReq{
		TenantId:    ctx.tenantId,
		Table:       opts.Table,
		Group:       opts.Group,
		Count:       opts.Count,
		FilterConds: opts.FilterConds,
		Page:        opts.Page,
		Selects:     opts.Selects,
		OrderBy:     opts.OrderBy,
		Desc:        opts.Desc,
	})

}

type tablejoinQueryOptions struct {
	Source        string             `json:"source,omitempty"`
	Group         string             `json:"group,omitempty"`
	Parent        string             `json:"parent,omitempty"`
	Child         string             `json:"child,omitempty"`
	OnParent      string             `json:"on_parent,omitempty"`
	OnChild       string             `json:"on_child,omitempty"`
	ParentFilters []dyndb.FilterCond `json:"parent_ft,omitempty"`
	ChildFilters  []dyndb.FilterCond `json:"child_ft,omitempty"`
}

func (ctx *goja2db) tableJoinQuery(opts tablejoinQueryOptions) (any, any) {

	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.JoinQuery(0, dyndb.JoinReq{
		TenantId:      ctx.tenantId,
		Group:         opts.Group,
		Parent:        opts.Parent,
		Child:         opts.Child,
		OnParent:      opts.OnParent,
		OnChild:       opts.OnChild,
		ParentFilters: opts.ChildFilters,
		ChildFilters:  opts.ChildFilters,
	})

}

type newRowOpts struct {
	Source string         `json:"source,omitempty"`
	Group  string         `json:"group,omitempty"`
	Table  string         `json:"table,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
	UserId string         `json:"user_id,omitempty"`
}

func (ctx *goja2db) newRow(opts newRowOpts) (any, any) {

	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.NewRow(0, dyndb.NewRowReq{
		TenantId: ctx.tenantId,
		Group:    opts.Group,
		Table:    opts.Table,
		Data:     opts.Data,
		ModCtx: dyndb.ModCtx{
			UserId: opts.UserId,
		},
	})
}

type getRowOpts struct {
	Source    string `json:"source,omitempty"`
	Group     string `json:"group,omitempty"`
	Table     string `json:"table,omitempty"`
	Id        int64  `json:"id,omitempty"`
	SkipCache bool   `json:"skip_cache,omitempty"`
}

func (ctx *goja2db) getRow(opts getRowOpts) (any, any) {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.GetRow(0, dyndb.GetRowReq{
		TenantId:  ctx.tenantId,
		Group:     opts.Group,
		Table:     opts.Table,
		Id:        opts.Id,
		SkipCache: opts.SkipCache,
	})
}

type updateRowOpts struct {
	Id      int64          `json:"id,omitempty"`
	Source  string         `json:"source,omitempty"`
	Version int64          `json:"version,omitempty"`
	Group   string         `json:"group,omitempty"`
	Table   string         `json:"table,omitempty"`
	Data    map[string]any `json:"data,omitempty"`
	UserId  string         `json:"user_id,omitempty"`
}

func (ctx *goja2db) updateRow(opts updateRowOpts) (any, any) {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.UpdateRow(0, dyndb.UpdateRowReq{
		TenantId: ctx.tenantId,
		Id:       opts.Id,
		Version:  opts.Version,
		Group:    opts.Group,
		Table:    opts.Table,
		Data:     opts.Data,
		ModCtx: dyndb.ModCtx{
			UserId: opts.UserId,
		},
	})
}

type deleteRowBatchOpts struct {
	Group       string             `json:"group,omitempty"`
	Source      string             `json:"source,omitempty"`
	Table       string             `json:"table,omitempty"`
	FilterConds []dyndb.FilterCond `json:"filter_conds,omitempty"`
	UserId      string             `json:"user_id,omitempty"`
}

func (ctx *goja2db) deleteRowBatch(opts deleteRowBatchOpts) ([]int64, any) {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.DeleteRowBatch(0, dyndb.DeleteRowBatchReq{
		TenantId:    ctx.tenantId,
		Group:       opts.Group,
		Table:       opts.Table,
		FilterConds: opts.FilterConds,
		ModCtx: dyndb.ModCtx{
			UserId: opts.UserId,
		},
	})
}

type deleteRowMultiOpts struct {
	Source string  `json:"source,omitempty"`
	Group  string  `json:"group,omitempty"`
	Table  string  `json:"table,omitempty"`
	Ids    []int64 `json:"sid,omitempty"`
	UserId string  `json:"user_id,omitempty"`
}

func (ctx *goja2db) deleteRowMulti(opts deleteRowMultiOpts) any {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	table.DeleteRowMulti(0, dyndb.DeleteRowMultiReq{
		TenantId: ctx.tenantId,
		Group:    opts.Group,
		Table:    opts.Group,
		Ids:      opts.Ids,
		ModCtx: dyndb.ModCtx{
			UserId: opts.UserId,
		},
	})

	return nil
}

type deleteRowOpts struct {
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
	Table  string `json:"table,omitempty"`
	Id     int64  `json:"id,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

func (ctx *goja2db) deleteRow(opts deleteRowOpts) any {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.DeleteRow(0, dyndb.DeleteRowReq{
		TenantId: ctx.tenantId,
		Group:    opts.Group,
		Table:    opts.Table,
		Id:       opts.Id,
		ModCtx: dyndb.ModCtx{
			UserId: opts.UserId,
		},
	})
}

type loadTableOpts struct {
	Source      string             `json:"source,omitempty"`
	Group       string             `json:"group,omitempty"`
	Table       string             `json:"table,omitempty"`
	View        string             `json:"view,omitempty"`
	ViewFilters []dyndb.FilterCond `json:"view_filters,omitempty"`
}

func (ctx *goja2db) loadTable(opts loadTableOpts) (any, any) {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.LoadTable(0, dyndb.LoadTableReq{
		TenantId:    ctx.tenantId,
		Table:       opts.Table,
		Group:       opts.Group,
		View:        opts.View,
		ViewFilters: opts.ViewFilters,
	})
}

type ftsQueryOpts struct {
	Source       string             `json:"source,omitempty"`
	Group        string             `json:"group,omitempty"`
	Table        string             `json:"table,omitempty"`
	SearchTerm   string             `json:"search_term,omitempty"`
	SearchColumn string             `json:"search_column,omitempty"`
	Count        int64              `json:"count,omitempty"`
	Filters      []dyndb.FilterCond `json:"filters,omitempty"`
	Page         int64              `json:"page,omitempty"`
	UsePattern   bool               `json:"use_pattern,omitempty"`
}

func (ctx *goja2db) ftsQuery(opts ftsQueryOpts) (any, any) {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.FTSQuery(0, dyndb.FTSQueryReq{
		TenantId:     ctx.tenantId,
		Table:        opts.Table,
		Group:        opts.Group,
		SearchTerm:   opts.SearchTerm,
		SearchColumn: opts.SearchColumn,
		Count:        opts.Count,
		Filters:      opts.Filters,
		Page:         opts.Page,
		UsePattern:   opts.UsePattern,
	})
}

type refResolveOpts struct {
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
	Table  string `json:"table,omitempty"`
	Column string `json:"column,omitempty"`
	Type   string `json:"type,omitempty"`
	Target string `json:"target,omitempty"`
	Object string `json:"object,omitempty"`
	RowIds []any  `json:"row_ids,omitempty"`
}

func (ctx *goja2db) refResolve(opts refResolveOpts) (any, any) {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.RefResolve(0, opts.Group, &dyndb.RefResolveReq{
		Column: opts.Column,
		Type:   opts.Type,
		Target: opts.Target,
		Object: opts.Object,
		RowIds: opts.RowIds,
	})
}

type refLoadOpts struct {
	Source      string `json:"source,omitempty"`
	Group       string `json:"group,omitempty"`
	Table       string `json:"table,omitempty"`
	Column      string `json:"column,omitempty"`
	Type        string `json:"type,omitempty"`
	Target      string `json:"target,omitempty"`
	Object      string `json:"object,omitempty"`
	CursorRowId int64  `json:"cursor_row_id,omitempty"`
}

func (ctx *goja2db) refLoad(opts refLoadOpts) (any, any) {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.RefLoad(0, opts.Group, &dyndb.RefLoadReq{
		Column:      opts.Column,
		Type:        opts.Type,
		Target:      opts.Target,
		Object:      opts.Object,
		CursorRowId: opts.CursorRowId,
	})
}

type reverseRefLoadOpts struct {
	Source       string `json:"source,omitempty"`
	Group        string `json:"group,omitempty"`
	Table        string `json:"table,omitempty"`
	CurrentTable string `json:"current_table,omitempty"`
	TargetTable  string `json:"target_table,omitempty"`
	TargetColumn string `json:"column,omitempty"`
	CurrentItem  any    `json:"current_item,omitempty"`
	CursorRowId  int64  `json:"cursor_row_id,omitempty"`
	Count        int    `json:"count,omitempty"`
}

func (ctx *goja2db) reverseRefLoad(opts reverseRefLoadOpts) (any, any) {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.ReverseRefLoad(0, opts.Group, &dyndb.RevRefLoadReq{
		CurrentTable: opts.CurrentTable,
		TargetTable:  opts.TargetTable,
		TargetColumn: opts.TargetColumn,
		CurrentItem:  opts.CurrentItem,
		CursorRowId:  opts.CursorRowId,
		Count:        opts.Count,
	})
}

type sqlQueryOpts struct {
	Source      string `json:"source,omitempty"`
	Group       string `json:"group,omitempty"`
	NoTransform bool   `json:"no_transform,omitempty"`
	Raw         bool   `json:"raw,omitempty"`
	QStr        string `json:"qstr,omitempty"`
}

func (ctx *goja2db) sqlQuery(opts sqlQueryOpts) (any, any) {
	table := ctx.datahub.GetDataTableHub(opts.Source, ctx.tenantId, opts.Group)

	return table.SqlQuery(0, dyndb.SqlQueryReq{
		NoTransform: true,
		Raw:         opts.Raw,
		Group:       opts.Group,
		QStr:        opts.QStr,
	})
}
