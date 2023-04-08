package zenerator

import (
	"fmt"

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
			
			`, activityTable,
		),
	)

	// fixme => add trigger

	// if t.model.ActivityType == dyndb.DynActivityTypeStrict {
	// 	wctx.Write(fmt.Sprintf(`CREATE TRIGGER
	// 		data_tg_%s AFTER INSERT OR UPDATE ON
	// 		%s FOR EACH ROW EXECUTE
	// 		FUNCTION data_activity_tg();`, t.tableSlug, t.tableSlug))
	// }

}
