package lifecycle

import "github.com/temphia/temphia/code/executors/backend/wizard/wmodels"

// before

type StageBeforeVerify struct {
	Models      *wmodels.Wizard
	SideEffects StageBeforeVerifyEffect
	SubData     *wmodels.Submission
}

type StageBeforeVerifyEffect struct {
	FailErr   string
	Errors    map[string]string
	SkipCheck []string
	NextStage string
}

type StageBeforeVerifyCtx struct {
	Type        string
	ParentSubId string
	ParentGroup string
	ParentStage string
	SubId       string
	Stage       string
	Group       string
}

func (s *StageBeforeVerify) Execute() error {
	return nil
}

func (s *StageBeforeVerify) Bindings() map[string]interface{} {

	return map[string]interface{}{
		"_wizard_set_err": func(e string) {
			s.SideEffects.FailErr = e
		},
		"_wizard_set_field_err": func(field, e string) {
			s.SideEffects.Errors[field] = e
		},

		"_wizard_skip_field_check": func(field string) {
			if s.SideEffects.SkipCheck == nil {
				s.SideEffects.SkipCheck = []string{field}
				return
			}
			s.SideEffects.SkipCheck = append(s.SideEffects.SkipCheck, field)
		},
		"_wizard_set_next_stage": func(name string) {
			s.SideEffects.NextStage = name
		},
		"_wizard_set_shared_var": func(name string, data interface{}) {
			s.SubData.SharedVars[name] = data
		},
		"_wizard_get_shared_var": func(name string) interface{} {
			return s.SubData.SharedVars[name]
		},

		"_wizard_get_prev_data": func(name string) interface{} {
			return s.SubData.Data[name]
		},
	}

}

// after

type StageAfterVerify struct {
	Models     *wmodels.Wizard
	SideEffect StageAfterVerifyEffect
	SubData    *wmodels.Submission
}

type StageAfterVerifyEffect struct {
	FailErr   string
	Errors    map[string]string
	NextStage string
}

type StageAfterVerifyCtx struct {
	Type        string
	ParentGroup string
	ParentStage string
	Stage       string
	Group       string
}

func (s *StageAfterVerify) Execute() error {
	return nil
}

func (s *StageAfterVerify) Bindings() map[string]interface{} {

	return map[string]interface{}{
		"_wizard_set_err": func(e string) {
			s.SideEffect.FailErr = e
		},
		"_wizard_set_field_err": func(field, e string) {
			s.SideEffect.Errors[field] = e
		},

		"_wizard_set_next_stage": func(name string) {
			s.SideEffect.NextStage = name
		},
		"_wizard_set_shared_var": func(name string, data interface{}) {
			s.SubData.SharedVars[name] = data
		},
		"_wizard_get_shared_var": func(name string) interface{} {
			return s.SubData.SharedVars[name]
		},

		"_wizard_get_prev_data": func(name string) interface{} {
			return s.SubData.Data[name]
		},
	}

}
