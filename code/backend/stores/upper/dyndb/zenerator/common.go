package zenerator

import (
	"strings"
)

func tableHead(tableName string) string {
	var buf strings.Builder
	buf.WriteString("create table ")
	buf.WriteString(tableName)
	buf.WriteString("(\n")
	buf.WriteString("\t__id serial primary key,\n\t__version integer not null default 0, \n\t__mod_sig text")
	return buf.String()
}

func bracketed(items []string, class string) string {
	var buf strings.Builder
	buf.Write([]byte("("))

	for i, item := range items {
		if i != 0 {
			buf.Write([]byte(","))
		}
		buf.Write([]byte(item))
		buf.Write([]byte(class))
	}
	buf.Write([]byte(")"))

	return buf.String()
}

func unique(spans []string) string {
	var buf strings.Builder
	buf.Write([]byte("UNIQUE"))
	buf.Write([]byte(bracketed(spans, "")))
	return buf.String()
}

func cTypeMap(mapping map[string]string) func(slug, ctype string, notnull bool, defval string) string {
	return func(slug, ctype string, notnull bool, defval string) string {
		var buf strings.Builder

		buf.WriteString("\n\t")
		buf.WriteString(slug)
		buf.WriteString(" ")
		buf.WriteString(mapping[ctype])

		if notnull {
			buf.WriteString(" not null")
		}

		buf.WriteString(defval)
		buf.WriteString("\n")
		return buf.String()
	}
}

func innerFKRef(target string, from []string, to []string) string {
	var buf strings.Builder
	buf.Write([]byte("FOREIGN KEY"))
	buf.Write([]byte(bracketed(from, "")))
	buf.Write([]byte(" REFERENCES "))
	buf.WriteString(target)
	buf.WriteString(bracketed(to, ""))
	//buf.Write([]byte(";"))
	return buf.String()
}

func indexName(tblname, iname string) string {
	return tblname + "_" + iname
}

func addNormalIndex(tbl string, iname, itype string, spans []string) string {
	var buf strings.Builder
	buf.Write([]byte("CREATE INDEX "))
	buf.WriteString(indexName(tbl, iname))
	buf.Write([]byte(" ON "))
	buf.WriteString(tbl)
	buf.Write([]byte(bracketed(spans, "")))
	buf.Write([]byte(";"))
	return buf.String()
}
