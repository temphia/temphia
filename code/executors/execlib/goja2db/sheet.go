package goja2db

import "github.com/temphia/temphia/code/backend/xtypes/store/dyndb"

type sheetOpts struct {
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
}

func (ctx *goja2db) listSheetGroup(opts sheetOpts) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.ListSheetGroup(0)
}

func (ctx *goja2db) listSheet(opts sheetOpts) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.ListSheet(0)
}

type newSheetOpts struct {
	Source string         `json:"source,omitempty"`
	Group  string         `json:"group,omitempty"`
	UserId string         `json:"user_id,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
}

func (ctx *goja2db) newSheet(opts newSheetOpts) any {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.NewSheet(0, opts.UserId, opts.Data)
}

type getSheetOpts struct {
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
	Id     int64  `json:"id,omitempty"`
}

func (ctx *goja2db) getSheet(opts getSheetOpts) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.GetSheet(0, opts.Id)
}

type updateSheetOpts struct {
	Source string         `json:"source,omitempty"`
	Group  string         `json:"group,omitempty"`
	UserId string         `json:"user_id,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
	Id     int64          `json:"id,omitempty"`
}

func (ctx *goja2db) updateSheet(opts updateSheetOpts) any {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.UpdateSheet(0, opts.Id, opts.UserId, opts.Data)
}

type deleteSheetOpts struct {
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
	UserId string `json:"user_id,omitempty"`
	Id     int64  `json:"id,omitempty"`
}

func (ctx *goja2db) deleteSheet(id int64, opts deleteSheetOpts) any {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.DeleteSheet(0, opts.Id, opts.UserId)
}

type listSheetColumnOpts struct {
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
	UserId string `json:"user_id,omitempty"`
	Sid    int64  `json:"sid,omitempty"`
}

func (ctx *goja2db) listSheetColumn(opts listSheetColumnOpts) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.ListSheetColumn(0, opts.Sid)
}

type newSheetColumnOpts struct {
	Source string         `json:"source,omitempty"`
	Group  string         `json:"group,omitempty"`
	UserId string         `json:"user_id,omitempty"`
	Sid    int64          `json:"sid,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
}

func (ctx *goja2db) newSheetColumn(opts newSheetColumnOpts) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.NewSheetColumn(0, opts.Sid, opts.UserId, opts.Data)
}

type getSheetColumnOpts struct {
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
	UserId string `json:"user_id,omitempty"`
	Sid    int64  `json:"sid,omitempty"`
	Cid    int64  `json:"cid,omitempty"`
}

func (ctx *goja2db) getSheetColumn(opts getSheetColumnOpts) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.GetSheetColumn(0, opts.Sid, opts.Cid)
}

type updateSheetColumnOpts struct {
	Source string         `json:"source,omitempty"`
	Group  string         `json:"group,omitempty"`
	UserId string         `json:"user_id,omitempty"`
	Sid    int64          `json:"sid,omitempty"`
	Cid    int64          `json:"cid,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
}

func (ctx *goja2db) updateSheetColumn(opts updateSheetColumnOpts) any {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.UpdateSheetColumn(0, opts.Sid, opts.Cid, opts.UserId, opts.Data)
}

type deleteSheetColumnOpts struct {
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
	UserId string `json:"user_id,omitempty"`
	Sid    int64  `json:"sid,omitempty"`
	Cid    int64  `json:"cid,omitempty"`
}

func (ctx *goja2db) deleteSheetColumn(opts deleteSheetColumnOpts) any {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.DeleteSheetColumn(0, opts.Sid, opts.Cid, opts.UserId)
}

type loadSheetOpts struct {
	Source      string `json:"source,omitempty"`
	Group       string `json:"group,omitempty"`
	SheetId     int64  `json:"sheet_id,omitempty"`
	View        string `json:"view,omitempty"`
	RowCursorId int64  `json:"row_cursor_id,omitempty"`
	Desc        bool   `json:"desc,omitempty"`
}

