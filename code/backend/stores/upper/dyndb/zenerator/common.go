package zenerator

import (
	"errors"
	"strings"
)

var (
	ErrUnknownColumn = errors.New("Unknown Column")
	ErrUnknownVendor = errors.New("Unknown Vendor")
)

func TableHead(tableName string) string {
	var buf strings.Builder
	buf.WriteString("create table ")
	buf.WriteString(tableName)
	buf.WriteString("(\n")
	buf.WriteString("\t__id serial primary key,\n\t__version integer not null default 0, \n\t__mod_sig text")
	return buf.String()
}

func Bracketed(items []string, class string) string {
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

type WriterCtx struct {
	buffer     strings.Builder
	seperated  bool // seperated with comma
	terminated bool // terminated with ';'
}

func (w *WriterCtx) DirectWrite(s string) {
	w.buffer.WriteString(s)
}

func (w *WriterCtx) CondWriteCol(write bool, s string) {
	if !write {
		return
	}
	w.Seperator()
	w.Write(s)
}

func (w *WriterCtx) Write(s string) {
	w.seperated = false
	w.buffer.WriteString(s)
}

func (w *WriterCtx) Seperator() {
	if w.seperated {
		return
	}
	w.buffer.WriteByte(byte(','))
}

func (w *WriterCtx) Terminate() {
	w.buffer.Write([]byte(");\n"))
}

func Unique(spans []string) string {
	var buf strings.Builder
	buf.Write([]byte("UNIQUE"))
	buf.Write([]byte(Bracketed(spans, "")))
	return buf.String()
}

func CTypeMap(mapping map[string]string) func(slug, ctype string, notnull bool, defval string) string {
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

func InnerFKRef(target string, from []string, to []string) string {
	var buf strings.Builder
	buf.Write([]byte("FOREIGN KEY"))
	buf.Write([]byte(Bracketed(from, "")))
	buf.Write([]byte(" REFERENCES "))
	buf.WriteString(target)
	buf.WriteString(Bracketed(to, ""))
	//buf.Write([]byte(";"))
	return buf.String()
}

func IndexName(tblname, iname string) string {
	return tblname + "_" + iname
}

func addNormalIndex(tbl string, iname, itype string, spans []string) string {
	var buf strings.Builder
	buf.Write([]byte("CREATE INDEX "))
	buf.WriteString(IndexName(tbl, iname))
	buf.Write([]byte(" ON "))
	buf.WriteString(tbl)
	buf.Write([]byte(Bracketed(spans, "")))
	buf.Write([]byte(";"))
	return buf.String()
}
