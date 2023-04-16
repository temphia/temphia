package pagequery

import (
	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

type PqLoadCtx struct {
	Binder  bindx.Bindings
	Model   *PgModel
	Message string
	Rt      *goja.Runtime

	Data map[string]any
}

/*

func (pf *PageQuery) new(data map[string]any) *PdCtx {

	return &PdCtx{
		Data:    data,
		Model:   pf.model,
		Message: "",
		Rt:      pf.jsruntime,
		Binder:  pf.binder,
	}
}

func (ctx *PdCtx) bind() {
	ctx.Rt.Set("apply_data", ctx.applyData)
	ctx.Rt.Set("get_data", ctx.getData)
	ctx.Rt.Set("get_data_value", ctx.getDataValue)
	ctx.Rt.Set("set_data_value", ctx.setDataValue)

	ctx.Rt.Set("get_bind_funcs", func() any {
		return []string{
			"apply_data",
			"get_data",
			"get_data_value",
			"set_data_value",
		}
	})

}

func (ctx *PdCtx) execute(method, version string) error {
	var fn func(version string) error
	err := getEntry(ctx.Rt, method, &fn)
	if err != nil {
		return err
	}

	return fn(method)
}

func (ctx *PdCtx) applyData(data map[string]any) {
	if data == nil {
		return
	}

	for k, v := range ctx.Data {
		data[k] = v
	}

	ctx.Data = data
}

func (ctx *PdCtx) getData() any {
	return ctx.Data
}

func (ctx *PdCtx) getDataValue(field string) any {
	return ctx.Data[field]
}

func (ctx *PdCtx) setDataValue(field string, value any) {
	pp.Println("@value", value)

	ctx.Data[field] = value
}

// helper

func getEntry(runtime *goja.Runtime, name string, entry interface{}) error {
	rawentry := runtime.Get(name)
	if rawentry == nil {
		return easyerr.NotFound()
	}

	return runtime.ExportTo(rawentry, entry)
}

*/
