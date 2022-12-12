package lifecycle

import "github.com/temphia/temphia/code/executors/backend/wizard/wmodels"

// before

type BeforeStart struct {
	SideEffects BeforeStartSideEffects
	ParentGroup string
	ParentStage string
	ExecData    interface{}
}

type BeforeStartSideEffects struct {
	FailErr   string
	NextStage string
}

type BeforeStartCtx struct {
	ParentGroup string
	ParentStage string
	Type        string
	ExecData    interface{}
}

func (b *BeforeStart) Execute() error { return nil }

func (b *BeforeStart) Bindings() map[string]interface{} {
	return map[string]interface{}{
		"_wizard_set_next_stage": func(name string) {
			b.SideEffects.NextStage = name
		},

		"_wizard_set_err": func(err string) {
			b.SideEffects.FailErr = err
		},
	}
}

// after

type AfterStart struct {
	SideEffects AfterStartSideEffects
	SubData     *wmodels.Submission
	ExecData    interface{}
}

type AfterStartCtx struct {
	Type        string
	ParentGroup string
	ParentStage string
	ParentSubId string
	SubId       string
	ExecData    interface{}
}

type AfterStartSideEffects struct {
	FailErr  string
	PrevData map[string]map[string]interface{} // map exec_data => prev_data
}

func (b *AfterStart) Execute() error {

	return nil
}

func (b *AfterStart) Bindings() map[string]interface{} {

	return map[string]interface{}{
		"_wizard_set_prev_data_field": func(stage, field string, value interface{}) {
			if b.SideEffects.PrevData == nil {
				b.SideEffects.PrevData = make(map[string]map[string]interface{})
			}

			spd := b.SideEffects.PrevData[stage]
			if spd == nil {
				spd = make(map[string]interface{})
				b.SideEffects.PrevData[stage] = spd
			}
			spd[field] = value
		},

		"_wizard_set_prev_data_stage": func(stage string, pdata map[string]interface{}) {
			if b.SideEffects.PrevData == nil {
				b.SideEffects.PrevData = make(map[string]map[string]interface{})
			}
			b.SideEffects.PrevData[stage] = pdata
		},

		"_wizard_set_prev_data": func(pdata map[string]map[string]interface{}) {
			b.SideEffects.PrevData = pdata
		},

		"_wizard_set_err": func(err string) {
			b.SideEffects.FailErr = err
		},
		"_wizard_set_shared_var": func(name string, data interface{}) {
			b.SubData.SharedVars[name] = data
		},
		"_wizard_get_shared_var": func(name string) interface{} {
			return b.SubData.SharedVars[name]
		},
	}
}
