package goja2db

import (
	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type goja2db struct {
	tenantId  string
	datahub   dyndb.DataHub
	jsruntime *goja.Runtime
}

func New(tenantId string,
	datahub dyndb.DataHub,
	jsruntime *goja.Runtime) *goja2db {
	return &goja2db{
		tenantId:  tenantId,
		datahub:   datahub,
		jsruntime: jsruntime,
	}
}

// fixme => remove extra struct and go back to original interface may be with invoker token

func (ctx *goja2db) Bind() {
	ctx.bindSheet()
	ctx.bindTable()
}

func (ctx *goja2db) bindTable() {
	ctx.jsruntime.Set("table_query", ctx.tableQuery)
	ctx.jsruntime.Set("table_join_query", ctx.tableJoinQuery)
	ctx.jsruntime.Set("new_row", ctx.newRow)
	ctx.jsruntime.Set("get_row", ctx.getRow)
	ctx.jsruntime.Set("update_row", ctx.updateRow)
	ctx.jsruntime.Set("delete_row_batch", ctx.deleteRowBatch)
	ctx.jsruntime.Set("delete_row_multi", ctx.deleteRowMulti)
	ctx.jsruntime.Set("delete_row", ctx.deleteRow)
	ctx.jsruntime.Set("load_table", ctx.loadTable)
	ctx.jsruntime.Set("fts_query", ctx.ftsQuery)
	ctx.jsruntime.Set("ref_resolve", ctx.refResolve)
	ctx.jsruntime.Set("ref_load", ctx.refLoad)
	ctx.jsruntime.Set("reverse_ref_load", ctx.reverseRefLoad)
	ctx.jsruntime.Set("sql_query", ctx.sqlQuery)

}

func (ctx *goja2db) bindSheet() {
	ctx.jsruntime.Set("list_sheet_group", ctx.listSheetGroup)
	ctx.jsruntime.Set("list_sheet", ctx.listSheet)
	ctx.jsruntime.Set("new_sheet", ctx.newSheet)
	ctx.jsruntime.Set("get_sheet", ctx.getSheet)
	ctx.jsruntime.Set("update_sheet", ctx.updateSheet)
	ctx.jsruntime.Set("delete_sheet", ctx.deleteSheet)
	ctx.jsruntime.Set("list_sheet_column", ctx.listSheetColumn)
	ctx.jsruntime.Set("new_sheet_column", ctx.newSheetColumn)
	ctx.jsruntime.Set("get_sheet_column", ctx.getSheetColumn)
	ctx.jsruntime.Set("update_sheet_column", ctx.updateSheetColumn)
	ctx.jsruntime.Set("delete_sheet_column", ctx.deleteSheetColumn)
	ctx.jsruntime.Set("load_sheet", ctx.loadSheet)
	ctx.jsruntime.Set("query_sheet", ctx.querySheet)
	ctx.jsruntime.Set("fts_query_sheet", ctx.ftsQuerySheet)
	ctx.jsruntime.Set("new_row_with_cell", ctx.newRowWithCell)
	ctx.jsruntime.Set("update_row_with_cell", ctx.updateRowWithCell)
	ctx.jsruntime.Set("delete_row_with_cell", ctx.deleteRowWithCell)
	ctx.jsruntime.Set("get_row_relations", ctx.getRowRelations)
}