func (ctx *goja2db) loadSheet(opts loadSheetOpts) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)

	return sheet.LoadSheet(0, &dyndb.LoadSheetReq{
		TenantId:    ctx.tenantId,
		Group:       opts.Group,
		SheetId:     opts.SheetId,
		View:        opts.View,
		RowCursorId: opts.RowCursorId,
		Desc:        opts.Desc,
	})
}

type sheetQueryOptions struct {
	Source      string `json:"source,omitempty"`
	Group       string `json:"group,omitempty"`
	SheetId     int64  `json:"sheet_id,omitempty"`
	RowCursorId int64  `json:"row_cursor_id,omitempty"`
	Desc        bool   `json:"desc,omitempty"`
}

func (ctx *goja2db) querySheet(opts sheetQueryOptions) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)

	return sheet.Query(0, &dyndb.QuerySheetReq{
		TenantId:    ctx.tenantId,
		Group:       opts.Group,
		SheetId:     opts.SheetId,
		RowCursorId: opts.RowCursorId,
		Desc:        opts.Desc,
	})
}

type ftsOpts struct {
	Source     string `json:"source,omitempty"`
	Group      string `json:"group,omitempty"`
	SheetId    int64  `json:"sheet_id,omitempty"`
	SearchTerm string `json:"search_term,omitempty"`
}

func (ctx *goja2db) ftsQuerySheet(opts ftsOpts) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)

	return sheet.FTSQuery(0, &dyndb.FTSQuerySheet{
		TenantId:   ctx.tenantId,
		Group:      opts.Group,
		SheetId:    opts.SheetId,
		SearchTerm: opts.SearchTerm,
	})
}

type newRowWithCellOpts struct {
	Source  string                   `json:"source,omitempty"`
	Group   string                   `json:"group,omitempty"`
	SheetId int64                    `json:"sheet_id,omitempty"`
	UserId  string                   `json:"user_id,omitempty"`
	Data    map[int64]map[string]any `json:"data,omitempty"`
}

func (ctx *goja2db) newRowWithCell(opts newRowWithCellOpts) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.NewRowWithCell(0, opts.SheetId, opts.UserId, opts.Data)
}

type updateRowWithCellOpts struct {
	Source  string                   `json:"source,omitempty"`
	Group   string                   `json:"group,omitempty"`
	SheetId int64                    `json:"sheet_id,omitempty"`
	RowId   int64                    `json:"row_id,omitempty"`
	UserId  string                   `json:"user_id,omitempty"`
	Data    map[int64]map[string]any `json:"data,omitempty"`
}

func (ctx *goja2db) updateRowWithCell(opts updateRowWithCellOpts) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.UpdateRowWithCell(0, opts.SheetId, opts.RowId, opts.UserId, opts.Data)
}

type deleteRowWithCellOpts struct {
	Source  string `json:"source,omitempty"`
	Group   string `json:"group,omitempty"`
	SheetId int64  `json:"sheet_id,omitempty"`
	RowId   int64  `json:"row_id,omitempty"`
	UserId  string `json:"user_id,omitempty"`
}

func (ctx *goja2db) deleteRowWithCell(opts deleteRowWithCellOpts) any {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.DeleteRowWithCell(0, opts.SheetId, opts.RowId, opts.UserId)
}

type getRowRelationsOpts struct {
	Source   string `json:"source,omitempty"`
	Group    string `json:"group,omitempty"`
	UserId   string `json:"user_id,omitempty"`
	SheetId  int64  `json:"sheet_id,omitempty"`
	RowId    int64  `json:"row_id,omitempty"`
	RefSheet int64  `json:"ref_sheet,omitempty"`
	RefCol   int64  `json:"ref_col,omitempty"`
}

func (ctx *goja2db) getRowRelations(opts getRowRelationsOpts) (any, any) {
	sheet := ctx.datahub.GetDataSheetHub(opts.Source, ctx.tenantId, opts.Group)
	return sheet.GetRowRelations(0, opts.SheetId, opts.RowId, opts.RefSheet, opts.RefCol)
}
