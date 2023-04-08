package zenerator

import (
	"strings"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

var (
	pgCtypeMap map[string]string = map[string]string{
		dyndb.CtypeShortText:   "text",
		dyndb.CtypePhone:       "text",
		dyndb.CtypeSelect:      "text",
		dyndb.CtypeRFormula:    "text",
		dyndb.CtypeMultiFile:   "text",
		dyndb.CtypeFile:        "text",
		dyndb.CtypeCheckBox:    "boolean",
		dyndb.CtypeCurrency:    "decimal",
		dyndb.CtypeNumber:      "integer",
		dyndb.CtypeLocation:    "geography(point,4326)",
		dyndb.CtypeDateTime:    "timestamptz",
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

func PgFTSIndex(table string, model *entities.Index) string {
	var buf strings.Builder
	buf.Write([]byte("CREATE INDEX "))
	buf.WriteString(IndexName(table, model.Slug))
	buf.Write([]byte(" ON "))
	buf.WriteString(table)
	buf.Write([]byte(" USING gin"))
	buf.Write([]byte(Bracketed(model.Spans, " gin_trgm_ops")))
	buf.Write([]byte(";"))
	return buf.String()
}

func (g *zenerator) innerColumnPg(cmap map[string]string) func(cslug, ctype string, notnull bool, defval string) string {
	fn := CTypeMap(cmap)

	return func(cslug, ctype string, notnull bool, defval string) string {

		if ctype == dyndb.CtypeDateTime && defval == "now" {
			defval = " default (now() at time zone 'utc')"
		}
		return fn(cslug, ctype, notnull, defval)
	}

}

func (g *zenerator) indexPg(tblname, iname, itype string, spans []string) string {

	switch itype {
	case "normal":
		return addNormalIndex(tblname, iname, itype, spans)
	case "fts":
		return PgFTSIndex(tblname, &entities.Index{
			Mtype: itype,
			Slug:  iname,
			Spans: spans,
		})
	default:
		panic("not supported index type:" + itype)
	}
}
