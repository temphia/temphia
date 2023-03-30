package pageform

import (
	"sync"

	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/thoas/go-funk"
)

type PfCtx struct {
	data           map[string]any
	model          *FormModel
	disabledFields []string
	message        string
	ok             bool
	final          bool
	nextStage      string
	rt             *goja.Runtime
	bindings       bindx.Bindings

	currStage string
}

func (pc *PfCtx) bind() {
	pc.rt.Set("get_data_value", pc.getDataValue)
	pc.rt.Set("get_data", pc.getData)
	pc.rt.Set("get_stage_item", pc.getStageItem)
	pc.rt.Set("get_stage", pc.getStage)
	pc.rt.Set("set_message", pc.setMessage)
	pc.rt.Set("set_final", pc.setFinal)
	pc.rt.Set("clear_data", pc.clearData)
	pc.rt.Set("delete_data_field", pc.deleteDataField)
	pc.rt.Set("set_next_stage", pc.setNextStage)
	pc.rt.Set("pick_next_stage", pc.pickNextStage)
	pc.rt.Set("disable_field", pc.disableField)
	pc.rt.Set("get_bind_funcs", func() any {
		return []string{
			"get_data_value",
			"get_data",
			"get_stage_item",
			"get_stage",
			"set_message",
			"set_final",
			"clear_data",
			"delete_data_field",
			"set_next_stage",
			"pick_next_stage",
			"disable_field",
		}
	})
}

func (pc *PfCtx) execute(name, mode, stage string) error {
	var fn func(mode, stage string) error
	err := getEntry(pc.rt, name, &fn)
	if err != nil {
		return err
	}

	pc.currStage = stage

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

func (pc *PfCtx) getDataValue(field string) any {
	return pc.data[field]
}

func (pc *PfCtx) getData() any {
	return pc.data
}

func (pc *PfCtx) getStageItem(stage, item string) any {

	stg, ok := pc.model.Stages[stage]
	if !ok {
		return nil
	}

	for _, fi := range stg.Items {
		return fi
	}
	return nil
}

func (pc *PfCtx) getStage(stage string) any {
	stg, ok := pc.model.Stages[stage]
	if !ok {
		return nil
	}
	return &stg
}

func (pc *PfCtx) setMessage(msg string, ok bool) {

	pp.Println("@set_message", msg)

	pc.message = "@@@@@@@"
	pc.ok = ok

	pp.Println("@set_message_after_set", pc.message)

}

func (pc *PfCtx) setFinal() {
	pc.final = true
}

func (pc *PfCtx) clearData(except []string) {
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

func (pc *PfCtx) deleteDataField(field string) {
	delete(pc.data, field)
}

func (pc *PfCtx) pickNextStage() {
	for idx, shint := range pc.model.ExecHint {
		if shint == pc.currStage {
			pc.nextStage = pc.model.ExecHint[idx+1]
			break
		}
	}
}

func (pc *PfCtx) disableField(field string) {
	pc.disabledFields = append(pc.disabledFields, field)
}

func (pc *PfCtx) setNextStage(stage string) {
	pc.nextStage = stage
}

// helper

func getEntry(runtime *goja.Runtime, name string, entry interface{}) error {
	rawentry := runtime.Get(name)
	if rawentry == nil {
		return easyerr.NotFound()
	}

	return runtime.ExportTo(rawentry, entry)
}

type HookFunc func(ctx *PfCtx) error

var (
	hookFuncs map[string]HookFunc
	hLock     sync.Mutex
)

func RegisterHookFunc(name string, hfunc HookFunc) {
	hLock.Lock()
	defer hLock.Unlock()

	hookFuncs[name] = hfunc
}
