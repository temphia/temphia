package pageform

import (
	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
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
	pc.rt.Set("GetDataValue", pc.GetDataValue)
	pc.rt.Set("GetData", pc.GetData)
	pc.rt.Set("GetStageItem", pc.GetStageItem)
	pc.rt.Set("GetStage", pc.GetStage)
}

func (pc *PfCtx) execute(name, mode string) error {
	var fn func(mode string) error
	err := setEntry(pc.rt, name, &fn)
	if err != nil {
		return err
	}

	return fn(mode)
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

func (pc *PfCtx) GetDataValue()             {}
func (pc *PfCtx) GetData()                  {}
func (pc *PfCtx) GetStageItem()             {}
func (pc *PfCtx) GetStage()                 {}
func (pc *PfCtx) SetError(msg string)       {}
func (pc *PfCtx) ClearData(except []string) {}
func (pc *PfCtx) DeleteDataField(string)    {}

// helper

func setEntry(runtime *goja.Runtime, name string, entry interface{}) error {
	rawentry := runtime.Get(name)
	if rawentry == nil {
		return easyerr.NotFound()
	}

	return runtime.ExportTo(rawentry, entry)
}
