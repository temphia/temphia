package lifecycle

import "github.com/temphia/temphia/code/executors/backend/wizard/wmodels"

type BeforeBack struct {
	SideEffects BeforeBackSideEffects
	SubData     *wmodels.Submission
}

type BeforeBackSideEffects struct {
	FailErr   string
	NextStage string
}

type BeforeBackCtx struct {
	Type         string
	CurrentStage string
	CurrentGroup string
	ParentGroup  string
	ParentStage  string
}

func (b *BeforeBack) Execute() error { return nil }

func (b *BeforeBack) Bindings() map[string]interface{} {

	return map[string]interface{}{

		"_wizard_set_shared_var": func(name string, data interface{}) {
			b.SubData.SharedVars[name] = data
		},
		"_wizard_get_shared_var": func(name string) interface{} {
			return b.SubData.SharedVars[name]
		},

		"_wizard_get_stage_data": func(stage string) map[string]interface{} {
			return b.SubData.Data[stage]
		},

		"_wizard_get_field_data": func(stage, field string) interface{} {
			sdata := b.SubData.Data[stage]
			if sdata == nil {
				return nil
			}

			return sdata[field]
		},

		"_wizard_get_visited_stage": func() []string {
			return b.SubData.VisitedStages
		},

		"_wizard_set_err": func(err string) {
			b.SideEffects.FailErr = err
		},
		"_wizard_set_next_stage": func(name string) {
			b.SideEffects.NextStage = name
		},
	}

}
