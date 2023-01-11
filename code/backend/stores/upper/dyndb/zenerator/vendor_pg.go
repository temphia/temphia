package zenerator

import (
	"strings"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

var (
	pgCtypeMap map[string]string = map[string]string{
		store.CtypeShortText:   "text",
		store.CtypePhone:       "text",
		store.CtypeSelect:      "text",
		store.CtypeRFormula:    "text",
		store.CtypeMultiFile:   "text",
		store.CtypeFile:        "text",
		store.CtypeCheckBox:    "boolean",
		store.CtypeCurrency:    "decimal",
		store.CtypeNumber:      "integer",
		store.CtypeLocation:    "geography(point,4326)",
		store.CtypeDateTime:    "timestamptz",
		store.CtypeMultSelect:  "text",
		store.CtypeLongText:    "text",
		store.CtypeSingleUser:  "text",
		store.CtypeMultiUser:   "text",
		store.CtypeEmail:       "text",
		store.CtypeJSON:        "json",
		store.CtypeRangeNumber: "integer",
		store.CtypeColor:       "text",
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

func (g *zenerator) innerColumnPg() func(cslug, ctype string, notnull bool, defval string) string {
	fn := CTypeMap(pgCtypeMap)

	return func(cslug, ctype string, notnull bool, defval string) string {

		if ctype == store.CtypeDateTime && defval == "now" {
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
