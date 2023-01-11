package scripter

import (
	"reflect"
	"strings"

	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
)

type Scripter struct {
	runtime   *goja.Runtime
	bindFuncs map[string]interface{}
}

func New(rt *goja.Runtime, bfuncs map[string]interface{}, methodHelper bool) *Scripter {

	if methodHelper {
		bfuncs["current_scope_methods"] = func() map[string]string {
			scopeMethods := make(map[string]string, len(bfuncs)-1)
			for fkey, fn := range bfuncs {
				scopeMethods[fkey] = signature(fn)
			}
			return scopeMethods
		}
	}

	return &Scripter{
		runtime:   rt,
		bindFuncs: bfuncs,
	}

}

func (se *Scripter) Bind() {
	for name, fn := range se.bindFuncs {
		se.runtime.Set(name, fn)
	}

}

func (se *Scripter) SetEntry(name string, entry interface{}) error {
	rawentry := se.runtime.Get(name)
	if rawentry == nil {
		return easyerr.NotFound()
	}

	return se.runtime.ExportTo(rawentry, entry)
}

func (se *Scripter) Clear() {
	for name := range se.bindFuncs {
		se.runtime.Set(name, nil) // fixme => probably goja.Undefined() better ?
	}
}

func signature(f interface{}) string {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		return "<not a function>"
	}

	buf := strings.Builder{}
	buf.WriteString("func (")
	for i := 0; i < t.NumIn(); i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(t.In(i).String())
	}
	buf.WriteString(")")
	if numOut := t.NumOut(); numOut > 0 {
		if numOut > 1 {
			buf.WriteString(" (")
		} else {
			buf.WriteString(" ")
		}
		for i := 0; i < t.NumOut(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(t.Out(i).String())
		}
		if numOut > 1 {
			buf.WriteString(")")
		}
	}

	return buf.String()
}
