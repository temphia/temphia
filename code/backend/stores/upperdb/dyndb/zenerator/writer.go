package zenerator

import "strings"

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
	if w.terminated {
		return
	}

	w.buffer.Write([]byte(");\n"))
}
