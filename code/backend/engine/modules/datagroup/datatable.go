package datagroup

import (
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

// fixme => uctx := d.binder.InvokerGet().ContextUser()

type tableQueryOptions struct {
	Table       string             `json:"table,omitempty"`
	Count       int64              `json:"count,omitempty"`
	FilterConds []dyndb.FilterCond `json:"filter_conds,omitempty"`
	Page        int64              `json:"page,omitempty"`
	Selects     []string           `json:"selects,omitempty"`
	OrderBy     string             `json:"order_by,omitempty"`
	Desc        bool               `json:"desc,omitempty"`
	TxId        uint32             `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) TableQuery(opts *tableQueryOptions) (any, error) {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.source)

	return table.SimpleQuery(opts.TxId, dyndb.SimpleQueryReq{
		TenantId:    ctx.tenantId,
		Group:       ctx.group,
		Table:       opts.Table,
		Count:       opts.Count,
		FilterConds: opts.FilterConds,
		Page:        opts.Page,
		Selects:     opts.Selects,
		OrderBy:     opts.OrderBy,
		Desc:        opts.Desc,
	})

}

type tablejoinQueryOptions struct {
	Parent        string             `json:"parent,omitempty"`
	Child         string             `json:"child,omitempty"`
	OnParent      string             `json:"on_parent,omitempty"`
	OnChild       string             `json:"on_child,omitempty"`
	ParentFilters []dyndb.FilterCond `json:"parent_ft,omitempty"`
	ChildFilters  []dyndb.FilterCond `json:"child_ft,omitempty"`
	TxId          uint32             `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) TableJoinQuery(opts *tablejoinQueryOptions) (any, error) {

	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.JoinQuery(opts.TxId, dyndb.JoinReq{
		TenantId:      ctx.tenantId,
		Group:         ctx.group,
		Parent:        opts.Parent,
		Child:         opts.Child,
		OnParent:      opts.OnParent,
		OnChild:       opts.OnChild,
		ParentFilters: opts.ChildFilters,
		ChildFilters:  opts.ChildFilters,
	})
}

type newRowOpts struct {
	Table  string         `json:"table,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
	UserId string         `json:"user_id,omitempty"`
	TxId   uint32         `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) NewRow(opts *newRowOpts) (any, error) {

	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.NewRow(opts.TxId, dyndb.NewRowReq{
		TenantId: ctx.tenantId,
		Group:    ctx.group,
		Table:    opts.Table,
		Data:     opts.Data,
		ModCtx: dyndb.ModCtx{
			UserId: opts.UserId,
		},
	})
}

type getRowOpts struct {
	Table     string `json:"table,omitempty"`
	Id        int64  `json:"id,omitempty"`
	SkipCache bool   `json:"skip_cache,omitempty"`
	TxId      uint32 `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) GetRow(opts *getRowOpts) (any, error) {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.GetRow(opts.TxId, dyndb.GetRowReq{
		TenantId:  ctx.tenantId,
		Group:     ctx.group,
		Table:     opts.Table,
		Id:        opts.Id,
		SkipCache: opts.SkipCache,
	})
}

type updateRowOpts struct {
	Id      int64          `json:"id,omitempty"`
	Version int64          `json:"version,omitempty"`
	Table   string         `json:"table,omitempty"`
	Data    map[string]any `json:"data,omitempty"`
	UserId  string         `json:"user_id,omitempty"`
	TxId    uint32         `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) UpdateRow(opts *updateRowOpts) (any, error) {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.UpdateRow(opts.TxId, dyndb.UpdateRowReq{
		TenantId: ctx.tenantId,
		Id:       opts.Id,
		Version:  opts.Version,
		Group:    ctx.group,
		Table:    opts.Table,
		Data:     opts.Data,
		ModCtx: dyndb.ModCtx{
			UserId: opts.UserId,
		},
	})
}

type deleteRowBatchOpts struct {
	Table       string             `json:"table,omitempty"`
	FilterConds []dyndb.FilterCond `json:"filter_conds,omitempty"`
	UserId      string             `json:"user_id,omitempty"`
	TxId        uint32             `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) DeleteRowBatch(opts *deleteRowBatchOpts) ([]int64, any) {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.DeleteRowBatch(opts.TxId, dyndb.DeleteRowBatchReq{
		TenantId:    ctx.tenantId,
		Group:       ctx.group,
		Table:       opts.Table,
		FilterConds: opts.FilterConds,
		ModCtx: dyndb.ModCtx{
			UserId: opts.UserId,
		},
	})
}

type deleteRowMultiOpts struct {
	Table  string  `json:"table,omitempty"`
	Ids    []int64 `json:"sid,omitempty"`
	UserId string  `json:"user_id,omitempty"`
	TxId   uint32  `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) DeleteRowMulti(opts *deleteRowMultiOpts) any {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	table.DeleteRowMulti(opts.TxId, dyndb.DeleteRowMultiReq{
		TenantId: ctx.tenantId,
		Group:    ctx.group,
		Table:    opts.Table,
		Ids:      opts.Ids,
		ModCtx: dyndb.ModCtx{
			UserId: opts.UserId,
		},
	})

	return nil
}

type deleteRowOpts struct {
	Table  string `json:"table,omitempty"`
	Id     int64  `json:"id,omitempty"`
	UserId string `json:"user_id,omitempty"`
	TxId   uint32 `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) DeleteRow(opts *deleteRowOpts) any {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.DeleteRow(opts.TxId, dyndb.DeleteRowReq{
		TenantId: ctx.tenantId,
		Group:    ctx.group,
		Table:    opts.Table,
		Id:       opts.Id,
		ModCtx: dyndb.ModCtx{
			UserId: opts.UserId,
		},
	})
}

type loadTableOpts struct {
	Table       string             `json:"table,omitempty"`
	View        string             `json:"view,omitempty"`
	ViewFilters []dyndb.FilterCond `json:"view_filters,omitempty"`
	TxId        uint32             `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) LoadTable(opts *loadTableOpts) (any, error) {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.LoadTable(opts.TxId, dyndb.LoadTableReq{
		TenantId:    ctx.tenantId,
		Table:       opts.Table,
		Group:       ctx.group,
		View:        opts.View,
		ViewFilters: opts.ViewFilters,
	})
}

type ftsQueryOpts struct {
	Table        string             `json:"table,omitempty"`
	SearchTerm   string             `json:"search_term,omitempty"`
	SearchColumn string             `json:"search_column,omitempty"`
	Count        int64              `json:"count,omitempty"`
	Filters      []dyndb.FilterCond `json:"filters,omitempty"`
	Page         int64              `json:"page,omitempty"`
	UsePattern   bool               `json:"use_pattern,omitempty"`
	TxId         uint32             `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) FtsQuery(opts *ftsQueryOpts) (any, error) {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.FTSQuery(opts.TxId, dyndb.FTSQueryReq{
		TenantId:     ctx.tenantId,
		Table:        opts.Table,
		Group:        ctx.group,
		SearchTerm:   opts.SearchTerm,
		SearchColumn: opts.SearchColumn,
		Count:        opts.Count,
		Filters:      opts.Filters,
		Page:         opts.Page,
		UsePattern:   opts.UsePattern,
	})
}

type refResolveOpts struct {
	Table  string `json:"table,omitempty"`
	Column string `json:"column,omitempty"`
	Type   string `json:"type,omitempty"`
	Target string `json:"target,omitempty"`
	Object string `json:"object,omitempty"`
	RowIds []any  `json:"row_ids,omitempty"`
	TxId   uint32 `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) RefResolve(opts *refResolveOpts) (any, error) {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.RefResolve(opts.TxId, ctx.group, &dyndb.RefResolveReq{
		Column: opts.Column,
		Type:   opts.Type,
		Target: opts.Target,
		Object: opts.Object,
		RowIds: opts.RowIds,
	})
}

type refLoadOpts struct {
	Table       string `json:"table,omitempty"`
	Column      string `json:"column,omitempty"`
	Type        string `json:"type,omitempty"`
	Target      string `json:"target,omitempty"`
	Object      string `json:"object,omitempty"`
	CursorRowId int64  `json:"cursor_row_id,omitempty"`
	TxId        uint32 `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) RefLoad(opts *refLoadOpts) (any, error) {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.RefLoad(opts.TxId, ctx.group, &dyndb.RefLoadReq{
		Column:      opts.Column,
		Type:        opts.Type,
		Target:      opts.Target,
		Object:      opts.Object,
		CursorRowId: opts.CursorRowId,
	})
}

type reverseRefLoadOpts struct {
	Table        string `json:"table,omitempty"`
	CurrentTable string `json:"current_table,omitempty"`
	TargetTable  string `json:"target_table,omitempty"`
	TargetColumn string `json:"column,omitempty"`
	CurrentItem  any    `json:"current_item,omitempty"`
	CursorRowId  int64  `json:"cursor_row_id,omitempty"`
	Count        int    `json:"count,omitempty"`
	TxId         uint32 `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) ReverseRefLoad(opts *reverseRefLoadOpts) (any, error) {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.ReverseRefLoad(opts.TxId, ctx.group, &dyndb.RevRefLoadReq{
		CurrentTable: opts.CurrentTable,
		TargetTable:  opts.TargetTable,
		TargetColumn: opts.TargetColumn,
		CurrentItem:  opts.CurrentItem,
		CursorRowId:  opts.CursorRowId,
		Count:        opts.Count,
	})
}

type sqlQueryOpts struct {
	NoTransform bool   `json:"no_transform,omitempty"`
	Raw         bool   `json:"raw,omitempty"`
	QStr        string `json:"qstr,omitempty"`
	TxId        uint32 `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) SqlQuery(opts *sqlQueryOpts) (any, error) {
	table := ctx.dynsrc.GetDataTableHub(ctx.tenantId, ctx.group)

	return table.SqlQuery(opts.TxId, dyndb.SqlQueryReq{
		NoTransform: true,
		Raw:         opts.Raw,
		Group:       ctx.group,
		QStr:        opts.QStr,
	})
}
