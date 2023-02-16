package pageform

import (
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/executors/helper"
	"gopkg.in/yaml.v2"
)

type PfBuilder struct {
	app    xtypes.App
	helper *helper.ExecutorHelper
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	h := helper.New("executor_pageform", true)

	return &PfBuilder{
		app:    app.(xtypes.App),
		helper: h,
	}, nil
}

func (pf *PfBuilder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {

	out, _, err := opts.Binder.GetFileWithMeta("form1.yaml") // fixme => get_this from somewhere
	if err != nil {
		return nil, err
	}

	form := &FormModel{}
	err = yaml.Unmarshal(out, form)
	if err != nil {
		return nil, err
	}

	return &Pageform{
		builder: pf,
		model:   form,
	}, nil
}

func (pf *PfBuilder) ExecFile(file string) ([]byte, error) {
	return pf.helper.Serve(file)
}
