package pagedash

import (
	"sync"

	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
)

type PdCtx struct {
	Data    map[string]any
	Model   *DashModel
	Message string
	Rt      *goja.Runtime
}

func (pd *PageDash) new(data map[string]any) *PdCtx {

	return &PdCtx{
		Data:    data,
		Model:   pd.model,
		Message: "",
		Rt:      pd.jsruntime,
	}
}

func (pd *PdCtx) bind() {
	pd.Rt.Set("apply_data", pd.applyData)
	pd.Rt.Set("get_data", pd.getData)
	pd.Rt.Set("get_data_value", pd.getDataValue)
	pd.Rt.Set("set_data_value", pd.setDataValue)

	pd.Rt.Set("get_bind_funcs", func() any {
		return []string{
			"apply_data",
			"get_data",
			"get_data_value",
			"set_data_value",
		}
	})

}

func (pd *PdCtx) execute(method, version string) error {
	var fn func(version string) error
	err := getEntry(pd.Rt, method, &fn)
	if err != nil {
		return err
	}

	return fn(method)
}

func (pd *PdCtx) applyData(data map[string]any) {
	if data == nil {
		return
	}

	for k, v := range pd.Data {
		data[k] = v
	}

	pd.Data = data
}

func (pd *PdCtx) getData() any {
	return pd.Data
}

func (pd *PdCtx) getDataValue(field string) any {
	return pd.Data[field]
}

func (pd *PdCtx) setDataValue(field string, value any) {
	pp.Println("@value", value)

	pd.Data[field] = value
}

// helper

func getEntry(runtime *goja.Runtime, name string, entry interface{}) error {
	rawentry := runtime.Get(name)
	if rawentry == nil {
		return easyerr.NotFound()
	}

	return runtime.ExportTo(rawentry, entry)
}

type HookFunc func(ctx *PdCtx) error

var (
	hookFuncs map[string]HookFunc
	hLock     sync.Mutex
)

func RegisterHookFunc(name string, hfunc HookFunc) {
	hLock.Lock()
	defer hLock.Unlock()

	hookFuncs[name] = hfunc
}
