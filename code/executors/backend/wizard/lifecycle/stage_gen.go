package lifecycle

import "github.com/temphia/temphia/code/executors/backend/wizard/wmodels"

// before

type StageBeforeGenerate struct {
	Models      *wmodels.Wizard
	SideEffects StageBeforeGenerateEffects
	SubData     *wmodels.Submission
}

type StageBeforeGenerateEffects struct {
	FailErr     string
	DataSources map[string]interface{}
}

type StageBeforeGenerateCtx struct {
	Type        string
	ParentSubId string
	ParentGroup string
	ParentStage string
	SubId       string
}

func (s *StageBeforeGenerate) Execute() error {
	return nil
}

func (s *StageBeforeGenerate) Bindings() map[string]interface{} {

	return map[string]interface{}{
		"_wizard_set_err": func(e string) {
			s.SideEffects.FailErr = e
		},
		"_wizard_set_shared_var": func(name string, data interface{}) {
			s.SubData.SharedVars[name] = data
		},
		"_wizard_get_shared_var": func(name string) interface{} {
			return s.SubData.SharedVars[name]
		},
		"_wizard_get_stage_data": func(name string) interface{} {
			return s.SubData.Data[name]
		},
		"_wizard_get_data_source": func(name string) interface{} {
			return s.SideEffects.DataSources[name]
		},
		"_wizard_set_data_source": func(name string, data interface{}) {
			s.SideEffects.DataSources[name] = data
		},
	}

}

// after

type StageAfterGenerate struct {
	Models      *wmodels.Wizard
	SideEffects StageAfterGenerateEffects
	SubData     *wmodels.Submission
}

type StageAfterGenerateEffects struct {
	FailErr     string
	DataSources map[string]interface{}
}

type StageAfterGenerateCtx struct {
	Type        string
	ParentSubId string
	ParentGroup string
	ParentStage string
	Stage       string
	Group       string
}

func (s *StageAfterGenerate) Execute() error {
	return nil
}

func (s *StageAfterGenerate) Bindings() map[string]interface{} {

	return map[string]interface{}{
		"_wizard_set_err": func(e string) {
			s.SideEffects.FailErr = e
		},
		"_wizard_set_shared_var": func(name string, data interface{}) {
			s.SubData.SharedVars[name] = data
		},
		"_wizard_get_shared_var": func(name string) interface{} {
			return s.SubData.SharedVars[name]
		},
		"_wizard_get_stage_data": func(name string) interface{} {
			return s.SubData.Data[name]
		},

		"_wizard_get_data_source": func(name string) interface{} {
			return s.SideEffects.DataSources[name]
		},

		"_wizard_set_data_source": func(name string, data interface{}) {
			s.SideEffects.DataSources[name] = data
		},
	}

}
