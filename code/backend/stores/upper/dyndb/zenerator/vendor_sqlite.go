package zenerator

import (
	"fmt"
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

var (
	sqliteCtypeMap map[string]string = map[string]string{
		dyndb.CtypeShortText:   "text",
		dyndb.CtypePhone:       "text",
		dyndb.CtypeSelect:      "text",
		dyndb.CtypeRFormula:    "text",
		dyndb.CtypeMultiFile:   "text",
		dyndb.CtypeFile:        "text",
		dyndb.CtypeCheckBox:    "boolean",
		dyndb.CtypeCurrency:    "real",
		dyndb.CtypeNumber:      "integer",
		dyndb.CtypeLocation:    "text",
		dyndb.CtypeDateTime:    "text",
		dyndb.CtypeMultSelect:  "text",
		dyndb.CtypeLongText:    "text",
		dyndb.CtypeSingleUser:  "text",
		dyndb.CtypeMultiUser:   "text",
		dyndb.CtypeEmail:       "text",
		dyndb.CtypeJSON:        "json",
		dyndb.CtypeRangeNumber: "integer",
		dyndb.CtypeColor:       "text",
	}
)

func TableHeadSqlite(tableName string) string {
	var buf strings.Builder
	buf.WriteString("create table ")
	buf.WriteString(tableName)
	buf.WriteString("(\n")
	buf.WriteString("\t__id integer primary key autoincrement not null,\n\t__version integer not null default 0, \n\t__mod_sig text")
	return buf.String()
}

func (g *zenerator) indexSqlite(tblname, iname, itype string, spans []string) string {

	switch itype {
	case "normal":
		return addNormalIndex(tblname, iname, itype, spans)
	default:
		panic("not supported index type:" + itype)
	}
}

func (t *tzz) activityTableSqlite(wctx *WriterCtx) {

	activityTable := t.tns.ActivityTable(t.tenantId, t.gslug, t.model.Slug)

	wctx.Write(
		fmt.Sprintf(`

			CREATE TABLE %s (
				id integer primary key autoincrement not null,
				type text not null,
				row_id integer not null,
				row_version integer not null,
				user_id text not null DEFAULT '',
				user_sign text not null DEFAULT '',
				init_sign text not null DEFAULT '',
				payload text not null DEFAULT '',
				message text not null DEFAULT '',
				created_at timestamptz not null default current_timestamp
			);

			CREATE TRIGGER data_tg_%s_insert
				AFTER INSERT ON %s
			BEGIN
				INSERT INTO %s(
					type,
					row_id,
					row_version,
					user_id,
					user_sign,
					init_sign,
					payload
				)
				VALUES (
					'insert', 
					NEW.__id, 
					NEW.__version, 
					COALESCE(json_extract(NEW.__mod_sig, '$.user_id' ), ''), 
					COALESCE(json_extract(NEW.__mod_sig, '$.user_sign'), ''), 
					COALESCE(json_extract(NEW.__mod_sig, '$.init_sign'), ''),
					json_object(
						`, activityTable, t.tableSlug, t.tableSlug, activityTable,
		))

	for cidx, col := range t.model.Columns {
		if cidx != 0 {
			wctx.Seperator()
		}
		wctx.Write(fmt.Sprintf("'%s', NEW.%s \n", col.Slug, col.Slug))
	}

	wctx.Write(`)
					);					
			END;`)

	wctx.Write(fmt.Sprintf(`			CREATE TRIGGER data_tg_%s_update
			AFTER UPDATE ON %s
		BEGIN
			INSERT INTO %s(
				type,
				row_id,
				row_version,
				user_id,
				user_sign,
				init_sign,
				payload
			)
			VALUES (
				'update', 
				NEW.__id, 
				NEW.__version, 
				COALESCE(json_extract(NEW.__mod_sig, '$.user_id' ), ''), 
				COALESCE(json_extract(NEW.__mod_sig, '$.user_sign'), ''), 
				COALESCE(json_extract(NEW.__mod_sig, '$.init_sign'), ''),
				json_object(
					`, activityTable, t.tableSlug, t.tableSlug))
	for cidx, col := range t.model.Columns {
		if cidx != 0 {
			wctx.Seperator()
		}
		wctx.Write(fmt.Sprintf("'%s', NEW.%s \n", col.Slug, col.Slug))
	}

	wctx.Write(`)
					);					
			END;`)

}
