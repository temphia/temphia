package pageform

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type PageForm struct {
	model *FormModel
	app   xtypes.App
}

func New(app any) (*PageForm, error) {

	return &PageForm{
		model: &FormModel{
			Items:          []FormItem{},
			Data:           make(map[string]any),
			Message:        "",
			ServerOnLoad:   "",
			ServerOnSubmit: "",
			ClientOnLoad:   "",
			ClientOnSubmit: "",
		},

		app: app.(xtypes.App),
	}, nil
}

func (pf *PageForm) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {
	return nil, nil
}

func (pf *PageForm) ExecFile(file string) ([]byte, error) {

	return nil, nil
}
