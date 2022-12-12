package lifecycle

import "github.com/temphia/temphia/code/executors/backend/wizard/wmodels"

type OnSplashLoad struct {
	Models      *wmodels.Wizard
	SideEffects OnSplashLoadSideEffects
	HasExecData bool
}

type OnSplashLoadSideEffects struct {
	FailErr     string
	SkipSplash  bool
	DataSources map[string]interface{}
}

type OnSplashLoadCtx struct {
	Type        string
	HasExecData bool
}

func (s *OnSplashLoad) Execute() error {
	return nil
}

func (s *OnSplashLoad) Bindings() map[string]interface{} {
	return map[string]interface{}{
		"_wizard_set_err": func(err string) {
			s.SideEffects.FailErr = err
		},
		"_wizard_set_skip_splash": func(skip bool) {
			s.SideEffects.SkipSplash = skip
		},
		"_wizard_set_data_source": func(name string, data interface{}) {
			s.SideEffects.DataSources[name] = data
		},
	}
}
