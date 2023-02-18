package pageform

import (
	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/thoas/go-funk"
)

type PfCtx struct {
	data           map[string]any
	model          *FormModel
	disabledFields []string
	message        string
	nextStage      string
	rt             *goja.Runtime
}

func (pc *PfCtx) bind() {
	pc.rt.Set("get_data_value", pc.GetDataValue)
	pc.rt.Set("get_data", pc.GetData)
	pc.rt.Set("get_stage_item", pc.GetStageItem)
	pc.rt.Set("get_stage", pc.GetStage)
}

func (pc *PfCtx) execute(name, mode, stage string) error {
	var fn func(mode, stage string) error
	err := getEntry(pc.rt, name, &fn)
	if err != nil {
		return err
	}

	return fn(mode, stage)
}

func (pc *PfCtx) applyData(data map[string]any) {
	if data == nil {
		return
	}

	for k, v := range pc.data {
		data[k] = v
	}

	pc.data = data
}

// binds

func (pc *PfCtx) GetDataValue(field string) any {
	return pc.data[field]
}

func (pc *PfCtx) GetData() any {
	return pc.data
}

func (pc *PfCtx) GetStageItem(stage, item string) any {

	stg, ok := pc.model.Stages[stage]
	if !ok {
		return nil
	}

	for _, fi := range stg.Items {
		return fi
	}
	return nil
}

func (pc *PfCtx) GetStage(stage string) any {
	stg, ok := pc.model.Stages[stage]
	if !ok {
		return nil
	}
	return &stg
}

func (pc *PfCtx) SetError(msg string) {
	pc.message = msg
}

func (pc *PfCtx) ClearData(except []string) {
	if (except) == nil {
		pc.data = map[string]any{}
		return
	}

	for k := range pc.data {
		if funk.ContainsString(except, k) {
			continue
		}
		delete(pc.data, k)
	}

}

func (pc *PfCtx) DeleteDataField(field string) {
	delete(pc.data, field)
}

// helper

func getEntry(runtime *goja.Runtime, name string, entry interface{}) error {
	rawentry := runtime.Get(name)
	if rawentry == nil {
		return easyerr.NotFound()
	}

	return runtime.ExportTo(rawentry, entry)
}
