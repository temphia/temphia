package pagequery

import (
	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/libx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

type ctxResponse struct {
	Stage    string             `json:"stage,omitempty" yaml:"stage,omitempty"`
	Data     map[string]any     `json:"data,omitempty" yaml:"data,omitempty"`
	Elements map[string]Element `json:"elements,omitempty" yaml:"elements,omitempty"`
}

type PqLoadCtx struct {
	parent    *PageQuery
	Binder    bindx.Bindings
	Model     *PgModel
	Rt        *goja.Runtime
	ExecData  map[string]any
	ParamData map[string]string
	Stage     string
}

func (pf *PageQuery) new(ed map[string]any, pd map[string]string, stage string) PqLoadCtx {
	return PqLoadCtx{
		parent:    pf,
		Binder:    pf.binder,
		Model:     pf.model,
		Rt:        pf.jsruntime,
		ExecData:  ed,
		ParamData: pd,
		Stage:     stage,
	}
}

func (ctx *PqLoadCtx) execute(script string) (*ctxResponse, error) {

	ctx.bind()

	var err error
	var val goja.Value

	perr := libx.PanicWrapper(func() {
		val, err = ctx.Rt.RunString(script)
	})
	if perr != nil {
		return nil, perr
	}

	if err != nil {
		return nil, err
	}

	cresp := &ctxResponse{}

	err = ctx.Rt.ExportTo(val, cresp)
	if err != nil {
		return nil, err
	}

	return cresp, nil
}

func (ctx *PqLoadCtx) bind() {
	ctx.Rt.Set("get_execdata", ctx.getExecdata)
	ctx.Rt.Set("get_execdata_item", ctx.getExecdataItem)
	ctx.Rt.Set("get_paramdata", ctx.getParamdata)
	ctx.Rt.Set("get_paramdata_item", ctx.getParamdataItem)
	ctx.Rt.Set("get_stage", ctx.getStage)

	ctx.Rt.Set("get_bind_funcs", func() any {
		return []string{
			"get_execdata",
			"get_execdata_item",
			"get_paramdata",
			"get_paramdata_item",
			"get_stage",
		}
	})
}

func (ctx *PqLoadCtx) getExecdata() any {
	return ctx.ExecData
}

func (ctx *PqLoadCtx) getExecdataItem(name string) any {
	return ctx.ExecData[name]
}

func (ctx *PqLoadCtx) getParamdata() any {
	return ctx.ParamData
}

func (ctx *PqLoadCtx) getParamdataItem(name string) any {
	return ctx.ParamData[name]
}

func (ctx *PqLoadCtx) getStage() string {
	return ctx.Stage
}
