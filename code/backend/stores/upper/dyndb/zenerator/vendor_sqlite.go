package zenerator

import (
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
