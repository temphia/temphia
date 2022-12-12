package lifecycle

import "github.com/temphia/temphia/code/executors/backend/wizard/wmodels"

type OnSplashSubmit struct {
	Models      *wmodels.Wizard
	SideEffects OnSplashSubmitSideEffects
	SubmitData  map[string]interface{}
	ExecData    interface{}
}

type OnSplashSubmitSideEffects struct {
	FailErr        string
	NextGroup      string
	SkipValidation bool
}

type OnSplashSubmitCtx struct {
	Type       string
	SubmitData map[string]interface{}
	ExecData   interface{}
}

func (s *OnSplashSubmit) Execute() error {

	return nil
}

func (s *OnSplashSubmit) Bindings() map[string]interface{} {

	return map[string]interface{}{
		"_wizard_set_next_stage_group": func(name string) {
			s.SideEffects.NextGroup = name
		},
		"_wizard_set_err": func(err string) {
			s.SideEffects.FailErr = err
		},

		"_wizard_set_skip_validation": func(skip bool) {
			s.SideEffects.SkipValidation = skip
		},
	}
}
