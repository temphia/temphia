package wizard

import (
	"reflect"
	"strings"

	"github.com/temphia/temphia/code/executors/backend/execlib/scripter"
)

func (sw *SimpleWizard) execScript(name string, ctxData interface{}, binds map[string]interface{}) error {

	binds["current_scope_methods"] = func() map[string]string {
		scopeMethods := make(map[string]string, len(binds)-1)
		for fkey, fn := range binds {
			scopeMethods[fkey] = signature(fn)
		}
		return scopeMethods
	}

	script := scripter.New(sw.jsRuntime, binds, true)

	script.Bind()
	defer script.Clear()

	var entry func(interface{}) error
	err := script.SetEntry(name, entry)
	if err != nil {
		return err
	}

	return entry(ctxData)
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
