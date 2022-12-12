package lifecycle

import "github.com/temphia/temphia/code/executors/backend/wizard/wmodels"

// before

type BeforeEnd struct {
	SideEffects BeforeEndSideEffects
	SubData     *wmodels.Submission
}

type BeforeEndSideEffects struct {
	FailErr     string
	GotoStage   string
	GotoMessage string
}

type BeforeEndCtx struct {
	Type        string
	ParentGroup string
	ParentStage string
	ParentSubId string
}

func (b *BeforeEnd) Execute() error                   { return nil }
func (b *BeforeEnd) Bindings() map[string]interface{} { return nil }

// after

type AfterEnd struct {
	SideEffects AfterEndSideEffects
	SubData     *wmodels.Submission
}

type AfterEndSideEffects struct {
	FailErr string
	Message string
}

type AfterEndCtx struct {
	Type        string
	ParentGroup string
	ParentStage string
}

func (b *AfterEnd) Execute() error { return nil }
func (b *AfterEnd) Bindings() map[string]interface{} {

	return map[string]interface{}{
		"_wizard_set_err": func(e string) {
			b.SideEffects.FailErr = e
		},
		"_wizard_set_shared_var": func(name string, data interface{}) {
			b.SubData.SharedVars[name] = data
		},
		"_wizard_get_shared_var": func(name string) interface{} {
			return b.SubData.SharedVars[name]
		},
		"_wizard_get_stage_data": func(name string) interface{} {
			return b.SubData.Data[name]
		},
		"_wizard_set_message": func(m string) {
			b.SideEffects.Message = m
		},
	}

}
