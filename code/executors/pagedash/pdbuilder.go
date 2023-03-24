package pagedash

import (
	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/executors/helper"

	gojaExec "github.com/temphia/temphia/code/backend/engine/executors/javascript1/goja"
)

type PdBuilder struct {
	app    xtypes.App
	helper *helper.ExecutorHelper
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {

	h := helper.New("executor_pagedash", true)

	return &PdBuilder{
		app:    app.(xtypes.App),
		helper: h,
	}, nil
}

func (pd *PdBuilder) Instance(opts etypes.ExecutorOption) (etypes.Executor, error) {

	rt := goja.New()

	if opts.File != "" {
		out, _, err := opts.Binder.GetFileWithMeta(opts.File)
		if err != nil {
			return nil, err
		}

		_, err = gojaExec.New(opts.Binder, rt)
		if err != nil {
			return nil, err
		}

		pp.Println(rt.RunString(string(out)))
	}

	return &PageDash{
		builder:   pd,
		jsruntime: rt,
		binder:    opts.Binder,
	}, nil
}

func (pd *PdBuilder) ExecFile(file string) ([]byte, error) {
	return pd.helper.Serve(file)
}
