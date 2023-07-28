package datagroup

import "github.com/temphia/temphia/code/backend/xtypes/store/dyndb"

type sheetOpts struct {
	TxId uint32 `json:"txid,omitempty"`
}

func (ctx *DatagroupModule) ListSheetGroup(opts *sheetOpts) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.ListSheetGroup(opts.TxId)
}

func (ctx *DatagroupModule) ListSheet(opts *sheetOpts) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.ListSheet(opts.TxId)
}

type newSheetOpts struct {
	TxId   uint32         `json:"txid,omitempty"`
	UserId string         `json:"user_id,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
}

func (ctx *DatagroupModule) NewSheet(opts *newSheetOpts) any {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)

	return sheet.NewSheet(opts.TxId, opts.UserId, opts.Data)
}

type getSheetOpts struct {
	TxId uint32 `json:"txid,omitempty"`
	Id   int64  `json:"id,omitempty"`
}

func (ctx *DatagroupModule) GetSheet(opts *getSheetOpts) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.GetSheet(opts.TxId, opts.Id)
}

type updateSheetOpts struct {
	TxId   uint32         `json:"txid,omitempty"`
	UserId string         `json:"user_id,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
	Id     int64          `json:"id,omitempty"`
}

func (ctx *DatagroupModule) UpdateSheet(opts *updateSheetOpts) any {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.UpdateSheet(opts.TxId, opts.Id, opts.UserId, opts.Data)
}

type deleteSheetOpts struct {
	TxId   uint32 `json:"txid,omitempty"`
	UserId string `json:"user_id,omitempty"`
	Id     int64  `json:"id,omitempty"`
}

func (ctx *DatagroupModule) DeleteSheet(opts *deleteSheetOpts) any {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.DeleteSheet(opts.TxId, opts.Id, opts.UserId)
}

type listSheetColumnOpts struct {
	TxId   uint32 `json:"txid,omitempty"`
	UserId string `json:"user_id,omitempty"`
	Sid    int64  `json:"sid,omitempty"`
}

func (ctx *DatagroupModule) ListSheetColumn(opts *listSheetColumnOpts) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.ListSheetColumn(opts.TxId, opts.Sid)
}

type newSheetColumnOpts struct {
	TxId   uint32         `json:"txid,omitempty"`
	UserId string         `json:"user_id,omitempty"`
	Sid    int64          `json:"sid,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
}

func (ctx *DatagroupModule) NewSheetColumn(opts *newSheetColumnOpts) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.NewSheetColumn(opts.TxId, opts.Sid, opts.UserId, opts.Data)
}

type getSheetColumnOpts struct {
	TxId   uint32 `json:"txid,omitempty"`
	UserId string `json:"user_id,omitempty"`
	Sid    int64  `json:"sid,omitempty"`
	Cid    int64  `json:"cid,omitempty"`
}

func (ctx *DatagroupModule) GetSheetColumn(opts *getSheetColumnOpts) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.GetSheetColumn(opts.TxId, opts.Sid, opts.Cid)
}

type updateSheetColumnOpts struct {
	TxId   uint32         `json:"txid,omitempty"`
	UserId string         `json:"user_id,omitempty"`
	Sid    int64          `json:"sid,omitempty"`
	Cid    int64          `json:"cid,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
}

func (ctx *DatagroupModule) UpdateSheetColumn(opts *updateSheetColumnOpts) any {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.UpdateSheetColumn(opts.TxId, opts.Sid, opts.Cid, opts.UserId, opts.Data)
}

type deleteSheetColumnOpts struct {
	TxId   uint32 `json:"txid,omitempty"`
	UserId string `json:"user_id,omitempty"`
	Sid    int64  `json:"sid,omitempty"`
	Cid    int64  `json:"cid,omitempty"`
}

func (ctx *DatagroupModule) DeleteSheetColumn(opts *deleteSheetColumnOpts) any {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.DeleteSheetColumn(opts.TxId, opts.Sid, opts.Cid, opts.UserId)
}

