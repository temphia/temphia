package pagedash

import (
	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
)

type PdCtx struct {
	data    map[string]any
	model   *DashModel
	message string
	ok      bool
	rt      *goja.Runtime
}

func (pd *PageDash) new(data map[string]any) *PdCtx {

	return &PdCtx{
		data:    data,
		model:   pd.model,
		message: "",
		ok:      true,
		rt:      pd.jsruntime,
	}
}

func (pd *PdCtx) bind() {
	pd.rt.Set("apply_data", pd.applyData)
	pd.rt.Set("get_data", pd.getData)
	pd.rt.Set("get_data_value", pd.getDataValue)
	pd.rt.Set("set_data_value", pd.setDataValue)

	pd.rt.Set("get_bind_funcs", func() any {
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
	err := getEntry(pd.rt, method, &fn)
	if err != nil {
		return err
	}

	return fn(method)
}

func (pd *PdCtx) applyData(data map[string]any) {
	if data == nil {
		return
	}

	for k, v := range pd.data {
		data[k] = v
	}

	pd.data = data
}

func (pd *PdCtx) getData() any {
	return pd.data
}

func (pd *PdCtx) getDataValue(field string) any {
	return pd.data[field]
}

func (pd *PdCtx) setDataValue(field string, value any) {
	pd.data[field] = value
}

// helper

func getEntry(runtime *goja.Runtime, name string, entry interface{}) error {
	rawentry := runtime.Get(name)
	if rawentry == nil {
		return easyerr.NotFound()
	}

	return runtime.ExportTo(rawentry, entry)
}