type loadSheetOpts struct {
	TxId        uint32 `json:"txid,omitempty"`
	SheetId     int64  `json:"sheet_id,omitempty"`
	View        string `json:"view,omitempty"`
	RowCursorId int64  `json:"row_cursor_id,omitempty"`
	Desc        bool   `json:"desc,omitempty"`
}

func (ctx *DatagroupModule) LoadSheet(opts *loadSheetOpts) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)

	return sheet.LoadSheet(opts.TxId, &dyndb.LoadSheetReq{
		TenantId:    ctx.tenantId,
		Group:       ctx.group,
		SheetId:     opts.SheetId,
		View:        opts.View,
		RowCursorId: opts.RowCursorId,
		Desc:        opts.Desc,
	})
}

type sheetQueryOptions struct {
	TxId        uint32 `json:"txid,omitempty"`
	SheetId     int64  `json:"sheet_id,omitempty"`
	RowCursorId int64  `json:"row_cursor_id,omitempty"`
	Desc        bool   `json:"desc,omitempty"`
}

func (ctx *DatagroupModule) QuerySheet(opts *sheetQueryOptions) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)

	return sheet.Query(opts.TxId, &dyndb.QuerySheetReq{
		TenantId: ctx.tenantId,
		Group:    ctx.group,
		SheetId:  opts.SheetId,
		Desc:     opts.Desc,
	})
}

type ftsOpts struct {
	TxId       uint32 `json:"txid,omitempty"`
	SheetId    int64  `json:"sheet_id,omitempty"`
	SearchTerm string `json:"search_term,omitempty"`
}

func (ctx *DatagroupModule) FtsQuerySheet(opts *ftsOpts) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)

	return sheet.FTSQuery(opts.TxId, &dyndb.FTSQuerySheet{
		TenantId:   ctx.tenantId,
		Group:      ctx.group,
		SheetId:    opts.SheetId,
		SearchTerm: opts.SearchTerm,
	})
}

type newRowWithCellOpts struct {
	TxId    uint32                   `json:"txid,omitempty"`
	SheetId int64                    `json:"sheet_id,omitempty"`
	UserId  string                   `json:"user_id,omitempty"`
	Data    map[int64]map[string]any `json:"data,omitempty"`
}

func (ctx *DatagroupModule) NewRowWithCell(opts *newRowWithCellOpts) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.NewRowWithCell(opts.TxId, opts.SheetId, opts.UserId, opts.Data)
}

type updateRowWithCellOpts struct {
	TxId    uint32                   `json:"txid,omitempty"`
	SheetId int64                    `json:"sheet_id,omitempty"`
	RowId   int64                    `json:"row_id,omitempty"`
	UserId  string                   `json:"user_id,omitempty"`
	Data    map[int64]map[string]any `json:"data,omitempty"`
}

func (ctx *DatagroupModule) UpdateRowWithCell(opts *updateRowWithCellOpts) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.UpdateRowWithCell(opts.TxId, opts.SheetId, opts.RowId, opts.UserId, opts.Data)
}

type deleteRowWithCellOpts struct {
	TxId    uint32 `json:"txid,omitempty"`
	SheetId int64  `json:"sheet_id,omitempty"`
	RowId   int64  `json:"row_id,omitempty"`
	UserId  string `json:"user_id,omitempty"`
}

func (ctx *DatagroupModule) DeleteRowWithCell(opts *deleteRowWithCellOpts) any {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.DeleteRowWithCell(opts.TxId, opts.SheetId, opts.RowId, opts.UserId)
}

type getRowRelationsOpts struct {
	TxId     uint32 `json:"txid,omitempty"`
	UserId   string `json:"user_id,omitempty"`
	SheetId  int64  `json:"sheet_id,omitempty"`
	RowId    int64  `json:"row_id,omitempty"`
	RefSheet int64  `json:"ref_sheet,omitempty"`
	RefCol   int64  `json:"ref_col,omitempty"`
}

func (ctx *DatagroupModule) GetRowRelations(opts *getRowRelationsOpts) (any, any) {
	sheet := ctx.dynsrc.GetDataSheetHub(ctx.tenantId, ctx.group)
	return sheet.GetRowRelations(opts.TxId, opts.SheetId, opts.RowId, opts.RefSheet, opts.RefCol)
}
